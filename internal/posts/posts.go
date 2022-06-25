package posts

import (
	"log"
	"reddit-clone-backend/internal/persons"

	database "reddit-clone-backend/internal/pkg/db/mysql"
)

type Post struct {
	Id        string
	Title     string
	Body      string
	Views     int
	CreatedAt string
	UpdatedAt string
	Person    *persons.Person
}

//#2
func (post Post) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO POSTS(TITLE, BODY, VIEWS) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(post.Title, post.Body, post.Views)
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

func GetAll() []Post {
	stmt, err := database.Db.Prepare("SELECT ID, TITLE, BODY FROM POSTS")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Body)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return posts
}
