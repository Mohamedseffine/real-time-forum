package api

import (
	"net/http"

	"rt_forum/backend/handlers"
	"rt_forum/backend/models"
)

func Multiplexer() {
	db := models.DatabaseExec()
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWS(w, r, db)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleRegister(w, r, db)
	})
	// http.HandleFunc("/", middleware.IsAlreadyLoggedIn(handlers.HandleRegister, db))
	http.HandleFunc("/frontend/", handlers.HandleStatic)
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleSignUp(w, r, db)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	})

	http.HandleFunc("/create_post", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePostHandler(w, r, db)
	})

	http.HandleFunc("/create_comment", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateCommentHandler(w, r, db)
	})

	http.HandleFunc("/retrieve_posts", func(w http.ResponseWriter, r *http.Request) {
		handlers.RetrievePosts(w, r, db)
	})
}
