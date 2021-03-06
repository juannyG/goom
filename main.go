package main

import (
	"github.com/gin-gonic/gin"

	"goom/db"
	"goom/product"
)

func main() {
	database := db.Connect()
	defer database.Close()

	router := gin.Default()
	router.GET("/product", func(c *gin.Context) {
		product.IsEligibleView(database, c)
	})
	router.Run()
}
