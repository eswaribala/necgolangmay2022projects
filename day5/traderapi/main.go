package main

import (
	"github.com/gorilla/mux"
	"github.com/necws/day5/traderapi/stores"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	// Create
	router.HandleFunc("/traders", stores.CreateTrader).Methods("POST")
	// Read
	router.HandleFunc("/traders/{traderId}", stores.GetTrader).Methods("GET")
	// Read-all
	router.HandleFunc("/traders", stores.GetTraders).Methods("GET")
	// Update
	router.HandleFunc("/traders/{traderId}", stores.UpdateTrader).Methods("PUT")
	// Delete
	router.HandleFunc("/traders/{traderId}", stores.DeleteTrader).Methods("DELETE")
	// Initialize db connection
	stores.InitDB()

	log.Fatal(http.ListenAndServe(":7070", router))

	// Swagger
	//router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":7070", router))
}