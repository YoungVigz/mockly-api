package routes

import (
	"github.com/YoungVigz/mockly-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	RegisterAuthRoutes(api)
}

func RegisterAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)
}
