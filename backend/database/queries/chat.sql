--name CreateChat :one
INSERT INTO Chat (username) VALUES ($1);

--name GetUserChat :many
SELECT * FROM Chat WHERE username = $1;