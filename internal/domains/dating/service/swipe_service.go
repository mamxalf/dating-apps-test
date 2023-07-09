package service

import (
	"context"
	"dating-apps/internal/domains/dating/model/dto"
	"dating-apps/shared/failure"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	maxAttempt = 10
)

func (u *DatingServiceImpl) SwipeProfile(ctx context.Context, req dto.SwipeRequest) (err error) {
	tx, err := u.DatingRepository.BeginTx()
	if err != nil {
		log.Err(err).Msg("[SwipeProfile] failed begin transactions")
		return
	}

	err = u.setDatingCacheSchema(req)
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

func (u *DatingServiceImpl) setDatingCacheSchema(req dto.SwipeRequest) (err error) {
	// trigger swipe
	amount, err := u.DatingRepository.SwipeIncr(req.UserID)
	if err != nil {
		log.Err(err).Msg("[SwipeProfile] failed incr redis")
		return
	}

	if amount > maxAttempt {
		err = failure.UnprocessableEntity("Reach limit for swipe profile!")
		log.Warn().Err(err).Msg("[SwipeProfile]")
		return
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
