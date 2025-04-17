package auth_api

import (
	"net/http"
	auth_resources "rz-server/internal/app/user/api/auth/resources"
	"rz-server/internal/app/user/application"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	json_helper "rz-server/internal/common/helpers/json"
	"rz-server/internal/common/interfaces"
)

type AuthApi struct {
	service application.AuthService
	util    *interfaces.Util
	server  interfaces.Server
}

func New(
	server interfaces.Server,
	service application.AuthService,
	util *interfaces.Util,
) *AuthApi {
	u := new(AuthApi)
	u.service = service
	u.util = util
	u.server = server

	return u
}

func (u *AuthApi) Register() {
	u.server.POST("/user/auth/register", u.register)
	u.server.POST("/user/auth/login", u.login)
	u.server.POST("/user/auth/logout", u.logout)
	u.server.POST("/user/auth/refresh-token", u.refreshToken)
}

func (u *AuthApi) register(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Registering user", map[string]any{
		"method": r.Method,
	})

	command, _ := json_helper.ParseJson[auth_commands.RegisterUserCommand](r)

	authData, err := u.service.Register(command)

	if err != nil {
		json_helper.RespondJsonError(err, w)
		return
	}

	json_helper.RespondJsonResourceSuccess(auth_resources.NewAuthMapper(authData), w)
}

func (u *AuthApi) login(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Logging in user", map[string]any{
		"method": r.Method,
	})

	command, _ := json_helper.ParseJson[auth_commands.LoginUserCommand](r)

	authData, err := u.service.Login(command)

	if err != nil {
		json_helper.RespondJsonError(err, w)
		return
	}

	json_helper.RespondJsonResourceSuccess(auth_resources.NewAuthMapper(authData), w)
}

func (u *AuthApi) logout(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Logging out user", map[string]any{
		"method": r.Method,
	})

	command, _ := json_helper.ParseJson[auth_commands.LogoutUserCommand](r)

	err := u.service.Logout(command)

	if err != nil {
		json_helper.RespondJsonError(err, w)
		return
	}

	json_helper.RespondSuccess(w)
}

func (u *AuthApi) refreshToken(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Refreshing token", map[string]any{
		"method": r.Method,
	})

	command, _ := json_helper.ParseJson[auth_commands.RefreshTokenCommand](r)

	authData, err := u.service.RefreshToken(command)

	if err != nil {
		json_helper.RespondJsonError(err, w)
		return
	}

	json_helper.RespondJsonResourceSuccess(auth_resources.NewAuthMapper(authData), w)
}
