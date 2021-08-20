package product

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"Name"`
	BookType  string `json:"bookType"`
	Price     string `json:"price"`
	OwnerName string `json:"ownername"`
}

func HandleGetAllbook(c *gin.Context) {

	var book Book
	var bookList []Book

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `SELECT * from book `

	row, err := db.Query(sqlStatement)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	for row.Next() {
		err := row.Scan(&book.ID, &book.Name, &book.BookType, &book.Price, &book.OwnerName)
		if err != nil {
			panic(err.Error())
		}
		bookList = append(bookList, book)
	}

	c.JSON(200, gin.H{
		"message": bookList,
	})
}

func HandleGetbook(c *gin.Context) {

	var book Book
	var bookList []Book

	db := c.MustGet("db").(*sql.DB)
	book.ID = c.Param("id")
	sqlStatement := `SELECT * from book WHERE id = ` + book.ID

	row, err := db.Query(sqlStatement)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	for row.Next() {
		err := row.Scan(&book.ID, &book.Name, &book.BookType, &book.Price, &book.OwnerName)
		if err != nil {
			panic(err.Error())
		}
		bookList = append(bookList, book)
	}

	c.JSON(200, gin.H{
		"message": bookList,
	})

}

func HandleInsertBook(c *gin.Context) {
	book := Book{}

	book.Name = c.PostForm("name")
	book.OwnerName = c.PostForm("ownername")
	book.Price = c.PostForm("price")
	book.BookType = c.PostForm("booktype")

	c.JSON(200, gin.H{
		"message": InsertUser(book, c),
	})

}

func InsertUser(book Book, c *gin.Context) string {

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `
INSERT INTO book (name, book_type, price, owner_name)
VALUES ($1, $2, $3, $4)
RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, book.Name, book.BookType, book.Price, book.OwnerName).Scan(&id)

	if err != nil {
		//panic(err)
		return (err.Error())
	}

	return "Insert successful new ID is :" + strconv.Itoa(id)
}

func HandleUpdateBook(c *gin.Context) {
	book := Book{}
	book.ID = c.PostForm("id")
	book.Name = c.PostForm("name")
	fmt.Println(book.ID, book.Name, c.PostForm("name"))

	c.JSON(200, gin.H{
		"message1": updateBook(book, c),
	})

}

func updateBook(book Book, c *gin.Context) string {

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `
	UPDATE book SET name= $1 
	WHERE id =$2
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, book.Name, book.ID).Scan(&id)

	if err != nil {
		//panic(err)
		return (err.Error())
	}

	return "update successful row ID is :" + strconv.Itoa(id)
}
