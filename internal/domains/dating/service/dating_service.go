package service

import (
	"context"
	"dating-apps/internal/domains/dating/model/dto"

	"github.com/rs/zerolog/log"
)

func (u *DatingServiceImpl) GetDatingProfile(ctx context.Context, req dto.GetDatingProfileRequest) (result dto.ResponseProfile, err error) {
	user, err := u.UserRepository.GetUserProfileByUserID(req.UserID)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile]")
		return
	}

	exceptIDs, err := u.DatingRepository.GetSwipeCacheListID(req.UserID)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile] failed get swipe cache")
		return
	}

	var findByGender string
	if user.Gender == "male" {
		findByGender = "female"
	} else {
		findByGender = "male"
	}

	res, err := u.DatingRepository.GetProfile(ctx, exceptIDs, findByGender, req.Page, req.Size)
	if err != nil {
		log.Err(err).Msg("[GetDatingProfile] failed get profile")
		return
	}

	result.Profiles = res
	if len(res) > 0 {
		result.Pagination.SetPagination(res[0].TotalData, req.Page, req.Size)
	}

	return
}
