package user

import (
	"context"
	"dating-apps/configs"
	"dating-apps/shared/failure"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, req RegisterRequest) (err error)
}

type UserServiceImpl struct {
	UserRepository UserRepository
	Config         *configs.Config
}

// ProvideUserServiceImpl is the provider for this service.
func ProvideUserServiceImpl(userRepository UserRepository, config *configs.Config) *UserServiceImpl {
	s := new(UserServiceImpl)
	s.UserRepository = userRepository
	s.Config = config

	return s
}

func (u *UserServiceImpl) Register(ctx context.Context, req RegisterRequest) (err error) {
	registerModel, err := req.ToModel()
	if err != nil {
		log.Error().Interface("params", req).Err(err).Msg("[Register - Service]")
		return
	}

	if err = u.UserRepository.RegisterNewUser(ctx, &registerModel); err != nil {
		if failure.GetCode(err) != http.StatusNotFound {
			log.Error().Interface("params", req).Err(err).Msg("[Register - Service]")
		}
	}
	return
}
