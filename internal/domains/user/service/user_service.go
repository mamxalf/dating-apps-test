package service

import (
	"dating-apps/internal/domains/user/model"
	"github.com/rs/zerolog/log"
)

func (u *UserServiceImpl) GetUserByEmail(email string) (user model.User, err error) {
	user, err = u.UserRepository.GetUserByEmail(email)
	if err != nil {
		log.Err(err).Msg("[GetUserByEmail - Service]")
		return
	}
	return
}
