package user_commands

import "github.com/google/uuid"

type Preference struct {
	VibeIds  []uuid.UUID
	StyleIds []uuid.UUID
}

type SetUpCompleteBody struct {
	Preferences Preference
}
