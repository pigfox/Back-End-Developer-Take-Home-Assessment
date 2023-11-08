package db

import (
	"bedtha/structs"
	"bedtha/utils"
	"errors"
)

func Register(user structs.User) (int, error) {
	db, err := conn()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	statement, err := db.Prepare("INSERT INTO users (email, hash) VALUES (?,?)")
	if err != nil {
		return 0, err
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	res, err := statement.Exec(user.Email, hash)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func Login(user structs.User) (int, error) {
	db, err := conn()
	if err != nil {
		return 500, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, email, hash FROM users WHERE email=?", user.Email)
	if err != nil {
		return 500, err
	}
	var user2 structs.User
	var users []structs.User
	for rows.Next() {
		rows.Scan(&user2.ID, &user2.Email, &user2.Password)
		users = append(users, user2)
	}

	if !utils.CheckPasswordHash(user.Password, users[0].Password) {
		return 400, errors.New("invalid credentials")
	}

	return 200, nil
}
