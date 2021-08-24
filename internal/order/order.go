package order

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Code string
	//book  product.Book
	Price string
}

// gormdb := database.SetupGormDB()
// // Migrate the schema
// gormdb.AutoMigrate(&Order{})

// // Create
// gormdb.Create(&Order{Code: "D42", Price: 100})

// // Read
// var product Order
// gormdb.First(&product, 1)                 // find product with integer primary key
// gormdb.First(&product, "code = ?", "D42") // find product with code D42

// // Update - update product's price to 200
// gormdb.Model(&product).Update("Price", 200)
// // Update - update multiple fields
// gormdb.Model(&product).Updates(Order{Price: 200, Code: "F42"}) // non-zero fields
// gormdb.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

// // Delete - delete product
// gormdb.Delete(&product, 1)

func HandleGetAllorder(c *gin.Context) {

	db := c.MustGet("gorm").(*gorm.DB)
	var order Order
	result := db.Find(&order)

	c.JSON(200, gin.H{
		"message": result.RowsAffected,
	})
}

func HandleInsertorder(c *gin.Context) {

	db := c.MustGet("gorm").(*gorm.DB)

	order := Order{}
	order.Code = c.PostForm("code")
	order.Price = c.PostForm("price")

	fmt.Println(order.Code, order.Price)
	orderNew := Order{Code: order.Code, Price: order.Price}

	result := db.Create(&orderNew)

	if order.ID != 0 {
		c.JSON(200, gin.H{
			"message":  "Update Sucess ID :",
			"message1": order.ID,
		})
	} else {
		c.JSON(200, gin.H{
			"message": result.Error,
		})
	}

}
