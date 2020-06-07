package grouprepository

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"user.sor/models"
)

type GroupRepository struct{}

// InsertGroup inserts the group to the db
func (u GroupRepository) InsertGroup(db *sql.DB, group models.Group) (models.Group, error) {
	var result models.Group
	err := db.QueryRow("select * from users.insert_group('" + group.Name + "')").Scan(&result.Name)

	if err != nil {
		log.Print(err)
		return group, err
	}

	return result, err
}

// GetGroup returns the userIDs associated with the group
func (u GroupRepository) GetGroup(db *sql.DB, groupName string) (models.UserGroup, error) {
	var userGroup models.UserGroup

	rows, err := db.Query("select * from users.get_user_group('" + groupName + "')")

	if err != nil {
		log.Print(err)
		return userGroup, err
	}

	defer rows.Close()
	for rows.Next() {
		var userID string
		err = rows.Scan(&userID)
		if err != nil {
			log.Print(err)
		}
		userGroup.UserIDs = append(userGroup.UserIDs, userID)
	}

	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	return userGroup, err
}

// UpdateGroup replaces the users associated with the group
func (u GroupRepository) UpdateGroup(db *sql.DB, groupName string, userGroup models.UserGroup) (models.UserGroup, error) {
	_, err := db.Exec("DELETE FROM users.user_group")
	if err != nil {
		log.Print(err)
		return userGroup, err
	}

	_, err = db.Exec("INSERT INTO users.user_group (user_id, group_name) VALUES (unnest($1::text[]), $2)", pq.Array(userGroup.UserIDs), groupName)
	if err != nil {
		log.Print(err)
		return userGroup, err
	}

	return userGroup, err
}

// DeleteGroup deletes the group from the db
func (u GroupRepository) DeleteGroup(db *sql.DB, groupName string) (int, error) {
	var result int
	err := db.QueryRow("select * from users.delete_group('" + groupName + "')").Scan(&result)

	if err != nil {
		log.Print(err)
		return result, err
	}

	return result, err
}
