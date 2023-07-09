package service

import (
	"context"
	"dating-apps/internal/domains/user/model"
	"dating-apps/internal/domains/user/model/dto"
	"dating-apps/shared/failure"
	"dating-apps/shared/token"
	"dating-apps/shared/util"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (u *UserServiceImpl) Register(ctx context.Context, req dto.RegisterRequest) (err error) {
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

func (u *UserServiceImpl) Login(ctx context.Context, req dto.LoginRequest) (res dto.LoginResponse, err error) {
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
	userSessionParams := &model.UserSession{
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

	res = dto.LoginResponse{
		Token:        generateToken.Token,
		RefreshToken: generateToken.RefreshToken,
	}

	return
}
