package handlers

import (
	"errors"
	"net/http"
	"ninja/caio/api/components"
	"ninja/caio/api/email"
	rep "ninja/caio/api/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleCheckout(orderRegistrar rep.Registrar[components.Order], userFetcher email.UserFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var purchaseOrder []components.Order
		if err := c.BindJSON(&purchaseOrder); err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "invalid request body",
			})
			return
		}

		userId, err := strconv.Atoi(c.GetHeader("user_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("user_id header is not a valid number")).SetMeta(map[string]interface{}{
				"error": "invalid user_id header",
			})
			return
		}

		// Payment service here (?)
		
		// Loop through the JSON request body
		for _, orderItem := range purchaseOrder {
			var err error
			orderItem.UserId = userId
			// Register the order item
			err = orderRegistrar.Register(orderItem)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
					"error": "error registering product",
				})
				return
			}
		}

		// Send a success email to the user
		userEmail, err := userFetcher.FetchUserById(userId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New("error fetching user")).SetMeta(map[string]interface{}{
				"error": "error fetching user by Id",
			})
			return
		}
		email.SendSuccessEmail(userEmail)

		c.JSON(http.StatusCreated, gin.H{"message": "Purchase order registered successfully"})

		}
	}
