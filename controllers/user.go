package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"user.sor/models"
	"user.sor/utils"
)

type UserController struct {
}

// CreateUser creates a user in the database
func (c UserController) CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		json.NewDecoder(r.Body).Decode(&user)

		if user.UserID == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "UserId is required.")
			return
		}

		if user.FirstName == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "FirstName is required.")
			return
		}

		if user.LastName == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "LastName is required.")
			return
		}

		var result models.User
		err := db.QueryRow("select * from users.insert_user('"+user.UserID+"', '"+user.FirstName+"', '"+user.LastName+"')").Scan(&result.UserID, &result.FirstName, &result.LastName)
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "duplicate key") {
				utils.RespondWithError(w, http.StatusConflict, fmt.Sprintf("A user with userID=%s already exists", user.UserID))
			}
		} else {
			utils.ResponseJSON(w, result)
		}
	}
}

// GetUser returns a user from the db using the userID provided on the url
func (c UserController) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]

		var result models.User
		err := db.QueryRow("select * from users.get_user('"+userID+"')").Scan(&result.UserID, &result.FirstName, &result.LastName)
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "no rows in result set") {
				utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
			}
		} else {
			utils.ResponseJSON(w, result)
		}
	}
}

// UpdateUser updates the user in the db
func (c UserController) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		json.NewDecoder(r.Body).Decode(&user)

		userID := mux.Vars(r)["userID"]

		var result models.User
		err := db.QueryRow("select * from users.update_user('"+userID+"', '"+user.FirstName+"', '"+user.LastName+"')").Scan(&result.UserID, &result.FirstName, &result.LastName)
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "no rows in result set") {
				utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
			}
		} else {
			utils.ResponseJSON(w, result)
		}
	}
}

// DeleteUser deletes a user from the db using the userID provided on the url
func (c UserController) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]

		var result int
		err := db.QueryRow("select * from users.delete_user('" + userID + "')").Scan(&result)
		if err != nil {
			log.Print(err)
		} else {
			if result == 0 {
				utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
			}
		}
	}
}
