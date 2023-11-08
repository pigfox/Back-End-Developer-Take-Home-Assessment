package main

import (
	"bedtha/db"
	"bedtha/structs"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func dbSetUp() {
	err := db.CreateTasks()
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateUsers()
	if err != nil {
		log.Fatal(err)
	}
}

func loadJWTKEY() {
	godotenv.Load(".env")
	structs.JwtKey.Value = os.Getenv("JWT_KEY")

	if structs.JwtKey.Value != "" {
		fmt.Println("JWT_KEY set")
	} else {
		log.Fatal("Quitting: JWT_KEY key Not Set")
	}
}
