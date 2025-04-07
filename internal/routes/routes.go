package routes

import (
	"github.com/YoungVigz/mockly-api/internal/handlers"
	"github.com/YoungVigz/mockly-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	RegisterAuthRoutes(api)
	RegisterSchemaRoutes(api)
}

func RegisterAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)

	auth.GET("/protected", middlewares.AuthGuard(), handlers.GetUser)
}

func RegisterSchemaRoutes(r *gin.RouterGroup) {
	schema := r.Group("/schema")

	schema.POST("/generate", middlewares.AuthGuard(), handlers.GenerateFromSchema)
	schema.POST("/", middlewares.AuthGuard(), handlers.SaveSchema)
	schema.GET("/", middlewares.AuthGuard(), handlers.GetAllUserSchemas)
}
