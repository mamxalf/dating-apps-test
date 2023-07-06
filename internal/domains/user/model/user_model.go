package model

import (
	"github.com/google/uuid"
	"time"
)

type UserRegister struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type User struct {
	ID         uuid.UUID `db:"id"`
	Username   string    `db:"username"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	IsVerified bool      `db:"is_verified"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
