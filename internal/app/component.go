package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ServiceWeaver/weaver"

	"projeto-api/internal/app/store"
	"projeto-api/pkg/crypt"
	"projeto-api/pkg/token"

	_ "github.com/go-sql-driver/mysql"
)

type Component interface {
	AllUsers(ctx context.Context) (out AllUsersOut, err error)
	GetOneUserById(ctx context.Context, id string) (out UsersOut, err error)
	CreateUser(ctx context.Context, in UserIn) (ok bool, err error)
	GetUserByEmailAndPassword(ctx context.Context, email string, password string) (err error)
	ListUserByEmail(ctx context.Context, email string) (out UsersOut, err error)
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

func (e implapp) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (err error) {

	result, err := e.db.GetUserByEmail(ctx, email)
	if err != nil {
		return errors.New("Email ou senha invalido!")
	}

	ok := crypt.New().ValidateUserByPassword(password, result.Password)
	if !ok {
		return errors.New("Email ou senha invalido!")
	}

	tokenString, err := token.New().CreateNewToken(result.Email, result.IsAdmin)
	if err != nil {
		return err
	}
	fmt.Println(tokenString)


	return nil
}

func (e implapp) ListUserByEmail(ctx context.Context, email string) (out UsersOut, err error) {
	users, err := e.db.ListUserByEmail(ctx, email)
	if err != nil {
		return out, err
	}

	return out.FromStore(users), nil
}
