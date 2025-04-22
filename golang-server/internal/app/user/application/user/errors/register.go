package user_errors

func (e *UserError) RegisterAllErrors() {
	e.RegisterValidationError(VIBE_IDS_OR_STYLE_IDS_REQUIRED, "Vibe ids or style ids are required")
	e.RegisterValidationError(INVALID_METADATA_IDS, "Invalid metadata ids")
}
