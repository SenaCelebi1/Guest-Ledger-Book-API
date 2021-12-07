package router

import (
	middleware "bookapi/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/mails", middleware.GuestLedger).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/mails", middleware.CreateMail).Methods("POST", "OPTIONS")

	return router
}
