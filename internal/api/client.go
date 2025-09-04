// Package api contains functions for making HTTP requests and handling their responses.
package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// RealizarSolicitud makes an HTTP request using the specified method, URL, and body.
// It also adds headers and authentication credentials if provided.
// The function returns the HTTP status code, response body as string, and any error if the request fails.
//
// Parameters:
//   - method (string): HTTP method (e.g., "GET", "POST").
//   - url (string): URL to send the request.
//   - body (string): Request body in JSON format, if any.
//   - headers (string): JSON string representing additional headers to add.
//   - auth (string): Authentication type ("Bearer" or "Basic").
//   - user (string): Username or token for authentication.
//   - password (string): Password for authentication, if needed.
//
// Returns:
//   - int: HTTP status code of the response.
//   - string: Response body as string.
//   - error: Error if any occurs during the request or response processing.
func RealizarSolicitud(method, url, body, headers, auth, user, password string) (int, string, error) {
	// Create HTTP client with timeout
	client := &http.Client{Timeout: 15 * time.Second}

	// Create request with body (if any)
	var reqBody io.Reader
	if body != "" {
		reqBody = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return 0, "", err
	}

	// Add headers from JSON
	if headers != "" {
		var headerMap map[string]string
		if err := json.Unmarshal([]byte(headers), &headerMap); err != nil {
			return 0, "", fmt.Errorf("error parsing headers JSON: %v", err)
		}
		for key, value := range headerMap {
			req.Header.Add(key, value)
		}
	}

	// Handle authentication
	if auth == "Bearer" && user != "" {
		req.Header.Add("Authorization", "Bearer "+user)
	} else if auth == "Basic" && user != "" && password != "" {
		creds := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))
		req.Header.Add("Authorization", "Basic "+creds)
	}

	// Execute HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("error in HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the entire response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error reading response body: %v", err)
	}

	return resp.StatusCode, string(bodyBytes), nil
}
