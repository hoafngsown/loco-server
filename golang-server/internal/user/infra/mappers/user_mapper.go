package mapper

import (
	"go-server/internal/user/domain"
	store "go-server/internal/user/infra"
)

func ToDBUser(user domain.User) store.User {
	return store.User{
		ID:   1,
		Name: user.Name,
	}
}

func FromDBUser(user store.User) domain.User {
	return domain.User{
		Id:   user.ID,
		Name: user.Name,
	}
}
