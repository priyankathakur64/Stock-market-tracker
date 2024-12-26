package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type StockData struct {
	Symbol string
	Price  float64
	Time   string
}

// InitializeDatabase creates or opens a SQLite database and sets up the table.
func InitializeDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "stocks.db") // Change to "sqlite" instead of "sqlite3"
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Create a table if it doesn't already exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		symbol TEXT NOT NULL,
		price REAL NOT NULL,
		time TEXT NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return db
}

// SaveStockData saves stock data into the database.
func SaveStockData(db *sql.DB, stock StockData) error {
	insertQuery := `
	INSERT INTO stocks (symbol, price, time)
	VALUES (?, ?, ?);`
	_, err := db.Exec(insertQuery, stock.Symbol, stock.Price, stock.Time)
	return err
}

// GetStockHistory retrieves all stored stock data for a specific symbol.
func GetStockHistory(db *sql.DB, symbol string) ([]StockData, error) {
	selectQuery := `
	SELECT symbol, price, time
	FROM stocks
	WHERE symbol = ?;`

	rows, err := db.Query(selectQuery, symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []StockData
	for rows.Next() {
		var stock StockData
		err := rows.Scan(&stock.Symbol, &stock.Price, &stock.Time)
		if err != nil {
			return nil, err
		}
		history = append(history, stock)
	}

	return history, nil
}
