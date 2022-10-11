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

// #2
func (post Post) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO POSTS(PERSON_ID, TITLE, BODY, VIEWS) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(post.Person.Id, post.Title, post.Body, post.Views)
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

func (post Post) Update() (int64, error) {
	stmt, err := database.Db.Prepare("UPDATE POSTS SET TITLE = ?, BODY = ? WHERE ID = ?")
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	ret, err := stmt.Exec(post.Title, post.Body, post.Id)
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	log.Print("Row Updated!", ret)
	return 0, nil
}

func (post Post) Delete() (int64, error) {
	stmt, err := database.Db.Prepare("DELETE FROM POSTS WHERE ID = ?")
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	ret, err := stmt.Exec(post.Id)
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	log.Print("Row Deleted!", ret)
	return 0, nil
}

func GetAll() []Post {
	stmt, err := database.Db.Prepare(`
	SELECT post.ID, post.TITLE, post.BODY, post.VIEWS, person.ID,
		person.USERNAME, person.CREATED_AT, post.CREATED_AT, post.UPDATED_AT FROM POSTS post
		JOIN PERSONS person ON post.PERSON_ID = person.ID
	`)
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
		var person persons.Person
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.Views,
			&person.Id,
			&person.Username,
			&person.CreatedAt,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		post.Person = &person
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return posts
}

func Get(id int) Post {
	var person persons.Person
	var post Post
	stmt, err := database.Db.Prepare(`
	SELECT post.ID, post.TITLE, post.BODY, post.VIEWS, person.ID,
		person.USERNAME, person.CREATED_AT, post.CREATED_AT, post.UPDATED_AT FROM POSTS post
		JOIN PERSONS person ON post.PERSON_ID = person.ID WHERE post.ID = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.Views,
		&person.Id,
		&person.Username,
		&person.CreatedAt,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	post.Person = &person
	if err != nil {
		log.Fatal(err)
	}
	return post
}
