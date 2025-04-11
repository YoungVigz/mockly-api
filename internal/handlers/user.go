package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/YoungVigz/mockly-api/internal/utils"
	"github.com/YoungVigz/mockly-api/internal/validators"
	"github.com/gin-gonic/gin"
)

// @Summary Get user data
// @Description Returns user data based on the provided authentication token.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.UserResponse "User data"
// @Failure 401 {object} models.ErrorResponse "Unauthorized, invalid token"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /user [get]
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
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// @Summary Delete user account
// @Description Deletes the user's account based on the provided authentication token and password.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body models.UserModifyRequest true "Current password"
// @Success 200 {object} string "Successfully deleted account"
// @Failure 400 {object} models.ErrorResponse "Bad request, invalid password"
// @Failure 401 {object} models.ErrorResponse "Unauthorized, invalid token"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /user [delete]
func DeleteUser(c *gin.Context) {

	var userRequest *models.UserModifyRequest = &models.UserModifyRequest{}

	if c.Bind(&userRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read provided values, ensure that your body is Valid",
		})

		return
	}

	_, err := validators.PasswordValidator(userRequest.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valid password is required",
		})

		return
	}

	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	userIdInt, err := utils.ConvertUserIdToInt(userId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	err = userService.DeleteUserById(userIdInt, userRequest)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted account",
	})
}

// @Summary Change user password
// @Description Allows the user to change their password based on the provided current password and new password.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body models.UserChangePassword true "Current and new password"
// @Success 200 {object} string "Password changed successfully"
// @Failure 400 {object} models.ErrorResponse "Bad request, invalid password"
// @Failure 401 {object} models.ErrorResponse "Unauthorized, invalid token"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /user [patch]
func ChangePassword(c *gin.Context) {
	var passwords *models.UserChangePassword = &models.UserChangePassword{}

	if c.Bind(&passwords) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read provided values, ensure that your body is Valid",
		})

		return
	}

	validationMessages, err := validators.PasswordValidator(passwords.NewPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"value":  "newPassword",
			"errors": validationMessages,
		})

		return
	}

	_, err = validators.PasswordValidator(passwords.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"value": "password (used to login with)",
			"error": "Valid password is required",
		})

		return
	}

	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	userIdInt, err := utils.ConvertUserIdToInt(userId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	err = userService.ChangePassword(userIdInt, passwords)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed succesfully",
	})

}
