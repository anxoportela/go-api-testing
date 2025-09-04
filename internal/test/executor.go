// Package test contains functions related to executing tests and comparing API responses
// against the expected responses defined in test cases.
package test

import (
	"encoding/json"
	"fmt"
	"go-api-testing/internal/api"
	"go-api-testing/models"
	"reflect"
)

// RunTest executes a test based on a specified test case.
// The function performs an HTTP request using the test case data,
// then compares the obtained response with the expected one.
// If both the status code and response body match the expectations, the test passes.
//
// Parameters:
//   - test (models.TestCase): Test case containing method, URL, headers, body, auth, and expected responses.
//
// Returns:
//   - bool: true if the test was successful, false otherwise.
//   - string: Detailed message about the test result.
func RunTest(test models.TestCase) (bool, string) {
	fullURL := test.URL + test.Endpoint

	// Perform HTTP request
	statusCode, response, err := api.RealizarSolicitud(
		test.Method,
		fullURL,
		test.Body,
		test.Headers,
		test.Authorization,
		test.User,
		test.Password,
	)

	if err != nil {
		// Request or authentication error
		msg := fmt.Sprintf("Error in request: %v", err)
		return false, msg
	}

	// Check status code
	if statusCode != test.ExpectedStatusCode {
		msg := fmt.Sprintf(
			"Incorrect status code: expected %d, got %d",
			test.ExpectedStatusCode,
			statusCode,
		)
		return false, msg
	}

	// Check expected response only if it's not empty
	if test.ExpectedResponse != "" {
		var expected, actual map[string]interface{}

		// Deserialize expected response
		if err := json.Unmarshal([]byte(test.ExpectedResponse), &expected); err != nil {
			msg := fmt.Sprintf("Error deserializing expected response: %v", err)
			return false, msg
		}

		// Deserialize obtained response
		if err := json.Unmarshal([]byte(response), &actual); err != nil {
			msg := fmt.Sprintf("Error deserializing obtained response: %v", err)
			return false, msg
		}

		// Compare JSON structures
		if !reflect.DeepEqual(expected, actual) {
			expectedJSON, _ := json.MarshalIndent(expected, "", "  ")
			actualJSON, _ := json.MarshalIndent(actual, "", "  ")
			msg := fmt.Sprintf(
				"Response does not match:\nExpected: %s\nObtained: %s",
				expectedJSON, actualJSON,
			)
			return false, msg
		}
	}

	// Everything is correct
	return true, "Test passed successfully"
}
