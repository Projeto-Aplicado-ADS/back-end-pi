package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

type implCrypt struct{}

type Crypt interface {
	CryptPassword(password string) string
}

func New() Crypt {
	return new(implCrypt)
}

func (e implCrypt) CryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(bytes)
	}

	return string(bytes)
}

/*
- Testar Codigo depois
func ValidateUserPassword(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
} */
