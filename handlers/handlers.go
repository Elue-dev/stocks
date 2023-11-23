package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Elue-dev/stocks/controllers"
	"github.com/Elue-dev/stocks/models"
	"github.com/gorilla/mux"
)


type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("failed to decode json body %v", err)
	}

	result := controllers.InsertStock(stock)

	json.NewEncoder(w).Encode(result)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}

	stock, err :=  controllers.GetStock(int64(id))
	if err != nil {
		log.Fatalf("Could not get stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stocks, err :=  controllers.GetAllStocks()

	if err != nil {
		log.Fatalf("Could not get all stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}

	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("failed to decode json body %v", err)
	}

	updatedRows, err :=  controllers.UpdateStock(int64(id), stock)
	if err != nil {
		log.Fatalf("could not update stock %v", err)
	}

	msg := fmt.Sprintf("Stock updated succeasfully Total rows affected %v:", updatedRows)
	json.NewEncoder(w).Encode(response{
		ID: int64(id),
		Message: msg,
	})
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}
		
	deletedRows, err := controllers.DeleteStock(int64(id))
	if err != nil {
		log.Fatalf("could not delete stock %v", err)
	}

	msg := fmt.Sprintf("Stock deleted succeasfully Total rows affected %v:", deletedRows)
	json.NewEncoder(w).Encode(response{
		ID: int64(id),
		Message: msg,
	})
}