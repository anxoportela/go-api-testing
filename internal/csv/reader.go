// Package csv contains functions to read CSV files and convert them into data structures.
package csv

import (
	"encoding/csv"
	"go-api-testing/models"
	"os"
	"strconv"
)

// ReadCSV reads test cases from a CSV file located at the specified path.
// The function converts each line of the file (except the first header line) into a
// TestCase structure and returns a slice of these structures.
//
// Parameters:
//   - path (string): The path to the CSV file containing the test cases.
//
// Returns:
//   - []models.TestCase: A slice of TestCase structures representing the test cases.
//   - error: An error in case there is an issue opening or reading the CSV file.
func ReadCSV(path string) ([]models.TestCase, error) {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	// Read all lines from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Convert CSV records into TestCase structures
	var testCases []models.TestCase
	for _, record := range records[1:] { // Skip the first line (headers)
		testCases = append(testCases, models.TestCase{
			TestId:             record[0],
			TestCase:           record[1],
			Run:                record[2],
			Method:             record[3],
			URL:                record[4],
			Endpoint:           record[5],
			Authorization:      record[6],
			User:               record[7],
			Password:           record[8],
			Headers:            record[9],
			Body:               record[10],
			ExpectedStatusCode: atoi(record[11]),
			ExpectedResponse:   record[12],
		})
	}

	// Return the slice of test cases and no error
	return testCases, nil
}

// atoi converts a string to an integer, handling conversion errors.
//
// Parameters:
//   - s (string): The string to convert to an integer.
//
// Returns:
//   - int: The integer value corresponding to the provided string.
//   - If conversion fails, returns 0.
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
