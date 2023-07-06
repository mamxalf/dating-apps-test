package user

import (
	"dating-apps/internal/domains/user/model/dto"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Register sign up user.
// @Summary Register User
// @Description This endpoint for Register User.
// @Tags Sessions
// @Accept  json
// @Produce json
// @Param request body dto.RegisterRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/sessions/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userRegisterRequest dto.RegisterRequest
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
// @Tags Sessions
// @Accept  json
// @Produce json
// @Param request body dto.LoginRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/sessions/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userLoginRequest dto.LoginRequest
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
