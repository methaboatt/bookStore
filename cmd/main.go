package main

import (
	"github.com/bookStore/internal/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := server.SetupDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	defer db.Close()

	r.POST("/testAddUser", server.HandleInsertUser)
	r.GET("/testGetUser", server.HandleGetUser)
	r.PUT("/testUpdateUser", server.HandleUpdateUser)
	r.PATCH("/testPatchUser", server.HandlePatchUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
