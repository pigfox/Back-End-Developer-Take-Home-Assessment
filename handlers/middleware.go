package handlers

import (
	"bedtha/structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Token")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Token missing")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(structs.JwtKey.Value), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Token is invalid")
			return
		}

		// Token is valid; proceed to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
