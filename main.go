package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "boat"
	dbname   = "bookStore"
)

type Customer struct {
	age       int
	email     string
	firstName string
	lastName  string
	address1  string
	address2  string
	username  string
	password  string
}

func main() {
	// insertUser(&Customer{age: 10,
	// 	email:     "jon111@calhoun.io",
	// 	firstName: "Meeeee",
	// 	lastName:  "R",
	// 	address1:  "Bangkok",
	// 	address2:  "Thailand",
	// 	username:  "MMM",
	// 	password:  "Password"})
	fmt.Println("")
}

func insertUser(c *gin.Context) {
	customer := Customer{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name, address1, address2, username, password)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, customer.age, customer.email, customer.firstName, customer.lastName, customer.address1, customer.address2, customer.username, customer.password).Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
