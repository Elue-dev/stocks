package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Elue-dev/stocks/connections"
	"github.com/Elue-dev/stocks/models"
)

func InsertStock(s models.Stock) models.Stock {
	fmt.Println("STOCK", s)
	db := connections.CeateConnection()
	defer db.Close()

    sqlQuery := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING *`
	var stock models.Stock

	err := db.QueryRow(sqlQuery, s.Name, s.Price, s.Company).Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	if err != nil {
		log.Fatalf("Could not execute SQL query %v", err)
	}

	return stock
}


func GetAllStocks() ([]models.Stock, error){
	db := connections.CeateConnection()
	defer db.Close()

	var stocks []models.Stock

	sqlQuery := `SELECT * FROM stocks`

	rows, err := db.Query(sqlQuery)

	if err != nil {
		log.Fatalf("Could not execute SQL query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Could not scan rows %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

func GetStock(id int64) (models.Stock, error) {
	db := connections.CeateConnection()
	defer db.Close()

	var stock models.Stock

	sqlQuery := `SELECT * FROM stocks WHERE stockid = $1`

	row := db.QueryRow(sqlQuery, id)

	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned.")
			return stock, nil
		case nil:
			return stock, nil
		default:
			log.Fatalf("Unable to scan rows %v", err)
	}

	return stock, err
}

func UpdateStock(id int64, s models.Stock) (int64, error) {
	db := connections.CeateConnection()
	defer db.Close()

	sqlQuery := `UPDATE stocks SET name = $2, price = $3, company = $4 WHERE stockid = $1`

	res, err := db.Exec(sqlQuery, id, s.Name, s.Price, s.Company)
	if err != nil {
		log.Fatalf("Could not execute query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Could not get affected rows %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)

	return rowsAffected, err

}

func DeleteStock(id int64) (int64, error) {
	db := connections.CeateConnection()
	defer db.Close()

	sqlQuery := `DELETE FROM stocks WHERE stockid = $1`

	res, err := db.Exec(sqlQuery, id)

	if err != nil {
		log.Fatalf("Could not execute query %v", err)
	}
	
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Could not get affected rows %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)

	return rowsAffected, err
}
