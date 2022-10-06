package posts


import (
    "testing"
	"log"

	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"

	database "reddit-clone-backend/internal/pkg/db/mysql"
//	"github.com/stretchr/testify/assert"
)

func TestSavePost(t *testing.T) {
	var mock sqlmock.Sqlmock
	database.Db, mock = NewMock()
	defer Close()
	mock.ExpectPrepare("INSERT INTO POSTS(PERSON_ID, TITLE, BODY, VIEWS) VALUES(\\?,\\?,\\?,\\?)")
	mock.ExpectExec("INSERT INTO POSTS(PERSON_ID, TITLE, BODY, VIEWS) VALUES(\\?,\\?,\\?,\\?)").WithArgs(1, 1, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1));


	var post Post
	post.Title = "Test"
	post.Body = "of the century"
	post.Views = 0
	post.Save()

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