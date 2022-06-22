package links

import (
	database "internal/pkg/db/mysql"
	"github.com/natttn/reddit-clone-backend/internal/users"
	"log"

	"github.com/go-sql-driver/mysql"
)


type Post struct {
    id float32
    title string
    body string
    views int
    createdAt string
    updatedAt string
    person *persons.person
}

//#2
func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
