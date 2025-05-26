package helpers

import (
	"database/sql"

	"net/http"
	"regexp"
	"strings"
	"time"

	"rt_forum/backend/models"
)

func IsValidUesrname(username string) bool {
	if len(username) < 2 {
		return false
	}
	for _, val := range username {
		if (val < 'a' || val > 'z') && (val < '0' || val > '9') && val != '_' {
			return false
		}
	}
	return true
}

func IsvalidName(name string) bool {
	if len(name) < 2 {
		return false
	}
	for _, val := range name {
		if (val < 'a' || val > 'z') && (val < 'A' || val > 'Z') {
			return false
		}
	}
	return true
}

func IsValidEmail(email string) bool {
	if len(email) < 5 {
		return false
	}
	reg, err := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(email))
}

func IsLoggedIn(db *sql.DB, r *http.Request) bool {
	token, err := r.Cookie("token")
	if err != nil {
		return !(err.Error() == "http: named cookie not present")
	}
	tok := strings.TrimPrefix(token.String(), "token=")
	n := models.CheckSession(db, tok)
	if n != 1 {
		return true
	}
	expires_at, err := models.IsExpired(db, tok)
	if err != nil {

		return true
	}
	return time.Now().After(expires_at)
}
