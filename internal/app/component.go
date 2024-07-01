package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ServiceWeaver/weaver"

	"projeto-api/internal/app/store"
	"projeto-api/pkg/crypt"

	_ "github.com/go-sql-driver/mysql"
)

type Component interface {
	AllUsers(ctx context.Context) (out AllUsersOut, err error)
	GetOneUserById(ctx context.Context, id string) (out UsersOut, err error)
	CreateUser(ctx context.Context, in UserIn) (ok bool, err error)
	GetUserByEmailAndPassword(ctx context.Context, email string, password string) (err error)
	ListUserByEmail(ctx context.Context, email string) (out UsersOut, err error)
	UpdateUserByEmail(ctx context.Context, in UpdateUsersEmailIn) (ok bool, err error)
	UpdateUserPhone(ctx context.Context, in UpdateUsersPhoneIn) (ok bool, err error)

	/* Hospedes */
	GetHospedes(ctx context.Context) (out AllHospodesOut, err error)
	CreateHospede(ctx context.Context, in HospedesIn) (ok bool, err error)
	RemoveHospede(ctx context.Context, out RemoveHospedesIn) (ok bool, err error)
	UpdateHospedesIn(ctx context.Context, in UpdateHospedesIn) (ok bool, err error)
	GetOneHospedeById(ctx context.Context, id string) (out HospedesOut, err error)

	/* Quartos */
	GetQuartos(ctx context.Context) (out AllQuartosOut, err error)
	CreateQuarto(ctx context.Context, in QuartosIn) (ok bool, err error)
	UpdateStatusQuarto(ctx context.Context, in UpdateQuartosStatusIn) (ok bool, err error)

	/* Reservas */

	GetReservas(ctx context.Context) (out AllReservasOut, err error) 
	CreateReserva (ctx context.Context, in ReservasIn) (ok bool, err error)
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

	return nil
}

func (e implapp) ListUserByEmail(ctx context.Context, email string) (out UsersOut, err error) {
	users, err := e.db.ListUserByEmail(ctx, email)
	if err != nil {
		return out, err
	}

	return out.FromStore(users), nil
}

func (e implapp) UpdateUserByEmail (ctx context.Context, in UpdateUsersEmailIn) (ok bool, err error) {
	_, err = e.db.AlterarEmailUser(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e implapp) UpdateUserPhone(ctx context.Context, in UpdateUsersPhoneIn) (ok bool, err error) {
	_, err = e.db.AlterarPhoneUser(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}


func (e implapp) GetHospedes(ctx context.Context) (out AllHospodesOut, err error) {
	hospedes, err := e.db.ListHospedes(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(hospedes), nil
}


func (e implapp) GetOneHospedeById(ctx context.Context, id string) (out HospedesOut, err error) {	
	hospede, err := e.db.ListHospedeById(ctx, id)
	if err != nil {
		return out, err
	}
	return out.FromStore(hospede), nil
}

func (e implapp) CreateHospede(ctx context.Context, in HospedesIn) (ok bool, err error) {
	_, err = e.db.CreateHospede(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e implapp) RemoveHospede(ctx context.Context, out RemoveHospedesIn) (ok bool, err error) {
	err = e.db.RemoverHospede(ctx, out.ToStore())
	if err != nil {
		return false, errors.New("Erro ao remover hospede")
	}
	return true, nil
}

func (e implapp) UpdateHospedesIn(ctx context.Context, in UpdateHospedesIn) (ok bool, err error) {
	_, err = e.db.AlterarHospede(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e implapp) GetQuartos(ctx context.Context) (out AllQuartosOut, err error) {
	quartos, err := e.db.ListQuartos(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(quartos), nil
} 

func (e implapp) CreateQuarto(ctx context.Context, in QuartosIn) (ok bool, err error) {
	_, err = e.db.CreateQuarto(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e implapp) UpdateStatusQuarto(ctx context.Context, in UpdateQuartosStatusIn) (ok bool, err error) {
	err = e.db.AlterarStatusQuarto(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}



/* Reservas */

func (e implapp) GetReservas(ctx context.Context) (out AllReservasOut, err error) {
	reservas, err := e.db.ListReservas(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(reservas), nil
}
 
func (e implapp) CreateReserva (ctx context.Context, in ReservasIn) (ok bool, err error) {
	_, err = e.db.CreateReserva(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}


// TODO: PUT AND UPDATE RESERVA


