//go:build wireinject
// +build wireinject

package main

import (
	"dating-apps/configs"
	"dating-apps/http"
	"dating-apps/http/middleware"
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
	user.ProvideUserServiceImpl,
	wire.Bind(new(user.UserService), new(*user.UserServiceImpl)),
	user.ProvideUserRepositoryPostgres,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryPostgres)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainUser,
)

var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	handlers.ProvideUserHandler,
	router.ProvideRouter,
)

var authMiddleware = wire.NewSet(
	middleware.ProvideJWTMiddleware,
)

func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// middleware
		authMiddleware,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
