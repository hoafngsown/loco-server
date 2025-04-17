package application

import (
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	"rz-server/internal/common/interfaces"
)

type AuthService interface {
	Register(command auth_commands.RegisterUserCommand) (*auth_data.AuthData, interfaces.ApplicationError)
	Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, interfaces.ApplicationError)
	RefreshToken(command auth_commands.RefreshTokenCommand) (*auth_data.AuthData, interfaces.ApplicationError)
	Logout(command auth_commands.LogoutUserCommand) interfaces.ApplicationError
}
