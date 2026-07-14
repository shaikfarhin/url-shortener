package handlers

import (
	"fmt"
	"net/http"

	"url-shortener/database"
	"url-shortener/utils"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	originalURL := r.FormValue("url")

	if originalURL == "" {
		http.Error(w, "URL Required", http.StatusBadRequest)
		return
	}

	shortCode := utils.GenerateShortCode(6)

	query := `
	INSERT INTO urls (original_url, short_code)
	VALUES ($1, $2)
	`

	_, err := database.DB.Exec(query, originalURL, shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Build URL dynamically (works locally and on Render)
	shortURL := fmt.Sprintf("%s://%s/%s",
		getScheme(r),
		r.Host,
		shortCode,
	)

	fmt.Fprintf(w,
		"<h2>✅ URL Shortened Successfully!</h2><br><p>Short URL: <a href='%s'>%s</a></p>",
		shortURL,
		shortURL,
	)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {

	shortCode := r.URL.Path[1:]

	var originalURL string

	query := `
	SELECT original_url
	FROM urls
	WHERE short_code = $1
	`

	err := database.DB.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE urls SET clicks = clicks + 1 WHERE short_code = $1",
		shortCode,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func getScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}

	if r.Header.Get("X-Forwarded-Proto") == "https" {
		return "https"
	}

	return "http"
}
