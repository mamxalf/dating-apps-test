package dto

import (
	"dating-apps/shared"
	"github.com/google/uuid"
)

type GetDatingProfileRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int       `validate:"required" json:"limit"`
	Offset int       `validate:"required" json:"offset"`
}

func (r *GetDatingProfileRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(r)
}
