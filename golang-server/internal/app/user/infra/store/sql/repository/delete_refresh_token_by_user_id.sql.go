// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: delete_refresh_token_by_user_id.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const deleteRefreshTokenByUserID = `-- name: DeleteRefreshTokenByUserID :exec
DELETE FROM refresh_token
WHERE user_id = $1
`

func (q *Queries) DeleteRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteRefreshTokenByUserID, userID)
	return err
}
