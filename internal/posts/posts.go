package posts

import (
	"log"
	"reddit-clone-backend/internal/pkg/persons"

	database "reddit-clone-backend/internal/pkg/db/mysql"
)

type Post struct {
	id        float32
	title     string
	body      string
	views     int
	createdAt string
	updatedAt string
	person    *persons.Person
}

//#2
func (post Post) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Posts(TITLE, BODY, VIEWS) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(post.title, post.body, post.views)
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
