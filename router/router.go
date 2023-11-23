package router

import (
	"github.com/Elue-dev/stocks/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stocks/{id}", handlers.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks", handlers.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks", handlers.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stocks/{id}", handlers.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stocks/{id}", handlers.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}