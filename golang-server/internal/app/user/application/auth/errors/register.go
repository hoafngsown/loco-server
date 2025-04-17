package auth_errors

func (e *AuthError) RegisterAllErrors() {
	e.RegisterBusinessError(HASH_PASSWORD_FAILED, "hash password failed")
	e.RegisterBusinessError(TOKEN_GENERATION_FAILED, "token generation failed")
	e.RegisterBusinessError(REFRESH_TOKEN_EXPIRED, "refresh token expired")

	e.RegisterNotFoundError(USER_NOT_FOUND, "user not found")
	e.RegisterNotFoundError(REFRESH_TOKEN_NOT_FOUND, "refresh token not found")

	e.RegisterInvalidInputError(INVALID_PASSWORD, "invalid password")
	e.RegisterInvalidInputError(USER_ALREADY_EXISTS, "user already exists")
}
