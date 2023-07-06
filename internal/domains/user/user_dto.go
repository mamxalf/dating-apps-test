package user

import (
	"dating-apps/shared"
	"dating-apps/shared/util"
	"github.com/rs/zerolog/log"
)

type RegisterRequest struct {
	Username        string `validate:"required" json:"username" example:"test"`
	Email           string `validate:"required,email" json:"email" example:"test@example.com"`
	Password        string `validate:"required,alphanum,min=8,max=30" json:"password,omitempty" example:"s3Cr3Tk3y"`
	ConfirmPassword string `validate:"required_with=Password,eqfield=Password" json:"confirm_password"`
}

func (r *RegisterRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(r)
}

func (r *RegisterRequest) ToModel() (register Register, err error) {
	hashPassword, err := util.HashPassword(r.Password)
	if err != nil {
		log.Err(err).Msg("[Hash Password]")
		return
	}
	register = Register{
		Username: r.Username,
		Email:    r.Email,
		Password: hashPassword,
	}
	return
}
