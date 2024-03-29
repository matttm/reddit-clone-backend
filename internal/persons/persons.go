package persons

import (
	"database/sql"
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

// variables for monkey patching during testing
var (
	CryptoHashPassword = crypto.HashPassword
	CryptoCheckPassword = crypto.CheckPasswordHash
)

/**
  * @function Create
  * @description Adds a person to connected db
  *
  * @return an integer representing the person id and an error
*/
func (person Person) Create() (int64, error) {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO PERSONS(USERNAME, PASSWORD) VALUES(?,?)")
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	hashPassword, err := CryptoHashPassword(person.Password)
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	res, err := stmt.Exec(person.Username, hashPassword)
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Print("Error:", err.Error())
		return 0, err
	}
	log.Print("Row inserted!")
	return id, nil
}

/**
  * @function Authenticate
  * @description determinning whether password is correct for iven user
  *
  * @return an boolean representing the authentication is successful and an error
*/
func Authenticate(username string, password string) (bool, error) {
	statement, err := database.Db.Prepare("SELECT PASSWORD FROM PERSONS WHERE USERNAME = ?")
	if err != nil {
		log.Printf(err.Error())
		return false, err
	}
	row := statement.QueryRow(username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			log.Printf(err.Error())
			return false, err
		}
	}

	return CryptoCheckPassword(password, hashedPassword), nil
}

/**
  * @function GetAll
  * @description Retrieve all persons from db
  *
  * @return an array containing all persons and an error
*/
func GetAll() ([]Person, error) {
	stmt, err := database.Db.Prepare("SELECT ID, USERNAME, CREATED_AT, UPDATED_AT FROM PERSONS")
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer rows.Close()
	var persons []Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Id, &person.Username, &person.CreatedAt, &person.UpdatedAt)
		if err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return persons, nil
}
/**
  * @function GetUserIdByUsername
  * @description get a user's id from their username
  *
  * @return an integer representing the person id and an error
*/
func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("SELECT ID FROM PERSONS WHERE USERNAME = ?")
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}
