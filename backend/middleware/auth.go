package middleware

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/helpers"
)

func IsAlreadyLoggedIn(next func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if helpers.CantLog(db, r) {
			http.SetCookie(w, &http.Cookie{
				Name:   "token",
				Value:  "",
				MaxAge: -1,
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "this is unauthorized",
			})
			return
		}

		next(w, r, db)
	}
}
