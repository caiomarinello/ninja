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
				"error": "param id is not a valid number",
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

func HandleUpdateProduct(updater rep.Updater[comp.Product], fetcher rep.Fetcher[comp.Product]) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdStr := c.Param("productId")
		productId, err := strconv.Atoi(productIdStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "param id is not a valid number",
			})
			return
		}
		product, err := fetcher.FetchById(productId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "product update failed to fetch product by id",
			})
			return
		}

		var updatedProductAttributes map[string]interface{}
		if err := c.BindJSON(&updatedProductAttributes); err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "invalid request body",
			})
			return
		}

		var ok bool
		// Update the product with the provided attributes
		for attribute, value := range updatedProductAttributes {
			switch attribute {
			case "product_name":
				product.Name, ok = value.(string)
				if !ok {
					c.AbortWithError(http.StatusBadRequest, errors.New("tried to alter product name with type different than string")).SetMeta(map[string]interface{}{
						"error": "invalid request body",
					})
					return
				}
			case "product_description":
				product.Description, ok = value.(string)
				if !ok {
					c.AbortWithError(http.StatusBadRequest, errors.New("tried to alter product description with type different than string")).SetMeta(map[string]interface{}{
						"error": "invalid request body",
					})
					return
				}
			case "product_price":
				product.Price, ok = value.(float64)
				if !ok {
					c.AbortWithError(http.StatusBadRequest, errors.New("tried to alter product price with type different than float64")).SetMeta(map[string]interface{}{
						"error": "invalid request body",
					})
					return
				}
			case "product_category":
				product.Category, ok = value.(string)
				if !ok {
					c.AbortWithError(http.StatusBadRequest, errors.New("tried to alter product category with type different than string")).SetMeta(map[string]interface{}{
						"error": "invalid request body",
					})
					return
				}
			case "product_stock":
				// Assert the value as float64 (JSON decodes numbers as float64 by default)
				stockFloat, ok := value.(float64)
				if !ok {
					c.AbortWithError(http.StatusBadRequest, errors.New("tried to alter product stock with type different than number")).SetMeta(map[string]interface{}{
						"error": "invalid request body",
					})
					return
				}

				// Convert float64 to int
				if stockFloat != float64(int(stockFloat)) {
					c.AbortWithError(http.StatusBadRequest, errors.New("tried to alter product stock with a non-integer value")).SetMeta(map[string]interface{}{
						"error": "invalid request body",
					})
					return
				}

				product.Stock = int(stockFloat)
			}
		}

		if err := updater.Update(product); err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to update product by product",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
	}
}

func HandleDeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete product!!",
		  })
	}
}