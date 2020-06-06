package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5436
	user     = "postgres"
	password = "postgres"
	dbname   = "user_sor"
)

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Group struct {
	Name string `json:"name"`
}

type UserGroup struct {
	UserIDs []string `json:"user_ids"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{userId}", getUser).Methods("GET")
	router.HandleFunc("/users/{userId}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", deleteUser).Methods("DELETE")

	router.HandleFunc("/groups", createGroup).Methods("POST")
	router.HandleFunc("/groups/{groupName}", getGroup).Methods("GET")
	router.HandleFunc("/groups/{groupName}", updateGroup).Methods("PUT")
	router.HandleFunc("/groups/{groupName}", deleteGroup).Methods("DELETE")

	log.Println("Listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	var group Group

	json.NewDecoder(r.Body).Decode(&group)

	if group.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is required.")
		return
	}

	var result Group
	err := db.QueryRow("select * from users.insert_group('" + group.Name + "')").Scan(&result.Name)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "duplicate key") {
			respondWithError(w, http.StatusConflict, fmt.Sprintf("A group with Name=%s already exists", group.Name))
		}
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	groupName := mux.Vars(r)["groupName"]

	rows, err := db.Query("select * from users.get_user_group('" + groupName + "')")
	if err != nil {
		log.Print(err)
	}

	var userGroup UserGroup
	defer rows.Close()
	for rows.Next() {
		var userID string
		err = rows.Scan(&userID)
		if err != nil {
			log.Print(err)
			respondWithError(w, http.StatusInternalServerError, "An error occurred, please try again later.")
		}
		userGroup.UserIDs = append(userGroup.UserIDs, userID)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	if len(userGroup.UserIDs) > 0 {
		json.NewEncoder(w).Encode(userGroup)
	} else {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Either there is no group named %s or there are no users associated with it", groupName))
	}
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	var userGroup UserGroup

	json.NewDecoder(r.Body).Decode(&userGroup)

	groupName := mux.Vars(r)["groupName"]

	_, err := db.Exec("DELETE FROM users.user_group")
	if err != nil {
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, "An error occurred, please try again later.")
	}

	_, err = db.Exec("INSERT INTO users.user_group (user_id, group_name) VALUES (unnest($1::text[]), $2)", pq.Array(userGroup.UserIDs), groupName)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "foreign key constraint") {
			respondWithError(w, http.StatusNotFound, "Invalid user or group name.")
		}
	}
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	groupName := mux.Vars(r)["groupName"]

	var result int
	err := db.QueryRow("select * from users.delete_group('" + groupName + "')").Scan(&result)
	if err != nil {
		log.Print(err)
	} else {
		if result == 0 {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("A group with groupName=%s does not exist", groupName))
		}
	}

}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	if user.UserID == "" {
		respondWithError(w, http.StatusBadRequest, "UserId is required.")
		return
	}

	if user.FirstName == "" {
		respondWithError(w, http.StatusBadRequest, "FirstName is required.")
		return
	}

	if user.LastName == "" {
		respondWithError(w, http.StatusBadRequest, "LastName is required.")
		return
	}

	var result User
	err := db.QueryRow("select * from users.insert_user('"+user.UserID+"', '"+user.FirstName+"', '"+user.LastName+"')").Scan(&result.UserID, &result.FirstName, &result.LastName)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "duplicate key") {
			respondWithError(w, http.StatusConflict, fmt.Sprintf("A user with userID=%s already exists", user.UserID))
		}
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	var result User
	err := db.QueryRow("select * from users.get_user('"+userID+"')").Scan(&result.UserID, &result.FirstName, &result.LastName)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
		}
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	userID := mux.Vars(r)["userId"]

	var result User
	err := db.QueryRow("select * from users.update_user('"+userID+"', '"+user.FirstName+"', '"+user.LastName+"')").Scan(&result.UserID, &result.FirstName, &result.LastName)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
		}
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	var result int
	err := db.QueryRow("select * from users.delete_user('" + userID + "')").Scan(&result)
	if err != nil {
		log.Print(err)
	} else {
		if result == 0 {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
		}
	}

}

func respondWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Error{Message: message})
}
