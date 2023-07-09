package repository

import (
	"dating-apps/shared/failure"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (repo *DatingRepositoryPostgres) SwipeIncr(userID uuid.UUID) (amount int64, err error) {
	key := fmt.Sprintf("user::swipe_amount::%s", userID.String())
	amount, err = repo.Cache.Client.Incr(key).Result()
	if err != nil {
		log.Error().Err(err).Msg("[SwipeIncr - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *DatingRepositoryPostgres) SwipeCacheListID(userID, profileID uuid.UUID) (err error) {
	key := fmt.Sprintf("user::swipe_list::%s", userID.String())
	val := profileID.String()
	err = repo.Cache.Client.RPush(key, val).Err()
	if err != nil {
		log.Error().Err(err).Msg("[SwipeCacheListID - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *DatingRepositoryPostgres) GetSwipeCacheListID(userID uuid.UUID) (data []string, err error) {
	key := fmt.Sprintf("user::swipe_list::%s", userID.String())
	data, err = repo.Cache.Client.LRange(key, 0, -1).Result()
	if err != nil {
		log.Error().Err(err).Msg("[GetSwipeCacheListIDSwipeCacheListID - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *DatingRepositoryPostgres) SwipeCacheExpiry(userID uuid.UUID) (err error) {
	key := fmt.Sprintf("user::swipe_list::%s", userID.String())
	expiration := 24 * time.Hour
	err = repo.Cache.Client.Expire(key, expiration).Err()
	if err != nil {
		log.Error().Err(err).Msg("[SwipeCacheExpiry - Repository] failed exec cache")
		err = failure.InternalError(err)
		return
	}
	return
}
