package main

import (
	"log"
	"ninja/caio/api/db"
	"ninja/caio/api/email"
	hdl "ninja/caio/api/handlers"
	mdw "ninja/caio/api/middleware"
	rep "ninja/caio/api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Use(mdw.HandleErrors())

	dbConn := db.OpenSqlConnection()
	defer dbConn.Close()

	r.GET("/products", hdl.HandleGetAllProducts(rep.NewProductRepository(dbConn)))
	r.GET("/product/:productId", hdl.HandleGetProduct(rep.NewProductRepository(dbConn)))
	r.POST("/product", hdl.HandleRegisterProduct(rep.NewProductRepository(dbConn)))
	r.PUT("/product/:productId", hdl.HandleUpdateProduct(rep.NewProductRepository(dbConn), rep.NewProductRepository(dbConn)))
	r.DELETE("/product/:productId", hdl.HandleDeleteProduct(rep.NewProductRepository(dbConn)))

	r.POST("/checkout", hdl.HandleCheckout(rep.NewOrderRepository(dbConn), email.NewUserEmailRepository(dbConn)))
	r.GET("/orders", hdl.HandleGetAllOrders(rep.NewOrderRepository(dbConn)))
	r.GET("/order/:orderId", hdl.HandleGetOrder(rep.NewOrderRepository(dbConn)))

	r.Run(":8080")
}
