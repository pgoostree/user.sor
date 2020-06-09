package grouprepository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"user.sor/models"
)

func TestShouldInsertGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was no expected when opening a stub db connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"group_name"}).
		AddRow("testGroupName")

	group := models.Group{Name: "testGroupName"}
	grouprepo := GroupRepository{}

	mock.ExpectQuery("users.insert_group").
		WithArgs(group.Name).
		WillReturnRows(rows)

	if group, err = grouprepo.InsertGroup(db, group); err != nil {
		t.Errorf("error was not expected while inserting group: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
