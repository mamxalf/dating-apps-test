package service

import (
	"context"
	"dating-apps/internal/domains/dating/model"
	"dating-apps/internal/domains/dating/model/dto"
	"github.com/rs/zerolog/log"
)

func (u *DatingServiceImpl) GetDatingProfile(ctx context.Context, req dto.GetDatingProfileRequest) (res []model.Profile, err error) {
	exceptIDs, err := u.DatingRepository.GetSwipeCacheListID(req.UserID)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile] failed get swipe cache")
		return
	}

	res, err = u.DatingRepository.GetProfile(ctx, exceptIDs, req.Limit, req.Offset)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile] failed get profile")
		return
	}
	return
}
