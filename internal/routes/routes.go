package routes

import (
	"github.com/YoungVigz/mockly-api/internal/handlers"
	"github.com/YoungVigz/mockly-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	ws := r.Group("ws")

	RegisterAuthRoutes(api)
	RegisterSchemaRoutes(api)
	RegisterUserRoutes(api)

	RegisterWebSocketRoutes(ws)
}

func RegisterAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/", middlewares.AuthGuard(), handlers.GetUser)
	user.DELETE("/", middlewares.AuthGuard(), handlers.DeleteUser)
	user.PATCH("/", middlewares.AuthGuard(), handlers.ChangePassword)
}

func RegisterSchemaRoutes(r *gin.RouterGroup) {
	schema := r.Group("/schema")

	schema.POST("/generate", middlewares.AuthGuard(), handlers.GenerateFromSchema)
	schema.POST("/", middlewares.AuthGuard(), handlers.SaveSchema)
	schema.GET("/", middlewares.AuthGuard(), handlers.GetAllUserSchemas)
	schema.GET("/:title", middlewares.AuthGuard(), handlers.GetUserSchemaByTitle)
	//TODO:
	//schema.PUT("/:title", middlewares.AuthGuard(), handlers.ChangeTitleAndContent)
	//schema.DELETE("/:title", middlewares.AuthGuard(), handlers.DeleteUserSchema)
}

func RegisterWebSocketRoutes(r *gin.RouterGroup) {
	r.GET("/", middlewares.AuthGuard(), handlers.WebSocketServer)
}
