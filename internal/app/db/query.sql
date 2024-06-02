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


-- name: Login :execresult
SELECT email, password FROM users 
WHERE email = ? and password = ?;

