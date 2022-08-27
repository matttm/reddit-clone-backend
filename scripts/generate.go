package main

import (
	"log"
	"strconv"

	"reddit-clone-backend/internal/persons"
	database "reddit-clone-backend/internal/pkg/db/mysql"
	"reddit-clone-backend/internal/posts"

	_ "github.com/go-sql-driver/mysql"

	"github.com/brianvoe/gofakeit/v6"
)

const defaultPort = "8080"

func main() {
	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	for i := 1; i <= 100; i++ {
		var person persons.Person
		person.Username = gofakeit.Username()
		person.Password = gofakeit.Password(true, true, true, true, false, 10)
		log.Print("User generated ", person.Username, " ", person.Password)
		person.Create()
	}
	for i := 1; i < 100; i++ {
		var post posts.Post
		var person persons.Person
		post.Title = gofakeit.Question()
		post.Body = gofakeit.Paragraph(1, 5, 5, " ")
		person.Id = strconv.Itoa(gofakeit.Number(1, 100))
		post.Person = &person
		log.Print("Post generzated ", post.Title, " ", post.Body, " ", post.Person.Id)
		post.Save()
	}
}
