package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketServer(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	userIdInt, err := utils.ConvertUserIdToInt(userId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	user, err := userService.GetUserById(userIdInt)

	if err != nil {
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Hello "+user.Username))
}
