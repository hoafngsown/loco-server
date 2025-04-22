-- name: MarkSetupCompleted :exec
UPDATE users
SET 
  preference = $2,
  has_completed_setup = true,
  updated_at = now()
WHERE id = $1;