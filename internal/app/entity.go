package app

import (
	"time"

	"projeto-api/internal/app/store"
	"projeto-api/pkg/crypt"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type (
	UsersOut struct {
		weaver.AutoMarshal
		ID        string    `json:"id"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone,omitempty"`
		IsAdmin   bool      `json:"is_admin"`
		Birthday  time.Time `json:"birthday,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	}
	UpdateUsersEmailIn struct {
		weaver.AutoMarshal
		ID    string `json:"id"`
		Email string `json:"email"`
	}
	UpdateUsersPhoneIn struct {
		weaver.AutoMarshal
		ID    string `json:"id"`
		Phone string `json:"phone"`
	}
)


type (
	HospedesOut struct {
		weaver.AutoMarshal
		ID             string    `json:"id"`
		Nome           string    `json:"nome"`
		Email          string    `json:"email"`
		Telefone       string    `json:"telefone,omitempty"`
		Cpf            string    `json:"cpf"`
		DataNascimento string    `json:"data_nascimento"`
		Sexo           string    `json:"sexo"`
		CreatedAt      time.Time `json:"created_at,omitempty"`
	}
	HospedesIn struct {
		weaver.AutoMarshal
		ID             string    `json:"id"`
		Nome           string    `json:"nome"`
		Email          string    `json:"email"`
		Telefone       string    `json:"telefone,omitempty"`
		Cpf            string    `json:"cpf"`
		DataNascimento string    `json:"data_nascimento"`
		Sexo           string    `json:"sexo"`
		CreateAt       time.Time `json:"created_at,omitempty"`
	}
	RemoveHospedesIn struct {
		weaver.AutoMarshal
		ID string `json:"id" validate:"required, uuid"`
		DeletedAt time.Time `json:"deleted_at,omitempty"`
	}
	UpdateHospedesIn struct {
		weaver.AutoMarshal
		ID string `json:"id" validate:"required, uuid"`
		Nome           string    `json:"nome,omitempty"`
		Email          string    `json:"email,omitempty"`
		Telefone       string    `json:"telefone,omitempty"`
		Cpf            string    `json:"cpf,omitempty"`
		DataNascimento string    `json:"data_nascimento,omitempty"`
		Sexo           string    `json:"sexo,omitempty"`
		UpdateAt       time.Time `json:"update_at,omitempty"`
	}
	QuartosOut struct {
		weaver.AutoMarshal
		ID           string    `json:"id"`
		Descricao    string    `json:"descricao,omitempty"`
		TipoQuarto   string    `json:"tipo_quarto"`
		StatusQuarto string    `json:"status_quarto"`
		NumeroQuarto int32     `json:"numero_quarto"`
		NumeroAndar  int32     `json:"numero_andar"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
	}
	QuartosIn struct {
		weaver.AutoMarshal
		ID           string    `json:"id"`
		Descricao    string    `json:"descricao,omitempty"`
		TipoQuarto   string    `json:"tipo_quarto"`
		StatusQuarto string    `json:"status_quarto"`
		NumeroQuarto int32     `json:"numero_quarto"`
		NumeroAndar  int32     `json:"numero_andar"`
		CreateAt     time.Time `json:"created_at,omitempty"`
	}

	UpdateQuartosStatusIn struct {
		weaver.AutoMarshal
		ID string `json:"id"`
		StatusQuarto string `json:"status_quarto"`
	}

	ReservasOut struct {
		weaver.AutoMarshal
		ID            string    `json:"id"`
		Nome          string    `json:"nome"`
		Cpf           string    `json:"cpf"`
		TipoReserva   string    `json:"tipo_reserva"`
		DataReserva   string    `json:"data_reserva"`
		DataCheckIn   string    `json:"data_check_in"`
		DataCheckOut  string    `json:"data_check_out"`
		StatusReserva string    `json:"status_reserva"`
		ValorReserva  string    `json:"valor_reserva"`
		NumeroQuarto  int32     `json:"numero_quarto"`
		NumeroAndar   int32     `json:"numero_andar"`
		TipoQuarto    string    `json:"tipo_quarto"`
		StatusQuarto  string    `json:"status_quarto"`
		CreatedAt     time.Time `json:"created_at,omitempty"`
	}
	ReservasIn struct {
		weaver.AutoMarshal
		ID            string    `json:"id"`
		TipoReserva   string    `json:"tipo_reserva"`
		DataReserva   string    `json:"data_reserva"`
		DataCheckIn   string    `json:"data_check_in"`
		DataCheckOut  string    `json:"data_check_out"`
		StatusReserva string    `json:"status_reserva"`
		ValorReserva  string    `json:"valor_reserva"`
		QuartoId      string    `json:"quarto_id"`
		HospedeId     string    `json:"hospede_id"`
		CreateAt      time.Time `json:"created_at,omitempty"`
	}
)

type AllHospodesOut []HospedesOut

type AllQuartosOut []QuartosOut

type AllReservasOut []ReservasOut

type UserGetByEmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
}

