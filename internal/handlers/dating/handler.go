package dating

import (
	"dating-apps/http/middleware"
	"dating-apps/internal/domains/dating/service"
	"github.com/go-chi/chi"
)

type DatingHandler struct {
	DatingService service.DatingService
	JWTMiddleware *middleware.JWT
}

func ProvideDatingHandler(datingService service.DatingService, jwt *middleware.JWT) DatingHandler {
	return DatingHandler{
		DatingService: datingService,
		JWTMiddleware: jwt,
	}
}

func (h *DatingHandler) Router(r chi.Router) {
	r.Route("/datings", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(h.JWTMiddleware.VerifyToken)
			r.Post("/swipe", h.SwipeProfile)
			r.Get("/swipe", h.GetSwipeProfile)
		})

	})
}
