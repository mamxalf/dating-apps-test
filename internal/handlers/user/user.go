package user

import (
	"dating-apps/http/middleware"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"net/http"
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
