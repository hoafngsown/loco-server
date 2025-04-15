package auth_errors

func (e *AuthError) RegisterAllErrors() {
	e.RegisterNotFoundError(USER_NOT_FOUND, "user not found")
	e.RegisterInvalidInputError(INVALID_PASSWORD, "invalid password")
	e.RegisterNotFoundError(REFRESH_TOKEN_NOT_FOUND, "refresh token not found")
	e.RegisterBusinessError(REFRESH_TOKEN_EXPIRED, "refresh token expired")
	e.RegisterInvalidInputError(USER_ALREADY_EXISTS, "user already exists")
}
