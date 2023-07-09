package service

import (
	"context"
	"dating-apps/internal/domains/user/model/dto"
	"dating-apps/shared/failure"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (u *UserServiceImpl) UpdateUserProfile(ctx context.Context, req dto.UpdateUserProfileRequest) (err error) {
	registerModel, err := req.ToModel()
	if err != nil {
		log.Error().Interface("params", req).Err(err).Msg("[UpdateUserProfile - Service]")
		return
	}

	if err = u.UserRepository.InsertUserprofile(ctx, &registerModel); err != nil {
		if failure.GetCode(err) != http.StatusNotFound {
			log.Error().Interface("params", req).Err(err).Msg("[UpdateUserProfile - Service]")
			return
		}
		return
	}
	return
}
