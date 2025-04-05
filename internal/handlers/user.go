package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	val, exists := c.Get("user_id")

	if !exists || val == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userIDfloat, ok := val.(float64)
	if !ok || userIDfloat <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := userService.GetUserById(int(userIDfloat))

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"code":  customError.Code,
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}
