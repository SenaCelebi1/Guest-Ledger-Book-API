package main

import (
	router "bookapi/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	log.Fatal(http.ListenAndServe(":8080", r))
}
