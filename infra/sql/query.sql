-- name: GetWish :one
SELECT * FROM wishes
WHERE id = $1 LIMIT 1;

-- name: GetWishes :many
SELECT * FROM wishes
WHERE couple_id = $1
ORDER BY created_at;

-- name: CreateWish :one
INSERT INTO wishes (
  title, description, url, price, completed, couple_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: UpdateWish :execresult
UPDATE wishes
SET title = $2, description = $3, url = $4, price = $5, completed = $6
WHERE id = $1;

-- name: CompleteWish :execresult
UPDATE wishes
SET completed = true
WHERE id = $1;

-- name: DeleteWish :exec
DELETE FROM wishes
WHERE id = $1;

-- name: GetUsers :many
SELECT users.id, users.name, users.username, users.phone, couples.id as couple_id
FROM users
LEFT JOIN couples
ON couples.user_id = users.id
OR couples.partner_id = users.id;

-- name: GetUser :one
SELECT users.id, users.name, users.username, users.phone, couples.id as couple_id
FROM users
LEFT JOIN couples 
ON couples.user_id = users.id 
OR couples.partner_id = users.id
WHERE users.id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT users.id, users.name, users.username, users.phone, couples.id as couple_id
FROM users
LEFT JOIN couples
ON couples.user_id = users.id
OR couples.partner_id = users.id
WHERE users.username = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT users.id, users.name, users.username, users.phone, couples.id as couple_id
FROM users
LEFT JOIN couples
ON couples.user_id = users.id
OR couples.partner_id = users.id
WHERE users.phone = $1 LIMIT 1;

-- name: CheckUserName :one
SELECT count(users.id) FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserWithPassword :one
SELECT users.*, couples.id as couple_id FROM users
LEFT JOIN couples
ON couples.user_id = users.id
OR couples.partner_id = users.id
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  name, username, password, phone
) VALUES (
  $1, $2, $3, $4
)
RETURNING id;

-- name: UpdateUser :execresult
UPDATE users
SET name = $2
WHERE id = $1;

-- name: ChangePassword :execresult
UPDATE users
SET password = $2
WHERE id = $1;

-- name: GetPartnerName :one
SELECT users.username AS username
FROM users
INNER JOIN couples
ON couples.id = $1
WHERE users.id != $2
LIMIT 1;

-- name: CreateCouple :one
INSERT INTO couples (
  user_id, partner_id
) VALUES (
  $1, $2
)
RETURNING id;

-- name: DeleteCouple :exec
DELETE FROM couples
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
