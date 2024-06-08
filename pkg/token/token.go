package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type implToken struct{}

type Token interface {
	CreateNewToken(username string) (string, error)
}

func New() Token {
	return new(implToken)
}

func (e implToken) CreateNewToken(email string) (string, error) {

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokeToString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokeToString, err
}

func (e implToken) ValidateToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return tokenString, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
