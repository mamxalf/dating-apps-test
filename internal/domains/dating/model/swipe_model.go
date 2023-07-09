package model

import (
	"time"

	"github.com/google/uuid"
)

type NewSwipe struct {
	UserID    uuid.UUID `db:"user_id"`
	ProfileID uuid.UUID `db:"profile_id"`
	IsLike    bool      `db:"is_like"`
}

type Swipe struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ProfileID uuid.UUID `db:"profile_id"`
	IsLike    bool      `db:"is_like"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
