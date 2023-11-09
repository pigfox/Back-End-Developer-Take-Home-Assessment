package structs

import "github.com/dgrijalva/jwt-go"

var JwtKey JwtKeyObj

type JwtKeyObj struct {
	Value string `json:"value"`
}

type JwtTokenObj struct {
	Token string `json:"token"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
