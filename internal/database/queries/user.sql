-- name: CountUsersWithUsername :one
SELECT COUNT(id) FROM "user" WHERE username = $1;

-- name: CreateUser :one
INSERT INTO "user" (username, password) VALUES ($1, $2) RETURNING id;

-- name: GetUserById :one
SELECT * FROM "user" WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM "user" WHERE username = $1;
