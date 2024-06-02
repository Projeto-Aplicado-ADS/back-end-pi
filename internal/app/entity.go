package app

import (
	"time"

	"projeto-api/internal/app/store"

	"github.com/ServiceWeaver/weaver"
)

type UsersOut struct {
	weaver.AutoMarshal
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type AllUsersOut []UsersOut

func (e UsersOut) FromStore(in store.User) UsersOut {
	e.ID = int(in.ID)
	e.FullName = in.FullName
	e.Email = in.Email
	e.Phone = in.Phone
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
	params.ID = int32(t.UTC().UnixNano()) // fake id
	params.FullName = e.FullName
	params.Email = e.Email
	params.Phone = e.Phone
	params.Password = e.Password // TODO emcript senha
	params.Birthday = e.Birthday.UnixMilli()
	params.CreatedAt = t.UnixMilli()
	return params
}
