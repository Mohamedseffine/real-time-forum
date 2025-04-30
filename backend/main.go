package main

import (
	"fmt"
	"log"
	"net/http"
	"rt_forum/backend/api"
)


func  main()  {
	api.Multiplexer()
	server := http.Server{
		Addr: ":8080",
	}
	fmt.Println("Running in : http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}