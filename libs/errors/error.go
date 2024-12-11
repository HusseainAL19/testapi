package libErrors

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type ErrorObj struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func ReturnError(connection *websocket.Conn) {
	returnMsg := ErrorObj{
		Status: 500,
		Msg:    "error during connection",
	}

	if connection == nil {
		return
	}

	connection.WriteJSON(returnMsg)
	writeError := connection.Close()
	if writeError != nil {
		fmt.Println("something went wrong")
	}
}
func ReturnUnAuth(connection *websocket.Conn) {
	returnMsg := ErrorObj{
		Status: 401,
		Msg:    "unautherized for reach this",
	}

	if connection == nil {
		return
	}
	connection.WriteJSON(returnMsg)
	writeError := connection.Close()
	if writeError != nil {
	}

}
