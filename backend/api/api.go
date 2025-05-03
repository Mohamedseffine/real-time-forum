package api

import (
	"net/http"
	"rt_forum/backend/handlers"
)

func Multiplexer() {
	http.HandleFunc("/chat", handlers.HandleWS)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/frontend/", handlers.HandleStatic)
	http.HandleFunc("/signup", handlers.HandleSignUp)
}