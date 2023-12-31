package service

import (
	"context"
	"dating-apps/internal/domains/dating/model/dto"
	"dating-apps/shared/failure"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	maxAttempt = 10
)

func (u *DatingServiceImpl) SwipeProfile(ctx context.Context, req dto.SwipeRequest) (err error) {
	user, err := u.UserRepository.GetUserProfileByUserID(req.UserID)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile]")
		return
	}

	tx, err := u.DatingRepository.BeginTx()
	if err != nil {
		log.Err(err).Msg("[SwipeProfile] failed begin transactions")
		return
	}

	err = u.setDatingCacheSchema(req, user.IsVerified)
	if err != nil {
		log.Err(err).Msg("[SwipeProfile - setDatingCacheSchema] failed cache data")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal().Err(err).Msg("Rollback")
		}
		return
	}

	newSwipe := req.ToModel()
	if err = u.DatingRepository.SwipeProfile(ctx, &newSwipe, tx); err != nil {
		if failure.GetCode(err) != http.StatusNotFound {
			log.Error().Interface("params", req).Err(err).Msg("[Register - Service]")
		}
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal().Err(err).Msg("Rollback")
		}
		return
	}

	if commitErr := tx.Commit(); commitErr != nil {
		log.Fatal().Err(err).Msg("Commit")
	}
	return
}

func (u *DatingServiceImpl) setDatingCacheSchema(req dto.SwipeRequest, isVerified bool) (err error) {
	// trigger swipe
	listSwipedIDS, err := u.DatingRepository.GetSwipeCacheListID(req.UserID)
	if err != nil {
		err = failure.InternalError(err)
		log.Err(err).Msg("[SwipeProfile] failed get swipe cache")
		return
	}

	if strings.Contains(strings.Join(listSwipedIDS, ","), req.ProfileID.String()) {
		err = failure.BadRequestFromString("swipe twice in a profile")
		log.Err(err).Msg("[SwipeProfile] swipe twice in a profile")
		return
	}

	amount, err := u.DatingRepository.SwipeIncr(req.UserID)
	if err != nil {
		log.Err(err).Msg("[SwipeProfile] failed incr redis")
		return
	}

	if !isVerified {
		if amount > maxAttempt {
			err = failure.UnprocessableEntity("Reach limit for swipe profile!")
			log.Warn().Err(err).Msg("[SwipeProfile]")
			return
		}
	}

	if amount <= 1 {
		err = u.DatingRepository.SwipeCacheExpiry(req.UserID)
		if err != nil {
			log.Err(err).Msg("[SwipeProfile] failed incr redis")
			return
		}
	}

	err = u.DatingRepository.SwipeCacheListID(req.UserID, req.ProfileID)
	if err != nil {
		log.Err(err).Msg("[SwipeProfile] failed push data")
		return
	}

	return
}
