package api

import (
	"net/http"
	"rt_forum/backend/handlers"
)

func Multiplexer() {
	http.HandleFunc("/ws", handlers.HandleWS)
}