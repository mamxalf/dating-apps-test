package repository

import (
	"dating-apps/shared/failure"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (repo *DatingRepositoryPostgres) SwipeIncr(userID uuid.UUID) (amount int64, err error) {
	amount, err = repo.Cache.Client.Incr(userID.String()).Result()
	if err != nil {
		log.Error().Err(err).Msg("[SwipeIncr - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *DatingRepositoryPostgres) SwipeCacheListID(userID, profileID uuid.UUID) (err error) {
	err = repo.Cache.Client.RPush(userID.String(), profileID).Err()
	if err != nil {
		log.Error().Err(err).Msg("[SwipeCacheListID - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *DatingRepositoryPostgres) GetSwipeCacheListID(userID uuid.UUID) (data []string, err error) {
	data, err = repo.Cache.Client.LRange(userID.String(), 0, -1).Result()
	if err != nil {
		log.Error().Err(err).Msg("[GetSwipeCacheListIDSwipeCacheListID - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *DatingRepositoryPostgres) SwipeCacheExpiry(userID uuid.UUID) (err error) {
	expiration := 24 * time.Hour
	err = repo.Cache.Client.Expire(userID.String(), expiration).Err()
	if err != nil {
		log.Error().Err(err).Msg("[SwipeCacheExpiry - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}
