package handlers

import (
	"bedtha/config"
	"bedtha/db"
	"bedtha/structs"
	"bedtha/token"
	"bedtha/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Register(w http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
	}
	var user structs.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
	}

	if !utils.IsValidEmail(user.Email) {
		data := structs.Response{Data: "Invalid email"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(data)
		return
	}

	min := config.MinPasswordLength
	max := config.MaxPasswordLength
	length := len(user.Password)
	if length < min || max < length {
		data := structs.Response{Data: "Password length must be min " + strconv.Itoa(min) + " and max " + strconv.Itoa(max)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(data)
		return
	}

	affectedRows, err := db.Register(user)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	if affectedRows == 1 {
		w.WriteHeader(201)
	}
}

func Login(w http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
	}
	var user structs.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
	}

	code, err := db.Login(user)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(data)
	}

	tokenString, err := token.New(user.Email)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	token := structs.JwtTokenObj{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(token)
}
