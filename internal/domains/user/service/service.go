package service

import (
	"context"
	"dating-apps/configs"
	"dating-apps/internal/domains/user/model/dto"
	"dating-apps/internal/domains/user/repository"

	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (err error)
	Login(ctx context.Context, req dto.LoginRequest) (res dto.LoginResponse, err error)
	GetUserByEmail(email string) (res dto.UserResponse, err error)
	UpdateUserProfile(ctx context.Context, req dto.UpdateUserProfileRequest) (err error)
	SubscribeUserPremium(ctx context.Context, userID uuid.UUID) (err error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Config         *configs.Config
}

// ProvideUserServiceImpl is the provider for this service.
func ProvideUserServiceImpl(userRepository repository.UserRepository, config *configs.Config) *UserServiceImpl {
	s := new(UserServiceImpl)
	s.UserRepository = userRepository
	s.Config = config

	return s
}
