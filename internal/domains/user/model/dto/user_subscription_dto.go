package dto

import (
	"dating-apps/shared"

	"github.com/google/uuid"
)

type UserSubscriptionRequest struct {
	UserID uuid.UUID
	Code   string `validate:"required,oneof=GRATIS FREE" json:"code" example:"GRATIS"`
}

func (r *UserSubscriptionRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(r)
}
