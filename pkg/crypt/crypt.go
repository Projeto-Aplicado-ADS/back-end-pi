package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

type implCrypt struct{}

type Crypt interface {
	CryptPassword(password string) (string, error)
	ValidateUserByPassword(password, hash string) bool
}

func New() Crypt {
	return new(implCrypt)
}

func (e implCrypt) CryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}


func (e implCrypt) ValidateUserByPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
