package main

import (
	"ninja/caio/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/product", handlers.HandleProduct())
  r.GET("/product/:productId", handlers.HandleProduct2())

  r.Run(":8080")
}