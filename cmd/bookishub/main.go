package main

import (
	//"fmt"
	"os"
	"net/http"
	"html/template"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	//"github.com/joho/godotenv"
)

func getEnvVariable(path string) {
	
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../web/templates/layout.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	})
	http.ListenAndServe("localhost:7000", r)
}
