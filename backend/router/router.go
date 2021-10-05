package router

import (
	"net/http"

	"github.com/ogierhaq/url_shortener/backend/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/URL", middleware.PostLink).Methods("POST", "OPTIONS")
	router.HandleFunc("/{ShortURL}", handleRedirect).Methods("GET")

	return router
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["ShortURL"]
	originURL := middleware.HandleShortenedUrl(shortURL)
	
	http.Redirect(w, r, originURL, http.StatusSeeOther)	
}