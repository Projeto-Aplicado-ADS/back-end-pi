package app

import (
	"time"

	"projeto-api/internal/app/store"
	"projeto-api/pkg/crypt"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type UsersOut struct {
	weaver.AutoMarshal
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	IsAdmin  	bool      `json:"is_admin,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

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
	e.Birthday = time.UnixMilli(in.Birthday)
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
	params.Birthday = e.Birthday.Unix()
	params.CreatedAt = t.UnixMilli()
	return params
}

func (e UserGetByEmailAndPassword) FromStoreByEmailAndPassword(in store.GetUserByEmailRow) UserGetByEmailAndPassword {
	e.Email = in.Email
	e.Password = in.Password
	e.IsAdmin = in.IsAdmin

	return e
}
