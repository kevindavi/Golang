package main

import (
	"assignment-2/internal/database"
	"assignment-2/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	handlers.RegisterOrderHandlers(router, db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
