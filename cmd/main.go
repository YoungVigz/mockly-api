// @title           Mockly API
// @version         1.0
// @description     Mockly docs
// @BasePath        /api
// @schemes         http
// @host            localhost:8080
package main

import (
	"fmt"

	"github.com/YoungVigz/mockly-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	/*
		err := database.DatabaseInit()

		if err != nil {
			log.Fatal("Unable to connect to database!")
		}*/

	fmt.Println("Connected to database")

	api := gin.Default()
	routes.RegisterRoutes(api)
	api.Run("0.0.0.0:8080")
}
