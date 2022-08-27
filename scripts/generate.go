package main

import (
	"log"
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

	for i := 1; i < 1000; i++ {
		var post posts.Post
		post.Title = gofakeit.Question()
		post.Body = gofakeit.Paragraph(1, 5, 5, " ")
		log.Print(post.Title, post.Body)
		post.Save()
	}
}
