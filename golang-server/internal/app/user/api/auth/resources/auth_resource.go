package auth_resources

import (
	auth_data "rz-server/internal/app/user/application/auth/data"
)

type AuthResourceData struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AuthMapper struct {
	data *auth_data.AuthData
}

func NewAuthMapper(data *auth_data.AuthData) *AuthMapper {
	return &AuthMapper{data}
}

func (m *AuthMapper) ToResource() AuthResourceData {
	return AuthResourceData{
		RefreshToken: m.data.RefreshToken,
		AccessToken:  m.data.AccessToken,
	}
}
