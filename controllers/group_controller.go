package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	grouprepository "user.sor/repository/group"

	"github.com/gorilla/mux"
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

		groupRepo := grouprepository.GroupRepository{}
		group, err := groupRepo.InsertGroup(db, group)

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				utils.RespondWithError(w, http.StatusConflict, fmt.Sprintf("A group with Name=%s already exists", group.Name))
			}
		} else {
			utils.ResponseJSON(w, group)
		}
	}
}

// GetGroup returns a group from the db using the group name provided on the url
func (c GroupController) GetGroup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := mux.Vars(r)["groupName"]

		groupRepo := grouprepository.GroupRepository{}
		userGroup, err := groupRepo.GetGroup(db, groupName)

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

		groupRepo := grouprepository.GroupRepository{}
		userGroup, err := groupRepo.UpdateGroup(db, groupName, userGroup)
		if err != nil {
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

		groupRepo := grouprepository.GroupRepository{}
		result, _ := groupRepo.DeleteGroup(db, groupName)

		if result == 0 {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("A group with groupName=%s does not exist", groupName))
		}
	}
}
