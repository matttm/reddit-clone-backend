package persons



import (
//	"github.com/stretchr/testify/assert"
	errors "reddit-clone-backend/pkg/errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	database "reddit-clone-backend/internal/pkg/db/mysql"
	"reddit-clone-backend/internal/utilities"
)

func TestPerson_Create(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

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
//	_, err := post.Save()
//	assert.Error(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}