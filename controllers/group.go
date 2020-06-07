package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"user.sor/models"
	"user.sor/utils"
)

type GroupController struct {
}

// CreateGroup creates a group in the database
func (c GroupController) CreateGroup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var group models.Group

		json.NewDecoder(r.Body).Decode(&group)

		if group.Name == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Name is required.")
			return
		}

		var result models.Group
		err := db.QueryRow("select * from users.insert_group('" + group.Name + "')").Scan(&result.Name)
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "duplicate key") {
				utils.RespondWithError(w, http.StatusConflict, fmt.Sprintf("A group with Name=%s already exists", group.Name))
			}
		} else {
			utils.ResponseJSON(w, result)
		}
	}
}

// GetGroup returns a group from the db using the group name provided on the url
func (c GroupController) GetGroup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := mux.Vars(r)["groupName"]

		rows, err := db.Query("select * from users.get_user_group('" + groupName + "')")
		if err != nil {
			log.Print(err)
		}

		var userGroup models.UserGroup
		defer rows.Close()
		for rows.Next() {
			var userID string
			err = rows.Scan(&userID)
			if err != nil {
				log.Print(err)
				utils.RespondWithError(w, http.StatusInternalServerError, "An error occurred, please try again later.")
			}
			userGroup.UserIDs = append(userGroup.UserIDs, userID)
		}

		err = rows.Err()
		if err != nil {
			log.Print(err)
		}

		if len(userGroup.UserIDs) > 0 {
			utils.ResponseJSON(w, userGroup)
		} else {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Either there is no group named %s or there are no users associated with it", groupName))
		}
	}
}

// UpdateGroup updates the group in the db
func (c GroupController) UpdateGroup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userGroup models.UserGroup

		json.NewDecoder(r.Body).Decode(&userGroup)

		groupName := mux.Vars(r)["groupName"]

		_, err := db.Exec("DELETE FROM users.user_group")
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "An error occurred, please try again later.")
		}

		_, err = db.Exec("INSERT INTO users.user_group (user_id, group_name) VALUES (unnest($1::text[]), $2)", pq.Array(userGroup.UserIDs), groupName)
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "foreign key constraint") {
				utils.RespondWithError(w, http.StatusNotFound, "Invalid user or group name.")
			}
		}
	}
}

// DeleteGroup deletes a group from the db using the group name provided on the url
func (c GroupController) DeleteGroup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := mux.Vars(r)["groupName"]

		var result int
		err := db.QueryRow("select * from users.delete_group('" + groupName + "')").Scan(&result)
		if err != nil {
			log.Print(err)
		} else {
			if result == 0 {
				utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A group with groupName=%s does not exist", groupName))
			}
		}
	}
}
