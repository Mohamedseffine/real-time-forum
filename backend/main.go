package main

import (
	"fmt"
	"log"
	"net/http"

	"rt_forum/backend/api"

	"github.com/gorilla/websocket"
)

func main() {
	api.Multiplexer()
	server := http.Server{
		Addr: ":8080",
	}
	upgrader :=websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	upgrader.Upgrade(nil, nil, nil)
	fmt.Println("Running in : http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
