package api

import (
	"net/http"

	"rt_forum/backend/models"
	"rt_forum/backend/handlers"
)

func Multiplexer() {
	db := models.DatabaseExec()
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWS(w, r, db)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, db)
	})
	http.HandleFunc("/frontend/", handlers.HandleStatic)
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleSignUp(w, r, db)
	})
}
