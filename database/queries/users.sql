-- name: GetUser :one
SELECT * FROM obiwan.users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM obiwan.users
WHERE deletedAt IS NULL;

-- name: CreateUser :one
INSERT INTO obiwan.users (
  email, password, firstName, lastName, dob
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateUser :one
UPDATE obiwan.users
  SET firstName = CASE WHEN $2 IS NOT NULL THEN $2 ELSE firstName END,
  lastName = CASE WHEN $3 IS NOT NULL THEN $3 ELSE lastName END,
  dob = CASE WHEN $4 IS NOT NULL THEN $4 ELSE dob END,
  updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
UPDATE obiwan.users
  SET deletedAt = CURRENT_TIMESTAMP
WHERE id = $1 AND deletedAt IS NULL;
