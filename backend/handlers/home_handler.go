package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
)

func HandleRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}
	tmp, err := template.ParseFiles("./frontend/index.html")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "can't serve page",
		})
		return
	}
	tmp.Execute(w, nil)
}
