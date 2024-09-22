package utils

import (
	"encoding/json" // Added json import
	"fmt"           // Added fmt import
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](fullUrl string, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, error) {
	client := http.Client{}
	u, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, err
	}

	// For GET requests, append query parameters
	if httpMethod == "GET" {
		q := u.Query()
		for k, v := range queryParameters {
			q.Set(k, strings.Join(v, ","))
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return responseType, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Make the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return responseType, err
	}
	defer res.Body.Close() // Defer after receiving response

	// Check if response is nil
	if res == nil {
		return responseType, fmt.Errorf("error: calling %s returned empty response", u.String())
	}

	// Read the response body
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return responseType, err
	}

	// Check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		return responseType, fmt.Errorf("error: calling %s: \n status: %s \n responseData: %s", u.String(), res.Status, responseData)
	}

	// Unmarshal the response body into the expected response type
	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseType, err
	}

	// Return the unmarshalled object
	return responseObject, nil
}
