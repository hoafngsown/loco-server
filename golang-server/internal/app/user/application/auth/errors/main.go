package auth_errors

import (
	"rz-server/internal/common/errors/application_error"
)

type AuthError struct {
	application_error.Manager
}

func New() *AuthError {
	applicationErrorManager := application_error.NewManager("user", "auth")

	return &AuthError{
		*applicationErrorManager,
	}
}
