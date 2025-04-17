package auth_service

import (
	"fmt"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	auth_errors "rz-server/internal/app/user/application/auth/errors"
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
	"rz-server/internal/common/interfaces"
)

func (s *AuthService) Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, interfaces.ApplicationError) {
	email := command.Email
	password := command.Password

	user := s.userStore.GetUserByEmail(email)

	fmt.Printf("user: %+v\n", user)

	if user == nil {
		return nil, s.errors.New(auth_errors.USER_NOT_FOUND)
	}

	isPasswordCorrect := s.auth.ComparePassword(password, user.Password)

	if !isPasswordCorrect {
		return nil, s.errors.New(auth_errors.INVALID_PASSWORD)
	}

	existingRefreshToken := s.authStore.GetRefreshTokenByUserID(user.Id)

	refreshToken := ""

	if existingRefreshToken != nil {
		s.authStore.UpdateRefreshTokenExpiredAt(existingRefreshToken.ID, s.auth.GetExpiredAtAfter())
		refreshToken = existingRefreshToken.Token
	} else {
		newRefreshToken, expiredAt, err := s.auth.GenerateRefreshToken(user.Id)

		if err != nil {
			return nil, s.errors.New(auth_errors.TOKEN_GENERATION_FAILED)
		}

		s.authStore.SaveRefreshToken(auth_store_data.RefreshTokenBody{
			UserID:    user.Id,
			Token:     newRefreshToken,
			ExpiredAt: expiredAt,
		})
		refreshToken = newRefreshToken
	}

	accessToken, err := s.auth.GenerateAccessToken(user.Id)

	if err != nil {
		return nil, s.errors.New(auth_errors.TOKEN_GENERATION_FAILED)
	}

	return &auth_data.AuthData{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
