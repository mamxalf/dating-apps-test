package dto

import (
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared"
	"github.com/google/uuid"
)

type SwipeRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	ProfileID uuid.UUID `validate:"required" json:"profile_id"`
	IsLike    bool      `validate:"required" json:"is_like"`
}

func (r *SwipeRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(r)
}

func (r *SwipeRequest) ToModel() (swipe model.NewSwipe) {
	swipe = model.NewSwipe{
		UserID:    r.UserID,
		ProfileID: r.ProfileID,
		IsLike:    r.IsLike,
	}
	return
}
