package auth_api

import (
	"rz-server/internal/app/user/application"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/middlewares"
)

type AuthRouting struct {
	publicRoute  interfaces.Route
	privateRoute interfaces.Route
}

func NewRoutingSetup(
	server interfaces.Server,
	util *interfaces.Util,
	service application.AuthService,
) *AuthRouting {
	publicRoute := server.NewRoute()
	publicRoute.SetPathPrefix("/api/user/auth")
	publicRoute.Use(middlewares.NewLoggingMiddleware(util.Log))

	privateRoute := server.NewRoute()

	return &AuthRouting{
		publicRoute:  publicRoute,
		privateRoute: privateRoute,
	}
}
