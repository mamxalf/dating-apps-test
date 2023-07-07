package dating

import (
	"dating-apps/http/middleware"
	"dating-apps/internal/domains/dating/model/dto"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

// GetSwipeProfile get info user.
// @Summary GetSwipeProfile User
// @Description This endpoint for GetSwipeProfile User.
// @Tags Datings
// @Accept  json
// @Produce json
// @Security BearerToken
// @Param size query int false "data size per page" default(15)
// @Param page query int false "number of page" default(1)
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/datings/swipe [get]
func (h *DatingHandler) GetSwipeProfile(w http.ResponseWriter, r *http.Request) {
	var filter dto.GetDatingProfileRequest

	query := r.URL.Query()
	filter.Page, _ = strconv.Atoi(query.Get("page"))
	filter.Size, _ = strconv.Atoi(query.Get("size"))

	if err := filter.Validate(); err != nil {
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

	filter.UserID = userID
	res, err := h.DatingService.GetDatingProfile(r.Context(), filter)
	if err != nil {
		log.Warn().Err(err).Msg("[GetSwipeProfile Handler]")
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
