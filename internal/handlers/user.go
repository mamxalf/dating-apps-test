package handlers

import (
	"dating-apps/internal/domains/user"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
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
			r.Post("/register", h.Register)
			r.Post("/login", h.Login)
		})

	})
}

// Register sign up user.
// @Summary Register User
// @Description This endpoint for Register User.
// @Tags Users
// @Accept  json
// @Produce json
// @Param request body user.RegisterRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/users/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userRegisterRequest user.RegisterRequest
	if err := decoder.Decode(&userRegisterRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := userRegisterRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err := h.UserService.Register(r.Context(), userRegisterRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Register Handler]")
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusOK, "User successfully register!")
}

// Login sign in user.
// @Summary Login User
// @Description This endpoint for Login User.
// @Tags Users
// @Accept  json
// @Produce json
// @Param request body user.LoginRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userLoginRequest user.LoginRequest
	if err := decoder.Decode(&userLoginRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := userLoginRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.UserService.Login(r.Context(), userLoginRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Register Handler]")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}
