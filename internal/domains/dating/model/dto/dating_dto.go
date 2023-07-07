package dto

import (
	"dating-apps/shared"
	"github.com/google/uuid"
)

type GetDatingProfileRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Size   int       `json:"size" validate:"omitempty,gte=0"`
	Page   int       `json:"page" validate:"omitempty,gte=0"`
}

func (r *GetDatingProfileRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(r)
}
