package auth_data

import (
	"encoding/json"

	"github.com/google/uuid"
)

type UserData struct {
	ID          uuid.UUID
	Email       string
	DisplayName string
	Preferences json.RawMessage
}
