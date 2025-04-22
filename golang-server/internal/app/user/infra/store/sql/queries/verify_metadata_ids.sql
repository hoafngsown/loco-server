-- name: VerifyMetadataIDs :one
SELECT NOT EXISTS (
  SELECT 1
  FROM unnest($1::uuid[]) AS input_id
  WHERE NOT EXISTS (
    SELECT 1
    FROM preference_metadata pm
    WHERE pm.id = input_id
  )
) AS all_ids_valid;
