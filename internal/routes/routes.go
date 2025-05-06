package routes

import (
	_ "github.com/YoungVigz/mockly-api/internal/docs"
	"github.com/YoungVigz/mockly-api/internal/handlers"
	"github.com/YoungVigz/mockly-api/internal/middlewares"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	ws := r.Group("ws")

	api.GET("/docs/*any", ginSwagger.WrapHandler(files.Handler))
	api.GET("/status", handlers.CheckHealth)

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
	schema.DELETE("/:title", middlewares.AuthGuard(), handlers.DeleteUserSchema)
	schema.GET("/download/:id", handlers.DownloadSchema)
}

func RegisterWebSocketRoutes(r *gin.RouterGroup) {
	r.GET("/", middlewares.AuthGuard(), handlers.WebSocketServer)
}

func SetupTestRoutes() *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r)
	return r
}
