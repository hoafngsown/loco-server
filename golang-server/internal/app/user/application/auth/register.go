package auth_service

import (
	"fmt"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	auth_errors "rz-server/internal/app/user/application/auth/errors"
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
	user_store_data "rz-server/internal/app/user/infra/store/sql/user/data"
	"rz-server/internal/common/errors/application_error"
	"rz-server/internal/common/interfaces"
)

func (s *AuthService) Register(command auth_commands.RegisterUserCommand) (*auth_data.AuthData, interfaces.ApplicationError) {
	user := s.userStore.GetUserByEmail(command.Email)

	fmt.Println("HH", user)

	if user != nil {
		return nil, s.errors.New(auth_errors.USER_ALREADY_EXISTS)
	}

	hashedPassword, err := s.auth.HashPassword(command.Password)

	if err != nil {
		return nil, s.errors.New(auth_errors.HASH_PASSWORD_FAILED)
	}

	newUser := s.userStore.CreateUser(user_store_data.CreateUserBody{
		Email:       command.Email,
		Password:    hashedPassword,
		DisplayName: command.DisplayName,
	})

	if newUser == nil {
		return nil, s.errors.New(application_error.STORE_SQL_ERROR)
	}

	refreshToken, expiredAt, err := s.auth.GenerateRefreshToken(newUser.Id)

	if err != nil {
		return nil, s.errors.New(auth_errors.TOKEN_GENERATION_FAILED)
	}

	s.authStore.SaveRefreshToken(auth_store_data.RefreshTokenBody{
		UserID:    newUser.Id,
		Token:     refreshToken,
		ExpiredAt: expiredAt,
	})

	accessToken, err := s.auth.GenerateAccessToken(newUser.Id)

	if err != nil {
		return nil, s.errors.New(auth_errors.TOKEN_GENERATION_FAILED)
	}

	return &auth_data.AuthData{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
