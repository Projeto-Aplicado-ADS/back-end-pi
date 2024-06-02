// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  id, full_name, email, phone, password, birthday, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
`

type CreateUserParams struct {
	ID        int32
	FullName  string
	Email     string
	Phone     string
	Password  string
	Birthday  int64
	CreatedAt int64
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Phone,
		arg.Password,
		arg.Birthday,
		arg.CreatedAt,
	)
}

const getUserById = `-- name: GetUserById :one
SELECT id, full_name, email, phone, password, is_admin, birthday, created_at, update_at FROM users
WHERE id = ?
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.IsAdmin,
		&i.Birthday,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const listAdmins = `-- name: ListAdmins :many
SELECT id, full_name, email, phone, password, is_admin, birthday, created_at, update_at FROM users
WHERE is_admin = 1
`

func (q *Queries) ListAdmins(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listAdmins)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.Phone,
			&i.Password,
			&i.IsAdmin,
			&i.Birthday,
			&i.CreatedAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, full_name, email, phone, password, is_admin, birthday, created_at, update_at FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.Phone,
			&i.Password,
			&i.IsAdmin,
			&i.Birthday,
			&i.CreatedAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const login = `-- name: Login :execresult
SELECT email, password FROM users 
WHERE email = ? and password = ?
`

type LoginParams struct {
	Email    string
	Password string
}

func (q *Queries) Login(ctx context.Context, arg LoginParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, login, arg.Email, arg.Password)
}