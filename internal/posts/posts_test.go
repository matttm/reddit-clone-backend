package posts


import (
	"reddit-clone-backend/internal/persons"
	errors "reddit-clone-backend/pkg/utilities"
	"testing"
	"log"

	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"

	database "reddit-clone-backend/internal/pkg/db/mysql"
//	"github.com/stretchr/testify/assert"
)

func TestPost_Save(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "INSERT INTO POSTS\\(PERSON_ID, TITLE, BODY, VIEWS\\) VALUES\\(\\?,\\?,\\?,\\?\\)"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(person.Id, post.Title, post.Body, post.Views).WillReturnResult(sqlmock.NewResult(1, 1));
	post.Save()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Save_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()

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
	post.Save()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Update(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "UPDATE POSTS SET TITLE = \\?, BODY = \\? WHERE ID = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(post.Title, post.Body, post.Id).WillReturnResult(sqlmock.NewResult(1, 1));
	post.Update()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Update_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()

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
	post.Update()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPost_Delete(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()

	var post Post
	var person persons.Person
	person.Id = "1"
	post.Person = &person
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0

	query := "DELETE FROM POSTS WHERE ID = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(post.Id).WillReturnResult(sqlmock.NewResult(1, 1));
	post.Delete()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


func TestPost_Delete_Error(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()

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
	post.Delete()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAll(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()
	postMockRows := sqlmock.NewRows([]string{"Id", "Title", "Body", "Viewx"}).
		AddRow("1", "Genesis Post", "my body", "0").
		AddRow("2", "Another Post", "my body", "1")

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
	database.Db, mock = NewMock()
	defer Close()
	id := 1
	postMockRow := sqlmock.NewRows([]string{"Id", "Title", "Body", "Viewx"}).
		AddRow("1", "Genesis Post", "my body", "0")

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


// TODO: find bwtter placement for this db mock
func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
// Close attaches the provider and close the connection
func Close() {
	database.Db.Close()
}