package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}
	filePath := strings.TrimPrefix(r.URL.Path, "/frontend/")
	fullPath := "frontend/" + filePath
	info, err := os.Stat(fullPath)
	if err != nil {
		if err == os.ErrNotExist {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "ressorce not found",
			})
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "internal server error",
		})
		return
	}
	if info.IsDir() {
		if err == os.ErrNotExist {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "you van't acces this ressources",
			})
			return
		}
	}
	http.ServeFile(w, r, "./"+r.URL.Path)
}
