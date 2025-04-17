// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PreferenceMetadatum struct {
	ID        uuid.UUID
	Key       string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          uuid.UUID
	Email       string
	Password    string
	DisplayName string
	Preferences json.RawMessage
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
