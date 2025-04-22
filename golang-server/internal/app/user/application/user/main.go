package user_service

import (
	application "rz-server/internal/app/user/application"
	user_errors "rz-server/internal/app/user/application/user/errors"
	domain "rz-server/internal/app/user/domain"
	store "rz-server/internal/app/user/infra/store"
	"rz-server/internal/common/interfaces"
)

type UserService struct {
	userStore     store.UserStore
	metadataStore store.PreferenceMetadataStore
	user          domain.User
	errors        interfaces.ApplicationErrorManager
}

var _ application.UserService = (*UserService)(nil)

func New(userStore store.UserStore, metadataStore store.PreferenceMetadataStore, user domain.User) *UserService {
	application_error := user_errors.New()
	application_error.RegisterAllErrors()

	return &UserService{userStore, metadataStore, user, application_error}
}
