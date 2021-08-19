package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()

// 	r.POST("/customers", insertUser)
// 	r.Run()
// }

func main() {
	r := gin.Default()
	r.POST("/test", exampleFunc)
	r.POST("/testJSON", exampleJSON)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
