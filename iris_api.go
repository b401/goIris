package goiris

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ApiMessage represents the structure of the api response meta fields
type ApiMeta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// VersionResponse represents the structure of the response from the /api/versions endpoint.
type VersionResponse struct {
	Data struct {
		ApiCurrent  string `json:"api_current"`
		ApiMin      string `json:"api_min"`
		IrisCurrent string `json:"iris_current"`
	} `json:"data"`
	ApiMeta
}

// PingResponse represents the structure of the response from the /api/ping endpoint.
type PingResponse struct {
	ApiMeta
}

// GetAPIVersion queries the API version from the /api/versions endpoint.
// It returns a VersionResponse containing the current API version, minimum API version,
// and the current IRIS version. If the request fails or the response cannot be decoded,
// an error is returned.
//
// Example usage:
//
//	client := api.NewAPIClient(httpClient, "https://api.example.com")
//	version, err := client.GetAPIVersion()
//	if err != nil {
//	    log.Fatalf("Failed to get API version: %v", err)
//	}
//	fmt.Println("API Version:", version.Data.ApiCurrent)
//
// Returns:
// - *VersionResponse: The response from the API containing version information.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) GetAPIVersion() (*VersionResponse, error) {
	builder := NewRequestBuilder().
		SetURL("/api/versions").
		SetMethod(http.MethodGet).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var versionResponse VersionResponse
	if err := json.NewDecoder(req.Body).Decode(&versionResponse); err != nil {
		return nil, err
	}

	return &versionResponse, nil
}

// Ping is used to test authentication against the /api/ping endpoint.
// It returns a PingResponse struct.
// If the request fails or the authentication was unsuccessfull an error is returned.
//
// Example usage:
//
//	    client := api.NewAPIClient(httpClient, "https://api.example.com")
//	    pong, err := client.Ping()
//	    if err != nil {
//	        log.Fatalf("Failed to authenticate: %v", err)
//	    }
//			 fmt.Printf("%s", pong.Message) // should print pong
//
// Returns:
// - *PingResponse*: A struct that contains a message and status field.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) Ping() (*PingResponse, error) {
	builder := NewRequestBuilder().
		SetURL("/api/ping").
		SetMethod(http.MethodGet).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var pingResponse PingResponse
	if err := json.NewDecoder(req.Body).Decode(&pingResponse); err != nil {
		return nil, err
	}

	return &pingResponse, nil

}
