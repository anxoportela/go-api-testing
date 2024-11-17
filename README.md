<!-- omit from toc -->
# 🧪 **Go API Testing Framework**

This repository contains a **Go-based API testing framework** designed to automate the testing of APIs using predefined test cases stored in CSV files. It supports reading and writing results to CSV, generating HTML reports, and executing test cases using configurable API requests.

<!-- omit from toc -->
## 📑 Table of Contents

- [🧑‍💻 **Overview**](#-overview)
- [🗂️ **Project Structure**](#️-project-structure)
- [🛠️ **Setup and Installation**](#️-setup-and-installation)
- [🚀 **Usage**](#-usage)
- [🤝 **Contributing**](#-contributing)

---

## 🧑‍💻 **Overview**

This project automates the testing of APIs by:

- 🔄 Loading test cases from a CSV file.
- 📡 Making HTTP requests based on the test cases.
- ✅ Verifying the response status and content.
- 💾 Saving results into CSV files.
- 📊 Generating HTML reports for test results.

---

## 🗂️ **Project Structure**

```bash
go-api-testing/
│
├── cmd/                        # Main execution code
│   └── main.go                 # Entry point for running tests
│
├── config/                     # Configuration files
│   └── config.go               # Global configurations (e.g., CSV reading)
├── data/
│   └── results.csv             # Test results in CSV format
│   └── test_cases.csv          # CSV file with test cases
│   └── report.html             # HTML report of test results
│
├── internal/                   # Internal logic for test execution
│   ├── api/                    # Logic for making API requests
│   │   └── client.go           # Functions for sending HTTP requests
│   ├── csv/                    # Logic for reading/writing CSV files
│   │   └── reader.go           # Functions for reading CSV files
│   │   └── writer.go           # Functions for writing results to CSV
│   ├── report/                 # Logic for generating HTML reports
│   │   └── html.go             # Functions for generating HTML reports
│   └── test/                   # Test execution logic
│       └── executor.go         # Functions to execute test cases
│
├── models/                     # Data structures for test cases
│   └── test_case.go            # Struct for test case data
│
├── scripts/                    # Setup and utility scripts
│   └── setup.bat               # Windows setup script
│   └── setup.sh                # Linux/Mac setup script
│
├── go.mod                      # Go module file with dependencies
├── go.sum                      # Go dependencies checksum
└── README.md                   # This file
```

---

## 🛠️ **Setup and Installation**

Before running the tests, you need to set up the required environment and dependencies. You can use either the `setup.sh` script for Linux/Mac or the `setup.bat` script for Windows.

<!-- omit from toc -->
### 🌍 **Linux/Mac Setup (using `setup.sh`)**

#### **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/go-api-testing.git
   cd go-api-testing
   ```

#### **Make the `setup.sh` script executable**

   ```bash
   chmod +x scripts/setup.sh
   ```

#### **Run the setup script**

   This script will install all required dependencies, create the `.env` file, and generate the `data/test_cases.csv` file.

   ```bash
   ./scripts/setup.sh
   ```

   The `setup.sh` script will:

- Install the necessary Go dependencies.
- Create a `.env` file with the necessary environment variables.
- Generate a sample `test_cases.csv` file in the `data/` directory.

#### **Verify Setup**

   Once the setup script completes, you should have the following files:

- `.env` (with paths for the test cases, results, and report files).
- `data/test_cases.csv` (predefined test cases to be executed).
- All required Go dependencies should be installed.

---

<!-- omit from toc -->
### 🖥️ **Windows Setup (using `setup.bat`)**

#### **Clone the repository**

   ```cmd
   git clone https://github.com/yourusername/go-api-testing.git
   cd go-api-testing
   ```

#### **Run the `setup.bat` script**

   This script will install all required dependencies, create the `.env` file, and generate the `data/test_cases.csv` file.

   ```cmd
   scripts\setup.bat
   ```

   The `setup.bat` script will:

- Install the necessary Go dependencies.
- Create a `.env` file with the necessary environment variables.
- Generate a sample `test_cases.csv` file in the `data/` directory.

#### **Verify Setup**

   After the script finishes, confirm that the following files have been created:

- `.env` (with paths for the test cases, results, and report files).
- `data/test_cases.csv` (contains sample test cases for API testing).

---

## 🚀 **Usage**

<!-- omit from toc -->
### **Prepare Test Cases**

   Edit the `data/test_cases.csv` file to include the test cases you wish to run. Each row should represent a test case with the following columns:

- `TestId`: Unique identifier for the test case.
- `TestCase`: A description of the test case.
- `Run`: "Y" to run the test, "N" to skip.
- `Method`: HTTP method (GET, POST, PUT, DELETE).
- `URL`: Base URL for the API.
- `Endpoint`: The specific endpoint to test.
- `Authorization`: Type of authorization (e.g., Bearer token).
- `User`: Username for authentication (if needed).
- `Password`: Password for authentication (if needed).
- `Headers`: JSON formatted headers (if any).
- `Body`: Request body (for POST/PUT).
- `ExpectedStatusCode`: The expected HTTP status code (e.g., 200).
- `ExpectedResponse`: Expected JSON response body (if applicable).

<!-- omit from toc -->
### **Run the Tests**

   After setting up the `.env` file and `test_cases.csv`, you can run the tests by executing the `main.go` file.

   ```bash
   go run cmd/main.go
   ```

   The tests will be executed, and the results will be printed in the terminal and saved to:

- `data/results.csv` (CSV file with test results).
- `data/report.html` (HTML report of the test results).

---

## 🤝 **Contributing**

If you'd like to contribute to this project, feel free to fork the repository, create a branch, and submit a pull request. All contributions are welcome! 🎉
