package service

import (
	"context"
	"dating-apps/configs"
	"dating-apps/internal/domains/dating/model/dto"
	"dating-apps/internal/domains/dating/repository"
)

type DatingService interface {
	SwipeProfile(ctx context.Context, req dto.SwipeRequest) (err error)
	GetDatingProfile(ctx context.Context, req dto.GetDatingProfileRequest) (result dto.ResponseProfile, err error)
}

type DatingServiceImpl struct {
	DatingRepository repository.DatingRepository
	Config           *configs.Config
}

// ProvideDatingServiceImpl is the provider for this service.
func ProvideDatingServiceImpl(userRepository repository.DatingRepository, config *configs.Config) *DatingServiceImpl {
	s := new(DatingServiceImpl)
	s.DatingRepository = userRepository
	s.Config = config

	return s
}
