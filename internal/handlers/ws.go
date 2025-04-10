package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/utils"
	"github.com/YoungVigz/mockly-api/internal/websockets"
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

	websockets.Manager.Mutex.Lock()
	websockets.Manager.Clients[userIdInt] = conn
	websockets.Manager.Mutex.Unlock()

	go func(conn *websocket.Conn, userID int) {
		defer func() {
			websockets.Manager.Mutex.Lock()
			delete(websockets.Manager.Clients, userID)
			websockets.Manager.Mutex.Unlock()
			conn.Close()
		}()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}(conn, userIdInt)
}
