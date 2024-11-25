package handlers

import (
	"errors"
	"net/http"
	"strconv"

	comp "ninja/caio/api/components"
	rep "ninja/caio/api/repositories"

	"github.com/gin-gonic/gin"
)

func HandleGetAllProducts(fetcher rep.Fetcher[comp.Product]) gin.HandlerFunc {
	return func(c *gin.Context) {
		

		products, err := fetcher.FetchAll()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to fetch all products",
			})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func HandleGetProduct(fetcher rep.Fetcher[comp.Product]) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdStr := c.Param("productId")
		productId, err := strconv.Atoi(productIdStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "param id is not a valid",
			})
			return
		}
		product, err := fetcher.FetchById(productId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to fetch product by id",
			})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func HandleRegisterProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "Product registration successful"})
	}
}

func HandleUpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "update product!!",
		  })
	}
}

func HandleDeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete product!!",
		  })
	}
}