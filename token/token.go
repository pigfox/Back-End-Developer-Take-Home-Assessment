package token

import (
	"bedtha/config"
	"bedtha/structs"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(structs.JwtKey.Value))
	return tokenStr, err
}
