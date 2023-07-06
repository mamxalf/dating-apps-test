//go:build wireinject
// +build wireinject

package main

import (
	"dating-apps/configs"
	"dating-apps/http"
	router "dating-apps/http/routers"
	"dating-apps/infras"
	"dating-apps/internal/domains/user"
	"dating-apps/internal/handlers"
	"github.com/google/wire"
)

var configurations = wire.NewSet(
	configs.Get,
)

var persistences = wire.NewSet(
	infras.ProvidePostgresConn,
)

var domainUser = wire.NewSet(
	// FooService interface and implementation
	user.ProvideUserServiceImpl,
	wire.Bind(new(user.UserService), new(*user.UserServiceImpl)),
	// FooRepository interface and implementation
	//foobarbaz.ProvideFooRepositoryMySQL,
	//wire.Bind(new(foobarbaz.FooRepository), new(*foobarbaz.FooRepositoryMySQL)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainUser,
)

var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "UserHandler"),
	handlers.ProvideUserHandler,
	router.ProvideRouter,
)

func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// middleware
		//authMiddleware,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
