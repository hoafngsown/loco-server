package user_service

import (
	user_commands "rz-server/internal/app/user/application/user/commands"
	user_errors "rz-server/internal/app/user/application/user/errors"
	"rz-server/internal/common/errors/application_error"
	common_helper "rz-server/internal/common/helpers/common"
	json_helper "rz-server/internal/common/helpers/json"
	"rz-server/internal/common/interfaces"

	"github.com/google/uuid"
)

func (u *UserService) SetUpComplete(body user_commands.SetUpCompleteBody) (bool, interfaces.ApplicationError) {

	// 2. verify all ids is in metadata ids
	// 3. if at least one id is not in metadata ids, return error
	// 4. convert preference to jsonb
	// 5. save to user_preference field (update hasUserCompletedSetup to true)
	// 6. return success

	preferences := body.Preferences
	vibeIds := body.Preferences.VibeIds
	styleIds := body.Preferences.StyleIds

	if len(vibeIds) == 0 || len(styleIds) == 0 {
		return false, u.errors.New(user_errors.VIBE_IDS_OR_STYLE_IDS_REQUIRED)
	}

	// 1. filter duplicate ids
	vibeIds = common_helper.RemoveDuplicate(vibeIds)
	styleIds = common_helper.RemoveDuplicate(styleIds)

	// 2. verify all ids is in metadata ids
	valid := u.metadataStore.VerifyMetadataIDs(
		append(vibeIds, styleIds...),
	)

	if !valid {
		return false, u.errors.New(user_errors.INVALID_METADATA_IDS)
	}

	// 3. Parse preferences to jsonb
	preferencesJSON := json_helper.ToJSON(preferences)

	// 4. Update user preferences and set hasUserCompletedSetup to true
	err := u.userStore.MarkSetupCompleted(uuid.UUID{}, preferencesJSON)

	if err != nil {
		return false, u.errors.New(application_error.STORE_SQL_ERROR)
	}

	return true, nil

}
