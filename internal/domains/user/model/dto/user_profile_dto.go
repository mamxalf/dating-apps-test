package dto

import (
	"dating-apps/internal/domains/user/model"
	"dating-apps/shared"
	"github.com/google/uuid"
)

type UpdateUserProfileRequest struct {
	UserID   uuid.UUID
	FullName string `validate:"required" json:"full_name" example:"test"`
	Age      int    `validate:"required,gte=0" json:"age" example:"18"`
	Gender   string `validate:"required,oneof=male female" json:"gender" example:"male"`
}

func (r *UpdateUserProfileRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(r)
}

func (r *UpdateUserProfileRequest) ToModel() (register model.UserProfile, err error) {
	register = model.UserProfile{
		UserID:   r.UserID,
		FullName: r.FullName,
		Age:      r.Age,
		Gender:   r.Gender,
	}
	return
}
