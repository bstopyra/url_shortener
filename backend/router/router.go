package router

import (
	"github.com/ogierhaq/url_shortener/backend/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/URL", middleware.PostLink).Methods("POST", "OPTIONS")

	return router
}