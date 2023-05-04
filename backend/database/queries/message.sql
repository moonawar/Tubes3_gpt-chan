-- name: CreateMessage :one
INSERT INTO "Message" (chat_id, question, answer) VALUES ($1, $2, $3) RETURNING *;

-- name: GetChatMessages :many
SELECT * FROM "Message" WHERE chat_id = $1 ORDER BY "no" ASC LIMIT $2 OFFSET $3;