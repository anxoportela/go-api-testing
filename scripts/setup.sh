#!/bin/bash

# --------------------------------------------------------
# setup.sh: Script to prepare the API testing environment
# --------------------------------------------------------

# Description:
#   This script performs the initial setup to run API tests:
#   - Checks if Go is installed.
#   - Initializes the Go module if it has not been done before.
#   - Downloads the required project dependencies.
#   - Creates a CSV file with basic test cases if it doesn't exist.
#   - Deletes previous result files (data/results.csv and data/report.html) if they exist.
#   - Creates a .env file with necessary environment variables if it doesn't exist.
#
# Usage:
#   Run the script from the terminal with the following command:
#   ./setup.sh
#
# Requirements:
#   - Go must be installed on the system.
#   - You must have write permissions in the directory where the script is run.

# --------------------------------------------------------

# Print start message
echo "Starting API testing environment setup..."

# --------------------------------------------------------
# 1. Ensure Go is installed
# --------------------------------------------------------
if ! command -v go &> /dev/null
then
    echo "Error: Go is not installed. Please install Go."
    exit 1
else
    echo "Go is installed: $(go version)"
fi

# --------------------------------------------------------
# 2. Initialize the Go module (if not done before)
# --------------------------------------------------------
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init go-api-testing
else
    echo "Go module is already initialized."
fi

# --------------------------------------------------------
# 3. Download required dependencies
# --------------------------------------------------------
echo "Downloading Go dependencies..."
go mod tidy

# --------------------------------------------------------
# 4. Check and create the 'data' folder if it doesn't exist
# --------------------------------------------------------
if [ ! -d "data" ]; then
    echo "The 'data' folder does not exist. Creating..."
    mkdir -p data
else
    echo "The 'data' folder already exists."
fi

# --------------------------------------------------------
# 5. Create the CSV file with test cases (if it doesn't exist)
# --------------------------------------------------------
TEST_CASES_FILE="data/test_cases.csv"
if [ ! -f "$TEST_CASES_FILE" ]; then
    echo "The file '$TEST_CASES_FILE' does not exist. Creating file with some basic test cases..."

    # Create CSV file with example tests
    cat <<EOL > $TEST_CASES_FILE
TestId,TestCase,Run,Method,URL,Endpoint,Authorization,User,Password,Headers,Body,ExpectedStatusCode,ExpectedResponse
TC-001,Get first user OK,Y,GET,https://jsonplaceholder.typicode.com,/users/1,,,,,,200,"{""id"":1,""name"":""Leanne Graham"",""username"":""Bret"",""email"":""Sincere@april.biz"",""address"":{""street"":""Kulas Light"",""suite"":""Apt. 556"",""city"":""Gwenborough"",""zipcode"":""92998-3874"",""geo"":{""lat"":""-37.3159"",""lng"":""81.1496""}},""phone"":""1-770-736-8031 x56442"",""website"":""hildegard.org"",""company"":{""name"":""Romaguera-Crona"",""catchPhrase"":""Multi-layered client-server neural-net"",""bs"":""harness real-time e-markets""}}"
TC-002,Get first user KO,Y,GET,https://jsonplaceholder.typicode.com,/users/1,,,,,,200,"{}"
TC-003,Get + StatusCode KO,Y,GET,https://jsonplaceholder.typicode.com,/users/1,,,,,,404,"{}"
TC-004,Invalid endpoint,Y,GET,https://jsonplaceholder.typicode.com,/invalid-endpoint,,,,,,404,"{}"
TC-005,Invalid Status Code,Y,GET,https://jsonplaceholder.typicode.com,/invalid-endpoint,,,,,,200,"{}"
TC-006,Skipped test,N,GET,https://jsonplaceholder.typicode.com,/invalid-endpoint,,,,,,404,"{}"
TC-007,Token Auth,Y,GET,https://httpbin.org,/bearer,Bearer test,,,,,200,"{""authenticated"":true,""token"":""test :""}"
TC-008,POST Example,Y,POST,https://jsonplaceholder.typicode.com,/users,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}",201,"{""id"":11}"
TC-009,POST Error Example,Y,POST,https://jsonplaceholder.typicode.com,/users,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}",201,"{}"
EOL

    echo "File '$TEST_CASES_FILE' created with example test cases."
else
    echo "The file '$TEST_CASES_FILE' already exists."
fi

# --------------------------------------------------------
# 6. Check if previous result files exist, delete if they do
# --------------------------------------------------------
RESULTS_FILE="data/results.csv"
if [ -f "$RESULTS_FILE" ]; then
    echo "The file '$RESULTS_FILE' already exists. Deleting previous file..."
    rm "$RESULTS_FILE"
fi

REPORT_FILE="data/report.html"
if [ -f "$REPORT_FILE" ]; then
    echo "The file '$REPORT_FILE' already exists. Deleting previous file..."
    rm "$REPORT_FILE"
fi

# --------------------------------------------------------
# 7. Create .env file if it doesn't exist
# --------------------------------------------------------
ENV_FILE=".env"
if [ ! -f "$ENV_FILE" ]; then
    echo "The file '$ENV_FILE' does not exist. Creating file with required environment variables..."

    # Create .env file with environment variables
    cat <<EOL > $ENV_FILE
TEST_CASES_FILE=data/test_cases.csv
RESULTS_FILE=data/results.csv
REPORT_FILE=data/report.html
EOL

    echo "File '$ENV_FILE' created with required environment variables."
else
    echo "The file '$ENV_FILE' already exists."
fi

# --------------------------------------------------------
# 8. Display final message
# --------------------------------------------------------
echo "Environment is ready. You can now run the tests with 'go run cmd/main.go'."
