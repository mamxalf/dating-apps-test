package user

import "dating-apps/configs"

type UserService interface {
	ResolvePing() (ping Ping, err error)
}

type UserServiceImpl struct {
	Config *configs.Config
}

// ProvideUserServiceImpl is the provider for this service.
func ProvideUserServiceImpl(config *configs.Config) *UserServiceImpl {
	s := new(UserServiceImpl)
	s.Config = config

	return s
}

func (u *UserServiceImpl) ResolvePing() (ping Ping, err error) {
	ping = Ping{Message: "test"}
	return
}
