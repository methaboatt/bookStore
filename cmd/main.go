package main

import (
	"fmt"

	"github.com/bookStore/internal/database"
	"github.com/bookStore/internal/order"
	"github.com/bookStore/internal/product"
	"github.com/bookStore/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetupDB()
	gormdb := database.SetupGormDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("gorm", gormdb)
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

	r.GET("/getOrder", order.HandleGetAllorder)
	r.POST("/AddOrder", order.HandleInsertorder)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
