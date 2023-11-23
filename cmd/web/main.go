package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Elue-dev/stocks/router"
	"github.com/joho/godotenv"
)

func main() {
	r := router.Router()
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	fmt.Println("Go server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}