package main

import (
	"fmt"

	"github.com/bookStore/internal/database"
	"github.com/bookStore/internal/product"
	"github.com/bookStore/internal/user"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := database.SetupDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	defer db.Close()

	fmt.Println("t")

	r.POST("/testAddUser", user.HandleInsertUser)
	r.GET("/testGetUser/:id", user.HandleGetUser)
	r.PUT("/testUpdateUser", user.HandleUpdateUser)

	r.GET("/getAllBook", product.HandleGetAllbook)
	r.GET("/getBook/:id", product.HandleGetbook)
	r.POST("/AddBook", product.HandleInsertBook)
	r.POST("/UpdateBook", product.HandleUpdateBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
