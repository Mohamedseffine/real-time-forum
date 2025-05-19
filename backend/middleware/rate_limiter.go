package middleware

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

var ratelimiter = make(chan *http.Request, 100)
var moment = time.Now()

func RateLimit(next func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(ratelimiter) == 100 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "too many requests",
			})
			return
		}
		ratelimiter <- r
		if time.Now().After(moment.Add(time.Hour)) {
			for range ratelimiter {
				<-ratelimiter
			}
		}
		next(w, r, db)
	}
}
