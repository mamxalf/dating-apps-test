package service

import (
	"context"
	"dating-apps/internal/domains/dating/model/dto"
	"dating-apps/shared/failure"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (u *DatingServiceImpl) SwipeProfile(ctx context.Context, req dto.SwipeRequest) (err error) {
	newSwipe := req.ToModel()
	if err = u.DatingRepository.SwipeProfile(ctx, &newSwipe); err != nil {
		if failure.GetCode(err) != http.StatusNotFound {
			log.Error().Interface("params", req).Err(err).Msg("[Register - Service]")
		}
	}
	return
}
