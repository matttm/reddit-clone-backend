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
	stmt, err := database.Db.Prepare("INSERT INTO POSTS(PERSON_ID, TITLE, BODY, VIEWS) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(1, post.Title, post.Body, post.Views)
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

func (post Post) Update() int64 {
	stmt, err := database.Db.Prepare("UPDATE POSTS SET TITLE = ?, BODY = ? WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	ret, err := stmt.Exec(post.Title, post.Body, post.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Row Updated!", ret)
	return 0
}

func (post Post) Delete() int64 {
	stmt, err := database.Db.Prepare("DELETE POSTS WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	ret, err := stmt.Exec(post.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Row Deleted!", ret)
	return 0
}

func GetAll() []Post {
	stmt, err := database.Db.Prepare("SELECT ID, TITLE, BODY, CREATED_AT, UPDATED_AT FROM POSTS")
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
		err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
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

func Get(id int) Post {
	var post Post
	stmt, err := database.Db.Prepare("SELECT ID, TITLE, BODY FROM POSTS WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&post.Id, &post.Title, &post.Body)
	if err != nil {
		log.Fatal(err)
	}
	return post
}
