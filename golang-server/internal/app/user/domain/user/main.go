package user

import (
	domain "rz-server/internal/app/user/domain"
)

type Entity struct {
}

var _ domain.User = (*Entity)(nil)

func New() *Entity {
	return &Entity{}
}
