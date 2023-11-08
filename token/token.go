package token

import (
	"bedtha/config"
	"bedtha/structs"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func New(email string) (string, error) {
	expirationTime := time.Now().Add(config.TokenExpiration * time.Minute)
	claims := &structs.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println(structs.JwtKey.Value)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(structs.JwtKey.Value))
	return tokenStr, err
}

func IsValid(token string) (bool, error) {
	claims := &structs.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(structs.JwtKey.Value), nil
	})

	if err != nil {
		return false, err
	}

	if tkn == nil || !tkn.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
