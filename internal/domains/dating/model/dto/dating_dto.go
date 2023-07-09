package dto

import (
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared"
	"dating-apps/shared/model/dto"

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

func (r *GetDatingProfileRequest) SetDefaultFilter(userID uuid.UUID) {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.Size <= 0 {
		r.Size = 10
	}

	r.UserID = userID
}

type ResponseProfile struct {
	Profiles []model.Profile `json:"profiles"`
	dto.Pagination
}
