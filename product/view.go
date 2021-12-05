package product

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)


func IsEligibleView(db *sql.DB, c *gin.Context) {
	merchantQ := c.DefaultQuery("merchant", "")
	productQ := c.DefaultQuery("product", "")

	product, err := GetProduct(db, merchantQ, productQ)
	if err != nil {
		// Alternative: if you can't find the product, still 200 but false/false
		// aka - the unknown product is considered ineligible :-)
		c.JSON(404, gin.H{
			"merchant": "",
			"product": "",
			"autoship_enabled": false,
			"live": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"merchant": merchantQ,
		"product": productQ,
		"autoship_enabled": product.AutoshipEnabled,
		"live": product.Live,
	})
}
