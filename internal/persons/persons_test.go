package persons



import (
//	"reddit-clone-backend/pkg/crypto"

		"github.com/stretchr/testify/assert"
	errors "reddit-clone-backend/pkg/errors"
	"reddit-clone-backend/pkg/mocks"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	database "reddit-clone-backend/internal/pkg/db/mysql"
	"reddit-clone-backend/internal/utilities"
)

func TestPerson_Create(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()
	CryptoHashPassword = mocks.HashPasswordMock

	var person Person
	person.Username = "matttm"
	person.Password = "bird314"

	query := "INSERT INTO PERSONS\\(USERNAME, PASSWORD\\) VALUES\\(\\?,\\?\\)"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(person.Username, person.Password).WillReturnResult(sqlmock.NewResult(1, 1));
	person.Create()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPerson_Create_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var person Person
	person.Username = "matttm"
	person.Password = "bird314"

	query := "INSERT INTO PERSONS\\(USERNAME, PASSWORD\\) VALUES\\(\\?,\\?\\)"
	mock.ExpectPrepare(query).WillReturnError(&errors.GenericError{"Error during prepare"})
//	mock.ExpectExec(query).WithArgs(person.Id, post.Title, post.Body, post.Views).WillReturnResult(sqlmock.NewResult(1, 1));
	_, err := person.Create()
	assert.Error(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthenticate(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()
	CryptoCheckPassword = mocks.CheckPasswordHashMock

	username := "matttm"
	password := "bird314"

	query := "SELECT PASSWORD FROM PERSONS WHERE USERNAME = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{
		"Password"}).AddRow("password"),
	)
	res, _ := Authenticate(username, password)
	assert.NotNil(t, res)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserIdByUsername(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()
	CryptoHashPassword = mocks.HashPasswordMock

	username := "matttm"

	query := "SELECT ID FROM PERSONS WHERE USERNAME = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{
		"Id"}).
		AddRow("1"),
	)
	res, _ := GetUserIdByUsername(username)
	assert.NotNil(t, res)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserIdByUsername_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()
	CryptoHashPassword = mocks.HashPasswordMock

	username := "matttm"

	query := "SELECT ID FROM PERSONS WHERE USERNAME = \\?"
	mock.ExpectPrepare(query).WillReturnError(&errors.GenericError{"Error during prepare"})
//	mock.ExpectQuery(query).WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{
//		"Id", "Username", "Password", "CreatedAt", "UpdatedAt",
//	}).
//		AddRow("1", "matttm", "password", "0", "0")
	_, err := GetUserIdByUsername(username)
	assert.Error(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
