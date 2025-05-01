package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)
var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func HandleWS(w http.ResponseWriter, r *http.Request){
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"error":"can't upgrage the protocol",
		})
	}
}