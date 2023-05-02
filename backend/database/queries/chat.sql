-- name: CreateChat :one
INSERT INTO "Chat" (username) VALUES ($1) RETURNING *;

-- name: GetUserChat :many
SELECT "chat_id" FROM "Chat" WHERE username = $1;