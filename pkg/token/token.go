package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type implToken struct{}

type Token interface {
	CreateNewToken(email string, isAdming bool) (string, error)
	ValidateToken(token string) (error)
}

type Claims struct {
	Email string `json:"email"`
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

func New() Token {
	return new(implToken)
}

var jwtKey = []byte("secret_key") 

func (e implToken) CreateNewToken(email string, isAdming bool) (string, error) {

	var claims = &Claims{
		Email: email,
		IsAdmin: isAdming,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenToString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenToString, err
}

func (e implToken) ValidateToken(token string) (error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}

		if !tkn.Valid {
			return errors.New("invalid token")
		}

		if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
			return errors.New("token expired")
		}

		return err
	}

	return nil
}
