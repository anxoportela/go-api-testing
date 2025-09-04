// Package csv contains functions for handling CSV files, including reading and writing data.
package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// WriteResults writes the provided results to a CSV file without escaping quotes or other special characters.
// The function creates a new CSV file or overwrites an existing one with the provided data.
//
// Parameters:
//   - results ([][]string): A slice of slices of strings, where each inner slice represents
//     a row to be written to the CSV file.
//   - filename (string): The name of the CSV file where the results will be saved.
//
// Returns:
//   - error: An error in case there is a problem creating or writing to the CSV file.
func WriteResults(results [][]string, filename string) error {
	// Open or create the CSV file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer file.Close()

	// Create a CSV writer that will handle writing the rows
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write each row of results to the CSV file
	for _, record := range results {
		// Write the row as a CSV line
		// NOTE: No automatic escaping of quotes or other special characters is performed,
		//       meaning the data must be properly formatted before writing.
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("error writing to CSV file: %v", err)
		}
	}

	// Return nil if everything went well
	return nil
}
