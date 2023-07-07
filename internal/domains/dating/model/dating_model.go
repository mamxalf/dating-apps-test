package model

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	FullName  string    `db:"full_name"`
	Age       int       `db:"age"`
	Gender    string    `db:"gender"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	TotalData int       `db:"total_data"`
}
