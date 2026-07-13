package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"url-shortener/database"
	"url-shortener/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	database.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{code}", handlers.RedirectURL).Methods("GET")

	router.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	log.Println("🚀 Server running at http://localhost:8081")

	log.Fatal(http.ListenAndServe(":8081", router))
}
