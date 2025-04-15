package application

import (
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	"rz-server/internal/common/errors/application_error"
)

type AuthService interface {
	Register(command auth_commands.RegisterUserCommand) (*auth_data.AuthData, *application_error.Error)
	Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, *application_error.Error)
	RefreshToken(command auth_commands.RefreshTokenCommand) (*auth_data.AuthData, *application_error.Error)
	Logout(command auth_commands.LogoutUserCommand) *application_error.Error
}
