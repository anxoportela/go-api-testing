package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // Pure Go SQLite driver, no cgo required
)

var DB *sql.DB

// InitDB initializes the SQLite database and creates the table if it does not exist
func InitDB(path string) {
	var err error
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS test_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		test_id TEXT,
		test_case TEXT,
		result BOOLEAN,
		message TEXT,
		run_date DATETIME
	);
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

// SaveResult saves a test result in the database
func SaveResult(testId, testCase string, result bool, message string) error {
	_, err := DB.Exec(`
		INSERT INTO test_results (test_id, test_case, result, message, run_date)
		VALUES (?, ?, ?, ?, ?)
	`, testId, testCase, result, message, time.Now())
	if err != nil {
		return fmt.Errorf("error saving result in DB: %v", err)
	}
	return nil
}

// GetHistory retrieves all historical test results
func GetHistory() ([]map[string]interface{}, error) {
	rows, err := DB.Query(`
		SELECT test_id, test_case, result, message, run_date
		FROM test_results
		ORDER BY run_date DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("error retrieving history from DB: %v", err)
	}
	defer rows.Close()

	var history []map[string]interface{}
	for rows.Next() {
		var testId, testCase, message string
		var result bool
		var runDate string
		if err := rows.Scan(&testId, &testCase, &result, &message, &runDate); err != nil {
			return nil, err
		}
		history = append(history, map[string]interface{}{
			"test_id":   testId,
			"test_case": testCase,
			"result":    result,
			"message":   message,
			"run_date":  runDate,
		})
	}
	return history, nil
}
