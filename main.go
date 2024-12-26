package main

import (
	"fmt"
	"log"
	"stock-market-tracker/fetcher"
	"stock-market-tracker/notifier"
	"stock-market-tracker/storage"
)

func main() {
	apiKey := "HBRVIZ2OE6LLS4U1"
	symbol := "AAPL"
	threshold := 150.0

	// Initialize the database
	db := storage.InitializeDatabase()
	defer db.Close()

	// Fetch stock data
	stock, err := fetcher.FetchStockPrice(apiKey, symbol)
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
	}

	fmt.Printf("Stock: %s\nPrice: $%.2f\nTime: %s\n", stock.Symbol, stock.Price, stock.Time)

	// Convert fetcher.StockData to storage.StockData
	storageStock := storage.StockData{
		Symbol: stock.Symbol,
		Price:  stock.Price,
		Time:   stock.Time,
	}

	// Save data to the database
	err = storage.SaveStockData(db, storageStock)
	if err != nil {
		log.Fatalf("Error saving stock data: %v", err)
	}
	fmt.Println("Stock data saved to the database.")

	// Trigger alerts
	notifier.Notify(stock.Symbol, stock.Price, threshold)

	// Retrieve and print stock history
	history, err := storage.GetStockHistory(db, symbol)
	if err != nil {
		log.Fatalf("Error retrieving stock history: %v", err)
	}

	fmt.Printf("Stock History for %s:\n", symbol)
	for _, record := range history {
		fmt.Printf("Price: $%.2f, Time: %s\n", record.Price, record.Time)
	}
}
