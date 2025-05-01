package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("")
	}
	filePath := strings.TrimPrefix(r.URL.Path, "/frontend/")
	fullPath := "frontend/" + filePath
	info, err := os.Stat(fullPath)
	if err!= nil {
		
	}
	if info.IsDir() {
		
	}
	http.ServeFile(w, r, "./"+r.URL.Path)
}
