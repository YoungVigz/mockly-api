package websockets

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	Clients map[int]*websocket.Conn
	Mutex   sync.RWMutex
}

var Manager = ClientManager{
	Clients: make(map[int]*websocket.Conn),
}

func SendToUser(userID int, message []byte) error {
	Manager.Mutex.RLock()
	conn, exists := Manager.Clients[userID]
	Manager.Mutex.RUnlock()

	if !exists {
		return nil
	}

	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil && err != websocket.ErrCloseSent {
		return err
	}
	return nil
}
