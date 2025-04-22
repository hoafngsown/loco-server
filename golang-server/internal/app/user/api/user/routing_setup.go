package user_api

import (
	"rz-server/internal/app/user/application"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/middlewares"
)

type UserRouting struct {
	publicRoute  interfaces.Route
	privateRoute interfaces.Route
}

func NewRoutingSetup(
	server interfaces.Server,
	util *interfaces.Util,
	service application.UserService,
) *UserRouting {
	publicRoute := server.NewRoute()
	publicRoute.SetPathPrefix("/api/user")
	publicRoute.Use(middlewares.NewLoggingMiddleware(util.Log))

	privateRoute := server.NewRoute()
	privateRoute.Use(
		middlewares.NewLoggingMiddleware(util.Log),
		middlewares.NewJWTAuthorizationMiddleware(),
	)

	return &UserRouting{
		publicRoute:  publicRoute,
		privateRoute: privateRoute,
	}
}
