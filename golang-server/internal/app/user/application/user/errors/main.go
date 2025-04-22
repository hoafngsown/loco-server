package user_errors

import (
	"rz-server/internal/common/errors/application_error"
)

type UserError struct {
	application_error.Manager
}

func New() *UserError {
	applicationErrorManager := application_error.NewManager("user", "user")

	return &UserError{
		*applicationErrorManager,
	}
}
