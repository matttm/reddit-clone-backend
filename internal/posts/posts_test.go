package posts

import (
	"reddit-clone-backend/internal/persons"
	"reddit-clone-backend/internal/utilities"

	"github.com/stretchr/testify/assert"

	errors "reddit-clone-backend/pkg/errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	database "reddit-clone-backend/internal/pkg/db/mysql"
)

func TestPost_Save(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "INSERT INTO POSTS\\(PERSON_ID, TITLE, BODY, VIEWS\\) VALUES\\(\\?,\\?,\\?,\\?\\)"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(person.Id, post.Title, post.Body, post.Views).WillReturnResult(sqlmock.NewResult(1, 1))
	post.Save()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Save_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "INSERT INTO POSTS\\(PERSON_ID, TITLE, BODY, VIEWS\\) VALUES\\(\\?,\\?,\\?,\\?\\)"
	mock.ExpectPrepare(query).WillReturnError(&errors.GenericError{"Error during prepare"})
	//	mock.ExpectExec(query).WithArgs(person.Id, post.Title, post.Body, post.Views).WillReturnResult(sqlmock.NewResult(1, 1));
	_, err := post.Save()
	assert.Error(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Update(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "UPDATE POSTS SET TITLE = \\?, BODY = \\? WHERE ID = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(post.Title, post.Body, post.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	post.Update()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Update_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "UPDATE POSTS SET TITLE = \\?, BODY = \\? WHERE ID = \\?"
	mock.ExpectPrepare(query).WillReturnError(&errors.GenericError{"Error during prepare"})
	//	mock.ExpectExec(query).WithArgs(post.Title, post.Body, post.Id).WillReturnResult(sqlmock.NewResult(1, 1));
	_, err := post.Update()
	assert.Error(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Delete(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "DELETE FROM POSTS WHERE ID = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(post.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	post.Delete()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Delete_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "DELETE FROM POSTS WHERE ID = \\?"
	mock.ExpectPrepare(query).WillReturnError(&errors.GenericError{"Error during prepare"})
	//	mock.ExpectExec(query).WithArgs(post.Id).WillReturnResult(sqlmock.NewResult(1, 1));
	_, err := post.Delete()
	assert.Error(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAll(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()
	postMockRows := sqlmock.NewRows([]string{
		"post.ID", "post.TITLE", "post.BODY", "post.VIEWS", "person.ID",
		"person.USERNAME", "person.CREATED_AT", "post.CREATED_AT", "post.UPDATED_AT",
	}).
		AddRow("1", "Genesis Post", "my body", "0", "1", "joe", "0", "0", "0").
		AddRow("2", "Srcond Post", "my body", "0", "1", "joe", "0", "0", "0")

	query := `
	SELECT post.ID, post.TITLE, post.BODY, post.VIEWS, person.ID,
		person.USERNAME, person.CREATED_AT, post.CREATED_AT, post.UPDATED_AT FROM POSTS post
		JOIN PERSONS person ON post.PERSON_ID = person.ID
	`
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WillReturnRows(postMockRows)
	GetAll()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGet(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = utilities.NewMock()
	defer utilities.Close()
	id := 1
	postMockRow := sqlmock.NewRows([]string{
		"post.ID", "post.TITLE", "post.BODY", "post.VIEWS", "person.ID",
		"person.USERNAME", "person.CREATED_AT", "post.CREATED_AT", "post.UPDATED_AT",
	}).
		AddRow("1", "Genesis Post", "my body", "0", "1", "joe", "0", "0", "0")

	query := `
	SELECT post.ID, post.TITLE, post.BODY, post.VIEWS, person.ID,
		person.USERNAME, person.CREATED_AT, post.CREATED_AT, post.UPDATED_AT FROM POSTS post
		JOIN PERSONS person ON post.PERSON_ID = person.ID WHERE post.ID = \\?
	`
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(postMockRow)
	Get(id)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
