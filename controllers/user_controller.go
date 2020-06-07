package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	userrepository "user.sor/repository/user"

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

		userrepo := userrepository.UserRepository{}
		user, err := userrepo.InsertUser(db, user)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				utils.RespondWithError(w, http.StatusConflict, fmt.Sprintf("A user with userID=%s already exists", user.UserID))
			}
		} else {
			utils.ResponseJSON(w, user)
		}
	}
}

// GetUser returns a user from the db using the userID provided on the url
func (c UserController) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]

		userrepo := userrepository.UserRepository{}
		user, err := userrepo.GetUser(db, userID)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
			}
		} else {
			utils.ResponseJSON(w, user)
		}
	}
}

// UpdateUser updates the user in the db
func (c UserController) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		json.NewDecoder(r.Body).Decode(&user)

		userID := mux.Vars(r)["userID"]

		userrepo := userrepository.UserRepository{}
		user, err := userrepo.UpdateUser(db, userID, user)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
			}
		} else {
			utils.ResponseJSON(w, user)
		}
	}
}

// DeleteUser deletes a user from the db using the userID provided on the url
func (c UserController) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]

		userrepo := userrepository.UserRepository{}
		result, _ := userrepo.DeleteUser(db, userID)
		if result == 0 {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A user with userID=%s does not exist", userID))
		}
	}
}
