package userrepository

import (
	"testing"

	"user.sor/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was no expected when opening a stub db connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"userID", "FirstName", "LastName"}).
		AddRow("testId", "testFirstName", "testLastName")

	user := models.User{UserID: "testId", FirstName: "testFirstName", LastName: "testLastName"}
	userrepo := UserRepository{}

	mock.ExpectQuery("users.insert_user").
		WithArgs(user.UserID, user.FirstName, user.LastName).
		WillReturnRows(rows)

	if user, err = userrepo.InsertUser(db, user); err != nil {
		t.Errorf("error was not expected while inserting user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was no expected when opening a stub db connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"userID", "FirstName", "LastName"}).
		AddRow("testId", "testFirstName", "testLastName")

	userID := "testId"
	userrepo := UserRepository{}

	mock.ExpectQuery("users.get_user").
		WithArgs(userID).
		WillReturnRows(rows)

	if _, err = userrepo.GetUser(db, userID); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was no expected when opening a stub db connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"userID", "FirstName", "LastName"}).
		AddRow("testId", "testFirstName", "testLastName")

	userID := "testId"
	user := models.User{UserID: userID, FirstName: "testFirstName", LastName: "testLastName"}
	userrepo := UserRepository{}

	mock.ExpectQuery("users.update_user").
		WithArgs(userID, user.FirstName, user.LastName).
		WillReturnRows(rows)

	if user, err = userrepo.UpdateUser(db, userID, user); err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was no expected when opening a stub db connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"out_rowsaffected"}).
		AddRow(1)

	userID := "testId"
	userrepo := UserRepository{}

	mock.ExpectQuery("users.delete_user").
		WithArgs(userID).
		WillReturnRows(rows)

	if _, err = userrepo.DeleteUser(db, userID); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
