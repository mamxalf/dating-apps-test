package service

import (
	"context"
	"dating-apps/internal/domains/dating/model/dto"

	"github.com/rs/zerolog/log"
)

func (u *DatingServiceImpl) GetDatingProfile(ctx context.Context, req dto.GetDatingProfileRequest) (result dto.ResponseProfile, err error) {
	exceptIDs, err := u.DatingRepository.GetSwipeCacheListID(req.UserID)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile] failed get swipe cache")
		return
	}

	res, err := u.DatingRepository.GetProfile(ctx, exceptIDs, "male", req.Page, req.Size)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile] failed get profile")
		return
	}

	result.Profiles = res
	result.Pagination.SetPagination(res[0].TotalData, req.Page, req.Size)

	return
}
