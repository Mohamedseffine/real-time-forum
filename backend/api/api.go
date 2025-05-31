package api

import (
	"net/http"

	"rt_forum/backend/handlers"
	"rt_forum/backend/middleware"
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

	http.HandleFunc("/create_post",middleware.IsAlreadyLoggedIn(handlers.CreatePostHandler, db))

	http.HandleFunc("/create_comment", middleware.IsAlreadyLoggedIn(handlers.CreateCommentHandler, db))

	http.HandleFunc("/retrieve_posts", middleware.IsAlreadyLoggedIn(handlers.RetrievePosts, db))

	http.HandleFunc("/logout", middleware.IsAlreadyLoggedIn(handlers.LogoutHandler, db))
	http.HandleFunc("/retrieve_comments", middleware.IsAlreadyLoggedIn(handlers.RetrieveComments, db))
	http.HandleFunc("/get_chat", middleware.IsAlreadyLoggedIn(handlers.GetChatMessages, db))
}
