package user_api

import (
	"rz-server/internal/app/user/application"
	"rz-server/internal/common/interfaces"
)

type UserApi struct {
	service application.UserService
	util    *interfaces.Util
	*UserRouting
}

func New(
	server interfaces.Server,
	service application.UserService,
	util *interfaces.Util,
) *UserApi {
	u := new(UserApi)
	u.service = service
	u.util = util
	u.UserRouting = NewRoutingSetup(server, util, service)

	return u
}

func (u *UserApi) Register() {
}
