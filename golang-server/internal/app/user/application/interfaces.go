package application

import (
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	user_commands "rz-server/internal/app/user/application/user/commands"
	"rz-server/internal/common/interfaces"
)

type AuthService interface {
	Register(command auth_commands.RegisterUserCommand) (*auth_data.AuthData, interfaces.ApplicationError)
	Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, interfaces.ApplicationError)
	RefreshToken(command auth_commands.RefreshTokenCommand) (*auth_data.AuthData, interfaces.ApplicationError)
	Logout(command auth_commands.LogoutUserCommand) interfaces.ApplicationError
}

type UserService interface {
	SetUpComplete(body user_commands.SetUpCompleteBody) (bool, interfaces.ApplicationError)
	// GetSetupState(id uuid.UUID)
	// GetPreferenceMetadata()
}
