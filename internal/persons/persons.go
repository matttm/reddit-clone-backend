package persons

import (
	"log"
	database "reddit-clone-backend/internal/pkg/db/mysql"
	"reddit-clone-backend/pkg/crypto"
)

type Person struct {
	Id        string
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
	// posts     []*posts.Post
}

//#2
func (person Person) Create() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO PERSONS(USERNAME, PASSWORD) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	hashPassword, err := crypto.HashPassword(person.Password)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(person.Username, hashPassword)
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

func GetAll() []Person {
	stmt, err := database.Db.Prepare("SELECT ID, USERNAME FROM PERSONS")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var persons []Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Id, &person.Username)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return persons
}
