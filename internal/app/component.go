package app

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ServiceWeaver/weaver"

	"projeto-api/internal/app/store"

    _ "github.com/go-sql-driver/mysql"
)

type Component interface {
	AllUsers(ctx context.Context) (out AllUsersOut, err error)
	GetOneUserById(ctx context.Context, id string) (out UsersOut, err error)
	CreateUser(ctx context.Context, in UserIn) (ok bool, err error)
}

type Config struct {
	Driver string
	Source string
}

type implapp struct {
	weaver.Implements[Component]
	weaver.WithConfig[Config]
	db *store.Queries
}

func (e *implapp) Init(ctx context.Context) error {
	db, err := sql.Open(e.Config().Driver, e.Config().Source)
	if err != nil {
		return fmt.Errorf("not open: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping: %w", err)
	}

	e.db = store.New(db)
	return nil
}

func (e implapp) AllUsers(ctx context.Context) (out AllUsersOut, err error) {
	users, err := e.db.ListUsers(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(users), nil
}

func (e implapp) GetOneUserById(ctx context.Context, id string) (out UsersOut, err error) {
	user, err := e.db.GetUserById(ctx, id)
	if err != nil {
		return out, err
	}
	return out.FromStore(user), nil
}

func (e implapp) CreateUser(ctx context.Context, in UserIn) (ok bool, err error) {
	_, err = e.db.CreateUser(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}