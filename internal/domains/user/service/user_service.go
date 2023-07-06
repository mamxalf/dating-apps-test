package service

import (
	"dating-apps/internal/domains/user/model/dto"
	"github.com/rs/zerolog/log"
)

func (u *UserServiceImpl) GetUserByEmail(email string) (res dto.UserResponse, err error) {
	dataUser, err := u.UserRepository.GetUserByEmail(email)
	if err != nil {
		log.Err(err).Msg("[GetUserByEmail - Service]")
		return
	}
	res = dto.UserResponse{
		Username:   dataUser.Username,
		Email:      dataUser.Email,
		IsVerified: dataUser.IsVerified,
		CreatedAt:  dataUser.CreatedAt,
		UpdatedAt:  dataUser.UpdatedAt,
	}
	return
}
