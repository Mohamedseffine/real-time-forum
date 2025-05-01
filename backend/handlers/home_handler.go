package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodGet {
		
	}
	tmp, err:=template.ParseFiles("./frontend/index.html")
	if err != nil {
		log.Println(err)
	}
	tmp.Execute(w, nil)
}