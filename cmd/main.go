package main

import (
	"fmt"
	"go-api-testing/config"
	"go-api-testing/internal/csv"
	"go-api-testing/internal/db"
	"go-api-testing/internal/report"
	"go-api-testing/internal/test"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize SQLite
	db.InitDB("data/test_history.db")

	// Read test cases from CSV
	testCases, err := csv.ReadCSV(config.AppConfig.TestCasesFile)
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Console table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TestId", "TestCase", "Result", "Message"})

	results := [][]string{{"TestId", "TestCase", "Result", "Message"}}

	// Execute tests
	for _, tc := range testCases {
		if tc.Run == "Y" {
			success, message := test.RunTest(tc)
			consoleMsg := "Test failed"
			if success {
				consoleMsg = "Test successful"
			}
			row := []string{tc.TestId, tc.TestCase, fmt.Sprintf("%t", success), consoleMsg}
			table.Append(row)
			results = append(results, []string{tc.TestId, tc.TestCase, fmt.Sprintf("%t", success), message})

			// Save to DB
			if err := db.SaveResult(tc.TestId, tc.TestCase, success, message); err != nil {
				log.Printf("Error saving to DB: %v", err)
			}
		}
	}

	// Show table in console
	table.Render()

	// Save CSV
	if err := csv.WriteResults(results, config.AppConfig.ResultsFile); err != nil {
		log.Fatalf("Error writing CSV: %v", err)
	}

	// Get full history
	historico, err := db.GetHistory()
	if err != nil {
		log.Fatalf("Error getting DB history: %v", err)
	}

	// Generate HTML report with history
	if err := report.GenerateUltimateReport(results, historico, config.AppConfig.ReportFile); err != nil {
		log.Fatalf("Error generating HTML report: %v", err)
	}

	fmt.Println("Tests executed. CSV and HTML report generated.")
}
