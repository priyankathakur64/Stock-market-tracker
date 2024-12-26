package main

import (
	"net/http"
	"stock-market-tracker/fetcher"

	"github.com/gin-gonic/gin"
)

func main() {
	apiKey := "HBRVIZ2OE6LLS4U1"
	r := gin.Default()

	r.GET("/stock", func(c *gin.Context) {
		symbol := c.Query("symbol")
		if symbol == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Symbol is required"})
			return
		}

		stock, err := fetcher.FetchStockPrice(apiKey, symbol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, stock)
	})

	r.Run(":8080")
}
