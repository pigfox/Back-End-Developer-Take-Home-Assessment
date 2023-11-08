package token

import (
	"bedtha/config"
	"bedtha/structs"
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
