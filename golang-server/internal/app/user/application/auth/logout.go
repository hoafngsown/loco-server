package auth_service

import (
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_errors "rz-server/internal/app/user/application/auth/errors"
	"rz-server/internal/common/errors/application_error"
	"rz-server/internal/common/interfaces"
)

func (s *AuthService) Logout(command auth_commands.LogoutUserCommand) interfaces.ApplicationError {
	userID := command.UserID

	existingRefreshToken := s.authStore.GetRefreshTokenByUserID(userID)

	if existingRefreshToken == nil {
		return s.errors.New(auth_errors.REFRESH_TOKEN_NOT_FOUND)
	}

	err := s.authStore.DeleteRefreshTokenByUserID(userID)

	if err != nil {
		return s.errors.New(application_error.STORE_SQL_ERROR)
	}

	return nil
}
