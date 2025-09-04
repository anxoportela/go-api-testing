@echo off
REM --------------------------------------------------------
REM setup.bat: Script to prepare the API testing environment on Windows
REM --------------------------------------------------------

REM Description:
REM   This script performs the initial setup to run API tests:
REM   - Checks if Go is installed.
REM   - Initializes the Go module if it has not been done before.
REM   - Downloads the required project dependencies.
REM   - Creates a CSV file with basic test cases if it doesn't exist.
REM   - Deletes previous result files (data/results.csv and data/report.html) if they exist.
REM   - Creates a .env file with necessary environment variables if it doesn't exist.
REM
REM Usage:
REM   Run the script from the terminal with the following command:
REM   setup.bat
REM
REM Requirements:
REM   - Go must be installed on the system.
REM   - You must have write permissions in the directory where the script is run.

REM --------------------------------------------------------

REM Print start message
echo Starting API testing environment setup...

REM --------------------------------------------------------
REM 1. Ensure Go is installed
REM --------------------------------------------------------
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Error: Go is not installed. Please install Go.
    exit /b 1
) else (
    echo Go is installed: %go_version%
)

REM --------------------------------------------------------
REM 2. Initialize the Go module (if not done before)
REM --------------------------------------------------------
if not exist go.mod (
    echo Initializing Go module...
    go mod init go-api-testing
) else (
    echo Go module is already initialized.
)

REM --------------------------------------------------------
REM 3. Download required dependencies
REM --------------------------------------------------------
echo Downloading Go dependencies...
go mod tidy

REM --------------------------------------------------------
REM 4. Check and create 'data' folder if it doesn't exist
REM --------------------------------------------------------
if not exist "data" (
    echo The 'data' folder does not exist. Creating...
    mkdir data
) else (
    echo The 'data' folder already exists.
)

REM --------------------------------------------------------
REM 5. Create the CSV file with test cases (if it doesn't exist)
REM --------------------------------------------------------
set TEST_CASES_FILE=data\test_cases.csv
if not exist "%TEST_CASES_FILE%" (
    echo The file "%TEST_CASES_FILE%" does not exist. Creating file with some basic test cases...

    REM Create CSV file with example tests
    echo TestId^,TestCase^,Run^,Method^,URL^,Endpoint^,Authorization^,User^,Password^,Headers^,Body^,ExpectedStatusCode^,ExpectedResponse > %TEST_CASES_FILE%
    echo TC-001^,Get first user OK^,Y^,GET^,https://jsonplaceholder.typicode.com^,/users/1^,,,,,,200^,"{""id"":1,""name"":""Leanne Graham"",""username"":""Bret"",""email"":""Sincere@april.biz"}" >> %TEST_CASES_FILE%
    echo TC-002^,Get first user KO^,Y^,GET^,https://jsonplaceholder.typicode.com^,/users/1^,,,,,,200^,"{}" >> %TEST_CASES_FILE%
    echo TC-003^,Get + StatusCode KO^,Y^,GET^,https://jsonplaceholder.typicode.com^,/users/1^,,,,,,404^,"{}" >> %TEST_CASES_FILE%
    echo TC-004^,Invalid endpoint^,Y^,GET^,https://jsonplaceholder.typicode.com^,/invalid-endpoint^,,,,,,404^,"{}" >> %TEST_CASES_FILE%
    echo TC-005^,Invalid Status Code^,Y^,GET^,https://jsonplaceholder.typicode.com^,/invalid-endpoint^,,,,,,200^,"{}" >> %TEST_CASES_FILE%
    echo TC-006^,Skipped test^,N^,GET^,https://jsonplaceholder.typicode.com^,/invalid-endpoint^,,,,,,404^,"{}" >> %TEST_CASES_FILE%
    echo TC-007^,Token Auth^,Y^,GET^,https://httpbin.org^,/bearer^,Bearer test^,,,,200^,"{""authenticated"":true,""token"":""test :""}" >> %TEST_CASES_FILE%
    echo TC-008^,POST Example^,Y^,POST^,https://jsonplaceholder.typicode.com^,/users^,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}"^,201^,"{""id"":11}" >> %TEST_CASES_FILE%
    echo TC-009^,POST Error Example^,Y^,POST^,https://jsonplaceholder.typicode.com^,/users^,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}"^,201^,"{}" >> %TEST_CASES_FILE%

    echo File "%TEST_CASES_FILE%" created with example test cases.
) else (
    echo The file "%TEST_CASES_FILE%" already exists.
)

REM --------------------------------------------------------
REM 6. Check if previous result files exist and delete them
REM --------------------------------------------------------
set RESULTS_FILE=data\results.csv
if exist "%RESULTS_FILE%" (
    echo The file "%RESULTS_FILE%" already exists. Deleting previous file...
    del "%RESULTS_FILE%"
)

set REPORT_FILE=data\report.html
if exist "%REPORT_FILE%" (
    echo The file "%REPORT_FILE%" already exists. Deleting previous file...
    del "%REPORT_FILE%"
)

REM --------------------------------------------------------
REM 7. Create .env file if it doesn't exist
REM --------------------------------------------------------
set ENV_FILE=.env
if not exist "%ENV_FILE%" (
    echo The file "%ENV_FILE%" does not exist. Creating file with required environment variables...

    REM Create .env file with environment variables
    echo TEST_CASES_FILE=data\test_cases.csv >> %ENV_FILE%
    echo RESULTS_FILE=data\results.csv >> %ENV_FILE%
    echo REPORT_FILE=data\report.html >> %ENV_FILE%

    echo File "%ENV_FILE%" created with required environment variables.
) else (
    echo The file "%ENV_FILE%" already exists.
)

REM --------------------------------------------------------
REM 8. Display final message
REM --------------------------------------------------------
echo The environment is ready. You can now run the tests with 'go run cmd\main.go'.
pause
