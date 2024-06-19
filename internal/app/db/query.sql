-- name: ListUsers :many
SELECT * FROM users;

-- name: ListAdmins :many
SELECT * FROM users
WHERE is_admin = 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (
  id, full_name, email, phone, password, birthday, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);


-- name: GetUserByEmail :one
SELECT email, password, is_admin FROM users 
WHERE email = ?;

-- name: ListUserByEmail :one
SELECT * FROM users
WHERE email = ?;


