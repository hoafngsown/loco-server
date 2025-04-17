package auth_service

import (
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	auth_errors "rz-server/internal/app/user/application/auth/errors"
	"rz-server/internal/common/errors/application_error"
	"rz-server/internal/common/interfaces"
)

func (s *AuthService) RefreshToken(command auth_commands.RefreshTokenCommand) (*auth_data.AuthData, interfaces.ApplicationError) {
	refreshToken := command.RefreshToken

	refreshTokenData := s.authStore.GetRefreshTokenByToken(refreshToken)

	if refreshTokenData == nil {
		return nil, s.errors.New(auth_errors.REFRESH_TOKEN_NOT_FOUND)
	}

	err := s.auth.ValidateExpired(refreshTokenData.ExpireAt)

	if err != nil {
		return nil, s.errors.New(auth_errors.REFRESH_TOKEN_EXPIRED)
	}

	err = s.authStore.UpdateRefreshTokenExpiredAt(refreshTokenData.ID, s.auth.GetExpiredAtAfter())

	if err != nil {
		return nil, s.errors.New(application_error.STORE_SQL_ERROR)
	}

	accessToken, err := s.auth.GenerateAccessToken(refreshTokenData.UserID)

	if err != nil {
		return nil, s.errors.New(auth_errors.TOKEN_GENERATION_FAILED)
	}

	return &auth_data.AuthData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
