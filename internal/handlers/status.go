package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/websockets"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Status string `json:"status"`
}

type WebsocketService struct {
	Status      string `json:"status"`
	ActiveUsers int    `json:"activeUsers"`
}

type Services struct {
	API       Service          `json:"api"`
	Database  Service          `json:"database"`
	Websocket WebsocketService `json:"websocket"`
}

// @Summary Check system health
// @Description Returns the health status of the API, Database, and Websocket services.
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {object} Services "Health check response with services status"
// @Failure 500 {object} models.ErrorResponse "API dose not respond thus every service is unhealthy"
// @Router /status [get]
func CheckHealth(c *gin.Context) {

	websockets.Manager.Mutex.RLock()
	activeUsers := len(websockets.Manager.Clients)
	websockets.Manager.Mutex.RUnlock()

	services := &Services{
		API: Service{
			Status: "Healthy",
		},
		Database: Service{
			Status: "Healthy",
		},
		Websocket: WebsocketService{
			Status:      "Healthy",
			ActiveUsers: activeUsers,
		},
	}

	_, err := database.GetDB()

	if err != nil {
		services.Database.Status = "Unhealthy"
	}

	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}
