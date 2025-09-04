// Package config handles loading the application configuration from
// environment variables, especially the test cases file, results file, and report file.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the structure containing the paths used by the application.
// TestCasesFile: Path to the CSV file containing the test cases.
// ResultsFile: Path to the CSV file where test results will be stored.
// ReportFile: Path to the HTML file where the test report will be generated.
type Config struct {
	TestCasesFile string // Path to the test cases CSV file
	ResultsFile   string // Path to the results CSV file
	ReportFile    string // Path to the HTML report file
}

// AppConfig is a global instance of the application configuration.
var AppConfig Config

// LoadConfig loads the configuration from environment variables.
// This function loads file paths from a .env file using the godotenv library.
// If any environment variable is not set, the program will terminate with an error.
func LoadConfig() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get test cases file path from the TEST_CASES_FILE environment variable
	testCasesFile := os.Getenv("TEST_CASES_FILE")
	if testCasesFile == "" {
		log.Fatal("TEST_CASES_FILE is not set. Please set this variable in the .env file.")
	}

	// Get results file path from the RESULTS_FILE environment variable
	resultsFile := os.Getenv("RESULTS_FILE")
	if resultsFile == "" {
		log.Fatal("RESULTS_FILE is not set. Please set this variable in the .env file.")
	}

	// Get HTML report file path from the REPORT_FILE environment variable
	reportFile := os.Getenv("REPORT_FILE")
	if reportFile == "" {
		log.Fatal("REPORT_FILE is not set. Please set this variable in the .env file.")
	}

	// Assign file paths to AppConfig struct
	AppConfig = Config{
		TestCasesFile: testCasesFile,
		ResultsFile:   resultsFile,
		ReportFile:    reportFile,
	}

	// Confirm that configuration is loaded correctly by showing file paths
	log.Printf("Configuration loaded successfully. TestCasesFile: %s, ResultsFile: %s, ReportFile: %s", AppConfig.TestCasesFile, AppConfig.ResultsFile, AppConfig.ReportFile)
}
