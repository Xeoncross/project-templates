
-- name: InsertUser :execlastid
INSERT INTO users (name, email) VALUES (?, ?);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: GetUsers :many
SELECT * FROM users;