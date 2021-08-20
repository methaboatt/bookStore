package server

import (
	"database/sql"
	"fmt"
	"strconv"

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

// SetupDB : initializing mysql database
func SetupDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func exampleFunc(c *gin.Context) {
	var result Result
	result.ID = c.Query("id")
	result.Name = c.PostForm("name")
	result.Message = c.PostForm("message")

	c.JSON(200, gin.H{
		"message": result,
	})
}

func exampleJSON(c *gin.Context) {
	var input Result
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	c.JSON(200, gin.H{
		"id":      input.ID,
		"name":    input.Name,
		"message": input.Message,
	})
}

type Result struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Customer struct {
	ID        string `json:"id"`
	Age       string `json:"age"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func HandleInsertUser(c *gin.Context) {
	customer := Customer{}
	customer.Age = c.PostForm("age")
	customer.Email = c.PostForm("email")
	customer.Address1 = c.PostForm("address1")
	customer.Address2 = c.PostForm("address2")
	customer.FirstName = c.PostForm("firstName")
	customer.LastName = c.PostForm("lastName")
	customer.Username = c.PostForm("username")
	customer.Password = c.PostForm("password")

	c.JSON(200, gin.H{
		"message": InsertUser(customer, c),
	})

}

func InsertUser(customer Customer, c *gin.Context) string {

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `
INSERT INTO users (age, email, first_name, last_name, address1, address2, username, password)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, customer.Age, customer.Email, customer.FirstName, customer.LastName, customer.Address1, customer.Address2, customer.Username, customer.Password).Scan(&id)

	if err != nil {
		//panic(err)
		return (err.Error())
	}

	return "Insert successful new ID is :" + strconv.Itoa(id)
}

func HandleGetUser(c *gin.Context) {
	var customer Customer
	var customerlst []Customer
	db := c.MustGet("db").(*sql.DB)
	sqlStatement := `SELECT * from users`

	row, err := db.Query(sqlStatement)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	for row.Next() {
		err := row.Scan(&customer.ID, &customer.Age, &customer.FirstName, &customer.LastName, &customer.Address1, &customer.Address2, &customer.Username, &customer.Password, &customer.Email)
		if err != nil {
			panic(err.Error())
		}
		customerlst = append(customerlst, customer)
	}

	c.JSON(200, gin.H{
		"message": customerlst,
	})
}

func HandleUpdateUser(c *gin.Context) {
	customer := Customer{}
	customer.ID = c.PostForm("id")
	customer.FirstName = c.PostForm("firstName")

	c.JSON(200, gin.H{
		"message":  customer,
		"message1": updateUser(customer, c),
	})

}

func updateUser(customer Customer, c *gin.Context) string {

	db := c.MustGet("db").(*sql.DB)

	fmt.Println(customer)
	sqlStatement := `
	UPDATE users SET  first_name= $1 
	WHERE id =$2
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, customer.FirstName, customer.ID).Scan(&id)

	if err != nil {
		//panic(err)
		return (err.Error())
	}

	return "update successful row ID is :" + strconv.Itoa(id)
}
func HandlePatchUser(c *gin.Context) {
	customer := Customer{}
	customer.ID = c.PostForm("id")
	customer.FirstName = c.PostForm("firstName")

	c.JSON(200, gin.H{
		"message":  customer,
		"message1": updateUser(customer, c),
	})

}
