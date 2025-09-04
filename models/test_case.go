// Package models defines the structures used in testing, such as test cases in CSV format.
package models

// TestCase represents a test case containing the necessary information to execute an API test.
// The structure maps directly to a CSV file where each column corresponds to a test field.
//
// The structure is used to read test cases from a CSV file and then execute HTTP requests
// according to the provided data.
//
// Fields:
//   - TestId: The unique identifier of the test case (e.g., "TC01").
//   - TestCase: The name or description of the test case.
//   - Run: Indicates whether the test case should be executed ("Y" for yes, "N" for no).
//   - Method: The HTTP method to use (e.g., "GET", "POST").
//   - URL: The base URL to which the endpoint will be appended.
//   - Endpoint: The specific endpoint to append to the base URL.
//   - Authorization: The type of authentication used for the request (e.g., "Bearer").
//   - User: The username for authentication, if required.
//   - Password: The password associated with the username, if required.
//   - Headers: Additional HTTP headers for the request, in JSON format.
//   - Body: The body of the request, sent in cases like "POST" or "PUT".
//   - ExpectedStatusCode: The expected HTTP status code in the response.
//   - ExpectedResponse: The expected API response in JSON format, to compare with the actual response.
type TestCase struct {
	TestId             string `json:"TestId"`             // Test case identifier.
	TestCase           string `json:"TestCase"`           // Name or description of the test case.
	Run                string `json:"Run"`                // Indicates if the test case should run ("Y" or "N").
	Method             string `json:"Method"`             // HTTP method (GET, POST, PUT, DELETE).
	URL                string `json:"URL"`                // Base URL for the request.
	Endpoint           string `json:"Endpoint"`           // Endpoint to append to the base URL.
	Authorization      string `json:"Authorization"`      // Type of authorization (e.g., "Bearer").
	User               string `json:"User"`               // Username for authentication.
	Password           string `json:"Password"`           // Password for authentication.
	Headers            string `json:"Headers"`            // HTTP headers in JSON format.
	Body               string `json:"Body"`               // Request body (for POST, PUT).
	ExpectedStatusCode int    `json:"ExpectedStatusCode"` // Expected HTTP status code in the response.
	ExpectedResponse   string `json:"ExpectedResponse"`   // Expected response in JSON format.
}
