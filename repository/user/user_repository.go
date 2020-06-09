package userrepository

import (
	"database/sql"
	"log"

	"user.sor/models"
)

// UserRepository is the DAL for the users schema user table
type UserRepository struct{}

// InsertUser inserts the user to the db
func (u UserRepository) InsertUser(db *sql.DB, user models.User) (models.User, error) {
	var result models.User
	row := db.QueryRow("select * from users.insert_user($1, $2, $3)", user.UserID, user.FirstName, user.LastName)
	err := row.Scan(&result.UserID, &result.FirstName, &result.LastName)

	if err != nil {
		log.Print(err)
		return user, err
	}

	return result, err
}

// GetUser returns the user from the db given the userID
func (u UserRepository) GetUser(db *sql.DB, userID string) (models.User, error) {
	var result models.User
	err := db.QueryRow("select * from users.get_user($1)", userID).Scan(&result.UserID, &result.FirstName, &result.LastName)

	if err != nil {
		log.Print(err)
		return result, err
	}

	return result, err
}

// UpdateUser updates the user in the db
func (u UserRepository) UpdateUser(db *sql.DB, userID string, user models.User) (models.User, error) {
	var result models.User
	err := db.QueryRow("select * from users.update_user($1, $2, $3)", userID, user.FirstName, user.LastName).Scan(&result.UserID, &result.FirstName, &result.LastName)

	if err != nil {
		log.Print(err)
		return result, err
	}

	return result, err
}

// DeleteUser deletes the user from the db
func (u UserRepository) DeleteUser(db *sql.DB, userID string) (int, error) {
	var result int = 0
	err := db.QueryRow("select * from users.delete_user($1)", userID).Scan(&result)

	if err != nil {
		log.Print(err)
		return result, err
	}

	return result, err
}
