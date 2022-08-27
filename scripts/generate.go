package main

import (
	"log"
	"strconv"

	database "reddit-clone-backend/internal/pkg/db/mysql"
	"reddit-clone-backend/internal/posts"
	"reddit-clone-backend/internal/persons"

	_ "github.com/go-sql-driver/mysql"

	"github.com/brianvoe/gofakeit/v6"
)

const defaultPort = "8080"

func main() {
	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	for i := 1; i < 100; i++ {
		var person persons.Person
		person.Username = gofakeit.Username()
		person.Password = gofakeit.Password(true, true, true, true, false, 10)
		log.Print("User generated ", person.Username, " ", person.Password)
		// person.Create()
	}
	for i := 1; i < 1000; i++ {
		var post posts.Post
		post.Title = gofakeit.Question()
		post.Body = gofakeit.Paragraph(1, 5, 5, " ")
		post.Person.Id = strconv.Itoa(gofakeit.Number(1, 100))
		// log.Print(p"Post generzated ", ost.Title, " ", post.Body)
		// post.Save()
	}
}
