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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "you can't log-in",
			})
			return
		}

		next(w, r, db)
	}
}
