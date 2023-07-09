package service

import (
	"context"
	"dating-apps/configs"
	"dating-apps/internal/domains/dating/model/dto"
	datingRepo "dating-apps/internal/domains/dating/repository"
	userRepo "dating-apps/internal/domains/user/repository"
)

type DatingService interface {
	SwipeProfile(ctx context.Context, req dto.SwipeRequest) (err error)
	GetDatingProfile(ctx context.Context, req dto.GetDatingProfileRequest) (result dto.ResponseProfile, err error)
}

type DatingServiceImpl struct {
	DatingRepository datingRepo.DatingRepository
	UserRepository   userRepo.UserRepository
	Config           *configs.Config
}

// ProvideDatingServiceImpl is the provider for this service.
func ProvideDatingServiceImpl(
	userRepository userRepo.UserRepository,
	datingRepository datingRepo.DatingRepository,
	config *configs.Config,
) *DatingServiceImpl {
	s := new(DatingServiceImpl)
	s.DatingRepository = datingRepository
	s.UserRepository = userRepository
	s.Config = config

	return s
}
