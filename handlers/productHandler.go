package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "products!!",
		  })
	}
}

func HandleProduct2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "one product only!!",
		  })
	}
}