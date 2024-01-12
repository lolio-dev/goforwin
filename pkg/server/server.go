package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}

	defer wsConn.Close()

	for {
		var msg string

		err := wsConn.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}

		fmt.Printf("Message Received: %s\n", msg)
		SendMessage("Hello, Client!")
	}
}

func SendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}

func StartServer() {

}
