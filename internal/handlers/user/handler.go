package user

import (
	"dating-apps/http/middleware"
	"dating-apps/internal/domains/user/service"

	"github.com/go-chi/chi"
)

type UserHandler struct {
	UserService   service.UserService
	JWTMiddleware *middleware.JWT
}

func ProvideUserHandler(userService service.UserService, jwt *middleware.JWT) UserHandler {
	return UserHandler{
		UserService:   userService,
		JWTMiddleware: jwt,
	}
}

func (h *UserHandler) Router(r chi.Router) {
	r.Route("/sessions", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/register", h.Register)
			r.Post("/login", h.Login)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(h.JWTMiddleware.VerifyToken)
			r.Get("/me", h.Me)
			r.Post("/profile", h.UpdateUserProfile)
		})
	})
}
