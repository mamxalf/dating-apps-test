package user

import (
	"dating-apps/http/middleware"
	"dating-apps/internal/domains/user/model/dto"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Me get info user.
// @Summary Me User
// @Description This endpoint for Me User.
// @Tags Users
// @Accept  json
// @Produce json
// @Security BearerToken
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/users/me [get]
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	res, err := h.UserService.GetUserByEmail(claimUser["email"].(string))
	if err != nil {
		log.Warn().Err(err).Msg("[Register Handler]")
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// UpdateUserProfile get info user.
// @Summary UpdateUserProfile User
// @Description This endpoint for UpdateUserProfile User.
// @Tags Users
// @Accept  json
// @Produce json
// @Security BearerToken
// @Param request body dto.UpdateUserProfileRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/users/profile [post]
func (h *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var updateUserProfileRequest dto.UpdateUserProfileRequest
	if err := decoder.Decode(&updateUserProfileRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := updateUserProfileRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	userID, err := uuid.Parse(claimUser["owner_id"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	updateUserProfileRequest.UserID = userID
	err = h.UserService.UpdateUserProfile(r.Context(), updateUserProfileRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Register Handler]")
		response.WithError(w, err)
		return
	}
	response.WithMessage(w, http.StatusOK, "Success insert profile!")
}

// UserSubscribePremium get info user.
// @Summary UserSubscribePremium User
// @Description This endpoint for UserSubscribePremium User.
// @Tags Users
// @Accept  json
// @Produce json
// @Security BearerToken
// @Param request body dto.UserSubscriptionRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/users/subscription [post]
func (h *UserHandler) UserSubscribePremium(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var UserSubscriptionRequest dto.UserSubscriptionRequest
	if err := decoder.Decode(&UserSubscriptionRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := UserSubscriptionRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	userID, err := uuid.Parse(claimUser["owner_id"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	UserSubscriptionRequest.UserID = userID
	err = h.UserService.SubscribeUserPremium(r.Context(), userID)
	if err != nil {
		log.Warn().Err(err).Msg("[UserSubscribePremium Handler]")
		response.WithError(w, err)
		return
	}
	response.WithMessage(w, http.StatusOK, "Success update premium user!")
}
