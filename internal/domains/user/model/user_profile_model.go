package model

import (
	"github.com/google/uuid"
	"time"
)

type UserProfile struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	FullName  string    `db:"full_name"`
	Age       int       `db:"age"`
	Gender    string    `db:"gender"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type FullUserProfile struct {
	ID         uuid.UUID `db:"id"`
	Username   string    `db:"username"`
	Email      string    `db:"email"`
	FullName   string    `db:"full_name"`
	Age        int       `db:"age"`
	Gender     string    `db:"gender"`
	IsVerified bool      `db:"is_verified"`
}
