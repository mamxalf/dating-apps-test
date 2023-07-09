package dating

import (
	"dating-apps/http/middleware"
	"dating-apps/internal/domains/dating/model/dto"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// SwipeProfile get info user.
// @Summary SwipeProfile User
// @Description This endpoint for SwipeProfile User.
// @Tags Datings
// @Accept  json
// @Produce json
// @Security BearerToken
// @Param request body dto.SwipeRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/datings/swipe [post]
func (h *DatingHandler) SwipeProfile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var swipeRequest dto.SwipeRequest
	if err := decoder.Decode(&swipeRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := swipeRequest.Validate(); err != nil {
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

	swipeRequest.UserID = userID
	err = h.DatingService.SwipeProfile(r.Context(), swipeRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Register Handler]")
		response.WithError(w, err)
		return
	}
	response.WithMessage(w, http.StatusOK, "Success swipe profile!")
}
