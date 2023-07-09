package service

import (
	"context"
	"dating-apps/shared/failure"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *UserServiceImpl) SubscribeUserPremium(ctx context.Context, userID uuid.UUID) (err error) {
	if err = u.UserRepository.SubscribeUserPremium(ctx, userID); err != nil {
		if failure.GetCode(err) != http.StatusNotFound {
			log.Error().Interface("params", userID).Err(err).Msg("[SubscribeUserPremium - Service]")
			return
		}
		return
	}
	return
}
