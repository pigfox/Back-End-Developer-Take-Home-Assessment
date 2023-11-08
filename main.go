package main

import (
	"bedtha/config"
	"bedtha/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	defer recovery()
	dbSetUp()
	loadJWTKEY()

	port := config.PortNumber
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.Handle("/task", handlers.ValidateToken(http.HandlerFunc(handlers.Create))).Methods("POST")
	router.Handle("/task/{task}", handlers.ValidateToken(http.HandlerFunc(handlers.Update))).Methods("PUT")
	router.Handle("/task/{task}", handlers.ValidateToken(http.HandlerFunc(handlers.Delete))).Methods("DELETE")
	router.Handle("/task/{task}", handlers.ValidateToken(http.HandlerFunc(handlers.View))).Methods("GET")
	router.Handle("/tasks/{page}", handlers.ValidateToken(http.HandlerFunc(handlers.List))).Methods("GET")

	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered. Error:\n", r)
	}
}
