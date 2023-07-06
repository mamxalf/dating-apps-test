package user

import (
	"github.com/google/uuid"
	"time"
)

type Register struct {
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

type UserSession struct {
	ID           uuid.UUID `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
