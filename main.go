package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

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

	// Connect to PostgreSQL
	database.ConnectDB()

	// Create router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{code}", handlers.RedirectURL).Methods("GET")

	// Serve static files
	router.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Get PORT from Render
	port := os.Getenv("PORT")

	// For local development
	if port == "" {
		port = "8081"
	}

	log.Println("🚀 Server running on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