type AllUsersOut []UsersOut

func (e UsersOut) FromStore(in store.User) UsersOut {
	e.ID = in.ID
	e.FullName = in.FullName
	e.Email = in.Email
	e.Phone = in.Phone
	e.IsAdmin = in.IsAdmin
	e.CreatedAt = time.UnixMilli(in.CreatedAt)
	return e
}

func (e AllUsersOut) FromStore(in []store.User) AllUsersOut {
	e = make(AllUsersOut, 0, len(in))
	for _, v := range in {
		e = append(e, UsersOut{}.FromStore(v))
	}
	return e
}

type UserIn struct {
	weaver.AutoMarshal
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone,omitempty"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday,omitempty"`
	/* UpdateAt time.Time `json:"update_at,omitempty"` */
}

func (e UserIn) ToStore() (params store.CreateUserParams) {
	t := time.Now()
	params.ID = uuid.New().String()
	params.FullName = e.FullName
	params.Email = e.Email
	params.Phone = e.Phone
	params.Password, _ = crypt.New().CryptPassword(e.Password)
	params.IsAdmin = false
	params.CreatedAt = t.UnixMilli()
	return params
}

func (e UserGetByEmailAndPassword) FromStoreByEmailAndPassword(in store.GetUserByEmailRow) UserGetByEmailAndPassword {
	e.Email = in.Email
	e.Password = in.Password
	e.IsAdmin = in.IsAdmin

	return e
}

func (e UpdateUsersEmailIn) ToStore() (params store.AlterarEmailUserParams) {
	params.ID = e.ID
	params.Email = e.Email
	return params
} 

func (e UpdateUsersPhoneIn) ToStore() (params store.AlterarPhoneUserParams) {
	params.ID = e.ID
	params.Phone = e.Phone
	return params
}

func (e HospedesOut) FromStore(in store.Hospede) HospedesOut {
	e.ID = in.ID
	e.Nome = in.Nome
	e.Email = in.Email
	e.Telefone = in.Telefone
	e.Cpf = in.Cpf
	e.DataNascimento = in.DataNascimento
	e.Sexo = string(in.Sexo)
	e.CreatedAt = time.UnixMilli(in.CreatedAt)
	return e
}

func (e AllHospodesOut) FromStore(in []store.Hospede) AllHospodesOut {
	e = make(AllHospodesOut, 0, len(in))
	for _, v := range in {
		e = append(e, HospedesOut{}.FromStore(v))
	}
	return e
}

func (e HospedesIn) ToStore() (params store.CreateHospedeParams) {
	t := time.Now()
	params.ID = uuid.New().String()
	params.Cpf = e.Cpf
	params.DataNascimento = e.DataNascimento
	params.Email = e.Email
	params.Nome = e.Nome
	params.Sexo = store.HospedesSexo(e.Sexo)
	params.Telefone = e.Telefone
	params.CreatedAt = t.UnixMilli()
	return params
}

func (e RemoveHospedesIn) ToStore() (out store.RemoverHospedeParams) {
	time := time.Now()
	out.ID = e.ID
	out.DeletedAt = time.UnixMilli()
	return out
}

func (e UpdateHospedesIn) ToStore() (params store.AlterarHospedeParams) {
	time := time.Now()
	params.ID = e.ID
	params.Cpf = e.Cpf
	params.DataNascimento = e.DataNascimento
	params.Email = e.Email
	params.Nome = e.Nome
	params.Sexo = store.HospedesSexo(e.Sexo)
	params.Telefone = e.Telefone
	params.UpdateAt = time.UnixMilli()
	return params
}


func (e QuartosOut) FromStore(in store.Quarto) QuartosOut {
	e.ID = in.ID
	e.Descricao = in.Descricao
	e.NumeroAndar = in.NumeroAndar
	e.NumeroQuarto = in.NumeroQuarto
	e.StatusQuarto = string(in.StatusQuarto)
	e.TipoQuarto = string(in.TipoQuarto)
	e.CreatedAt = time.UnixMilli(in.CreatedAt)
	return e
}

func (e AllQuartosOut) FromStore(in []store.Quarto) AllQuartosOut {
	e = make(AllQuartosOut, 0, len(in))
	for _, v := range in {
		e = append(e, QuartosOut{}.FromStore(v))
	}
	return e
}

func (e QuartosIn) ToStore() (params store.CreateQuartoParams) {
	t := time.Now()
	params.ID = uuid.New().String()
	params.Descricao = e.Descricao
	params.NumeroAndar = e.NumeroAndar
	params.NumeroQuarto = e.NumeroQuarto
	params.StatusQuarto = store.QuartosStatusQuarto(e.StatusQuarto)
	params.TipoQuarto = store.QuartosTipoQuarto(e.TipoQuarto)
	params.CreatedAt = t.UnixMilli()
	return params
}

func (e UpdateQuartosStatusIn) ToStore() (out store.AlterarStatusQuartoParams) {
	out.ID = e.ID
	out.StatusQuarto = store.QuartosStatusQuarto(e.StatusQuarto)
	return out
}


func (e ReservasOut) FromStore(in store.ListReservasRow) ReservasOut {
	e.ID = in.ID
	e.Cpf = in.Cpf
	e.Nome = in.Nome
	e.DataReserva = in.DataReserva
	e.DataCheckIn = in.DataCheckin
	e.DataCheckOut = in.DataCheckout
	e.ValorReserva = in.ValorReserva
	e.NumeroAndar = in.NumeroAndar
	e.NumeroQuarto = in.NumeroQuarto
	e.StatusQuarto = string(in.StatusQuarto)
	e.StatusReserva = string(in.StatusReserva)
	e.TipoQuarto = string(in.TipoQuarto)
	e.TipoReserva = string(in.TipoReserva)
	e.CreatedAt = time.UnixMilli(in.CreatedAt)
	return e
}

func (e AllReservasOut) FromStore(in []store.ListReservasRow) AllReservasOut {
	e = make(AllReservasOut, 0, len(in))
	for _, v := range in {
		e = append(e, ReservasOut{}.FromStore(v))
	}
	return e
}

func (e ReservasIn) ToStore() (params store.CreateReservaParams) {
	t := time.Now()
	params.ID = uuid.New().String()
	params.DataCheckin = e.DataCheckIn
	params.DataCheckout = e.DataCheckOut
	params.ValorReserva = e.ValorReserva
	params.StatusReserva = store.ReservasStatusReserva(e.StatusReserva)
	params.TipoReserva = store.ReservasTipoReserva(e.TipoReserva)
	params.IDQuarto = e.QuartoId
	params.IDHospede = e.HospedeId
	params.CreatedAt = t.UnixMilli()
	return params
}
