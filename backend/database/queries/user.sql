-- name: CreateUser :one
INSERT INTO "User" (username) VALUES ($1) RETURNING *;

-- name: GetUser :one
SELECT * FROM "User" WHERE username = $1;