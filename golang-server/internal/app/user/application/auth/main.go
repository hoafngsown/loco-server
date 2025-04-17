package auth_service

import (
	application "rz-server/internal/app/user/application"
	auth_errors "rz-server/internal/app/user/application/auth/errors"
	domain "rz-server/internal/app/user/domain"
	store "rz-server/internal/app/user/infra/store"
	"rz-server/internal/common/interfaces"
)

type AuthService struct {
	authStore store.AuthStore
	userStore store.UserStore
	auth      domain.Auth
	errors    interfaces.ApplicationErrorManager
}

var _ application.AuthService = (*AuthService)(nil)

func New(authStore store.AuthStore, userStore store.UserStore, auth domain.Auth) *AuthService {
	application_error := auth_errors.New()
	application_error.RegisterAllErrors()

	return &AuthService{authStore, userStore, auth, application_error}
}
