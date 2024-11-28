package handlers

import (
	"errors"
	"net/http"
	comp "ninja/caio/api/components"
	rep "ninja/caio/api/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetAllOrders(fetcher rep.Fetcher[comp.Order]) gin.HandlerFunc {
	return func(c *gin.Context) {
		

		orders, err := fetcher.FetchAll()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to fetch all orders",
			})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

func HandleGetOrder(fetcher rep.Fetcher[comp.Order]) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderIdStr := c.Param("orderId")
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "param id is not a valid number",
			})
			return
		}
		order, err := fetcher.FetchById(orderId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to fetch order by id",
			})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}