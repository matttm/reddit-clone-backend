package utilities

import (
	"log"

	"database/sql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	database "reddit-clone-backend/internal/pkg/db/mysql"
)

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
