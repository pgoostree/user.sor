package main

import (
	"database/sql"
	"log"
	"net/http"

	"user.sor/controllers"

	"github.com/gorilla/mux"
	"user.sor/driver"
)

var db *sql.DB

func main() {
	db = driver.Database()

	router := mux.NewRouter()

	userController := controllers.UserController{}
	router.HandleFunc("/users", userController.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{userID}", userController.GetUser(db)).Methods("GET")
	router.HandleFunc("/users/{userID}", userController.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{userID}", userController.DeleteUser(db)).Methods("DELETE")

	groupController := controllers.GroupController{}
	router.HandleFunc("/groups", groupController.CreateGroup(db)).Methods("POST")
	router.HandleFunc("/groups/{groupName}", groupController.GetGroup(db)).Methods("GET")
	router.HandleFunc("/groups/{groupName}", groupController.UpdateGroup(db)).Methods("PUT")
	router.HandleFunc("/groups/{groupName}", groupController.DeleteGroup(db)).Methods("DELETE")

	log.Println("Listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
