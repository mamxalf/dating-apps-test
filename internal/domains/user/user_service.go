package user

import (
	"context"
	"dating-apps/configs"
	"dating-apps/shared/failure"
	"dating-apps/shared/token"
	"dating-apps/shared/util"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, req RegisterRequest) (err error)
	Login(ctx context.Context, req LoginRequest) (res LoginResponse, err error)
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

func (u *UserServiceImpl) Login(ctx context.Context, req LoginRequest) (res LoginResponse, err error) {
	user, err := u.UserRepository.GetUserByEmail(req.Email)
	if err != nil {
		log.Err(err).Msg("[Login - Service]")
		return
	}

	err = util.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Wrong Password")
		err = fmt.Errorf("invalid Password")
		return
	}

	generateTokenParams := &token.GenerateTokenParams{
		AccessTokenSecret:  u.Config.Auth.AccessTokenSecret,
		RefreshTokenSecret: u.Config.Auth.RefreshTokenSecret,
		AccessTokenExpiry:  u.Config.Auth.AccessTokenExpiry,
		RefreshTokenExpiry: u.Config.Auth.RefreshTokenExpiry,
	}
	userData := &token.UserData{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}
	generateToken, err := token.GenerateToken(userData, generateTokenParams)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Generate Token Error")
		return
	}
	userSessionParams := &UserSession{
		UserID:       user.ID,
		AccessToken:  generateToken.Token,
		RefreshToken: generateToken.RefreshToken,
		IsActive:     true,
	}
	err = u.UserRepository.CreateUserSession(ctx, userSessionParams)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Generate Token Error")
		return
	}

	res = LoginResponse{
		Token:        generateToken.Token,
		RefreshToken: generateToken.RefreshToken,
	}

	return
}
