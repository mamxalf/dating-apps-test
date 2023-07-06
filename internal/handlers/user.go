package handlers

import (
	"dating-apps/internal/domains/user"
	"dating-apps/shared/response"
	"github.com/go-chi/chi"
	"net/http"
)

type UserHandler struct {
	UserService user.UserService
}

func ProvideUserHandler(userService user.UserService) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) Router(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			//r.Use(h.AuthMiddleware.ClientCredential)
			r.Get("/ping", h.ResolvePing)
		})

	})
}

// ResolvePing resolves a Foo by its ID.
// @Summary Resolve Foo by ID
// @Description This endpoint resolves a Foo by its ID.
// @Tags Users
// @Security EVMOauthToken
// @Param id path string true "The Foo's identifier."
// @Param withItems query string false "Fetch with items, default false."
// @Produce json
// @Success 200 {object} response.Base{}
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/users/ping [get]
func (h *UserHandler) ResolvePing(w http.ResponseWriter, r *http.Request) {
	foo, err := h.UserService.ResolvePing()
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, foo)
}
