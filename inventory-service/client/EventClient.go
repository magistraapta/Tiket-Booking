package client

import (
	"encoding/json"
	"fmt"
	"inventory-service/model"
	"net/http"
	"time"
)

type EventClient interface {
	GetEventById(eventID string) (*model.EventModel, error)
}

type HTTPEventClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPEventClient(baseURL string) *HTTPEventClient {
	return &HTTPEventClient{baseURL: baseURL, httpClient: &http.Client{Timeout: 30 * time.Second}}
}

func (c *HTTPEventClient) GetEventById(eventID string) (*model.EventModel, error) {
	/*
		Params: eventID (string) is the id of the event to get.
		Returns: EventModel (struct) is the event model.
		Error: error is the error if the request fails.
	**/

	// url
	url := fmt.Sprintf("%s/events/%s", c.baseURL, eventID)

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	// do request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	// close response body
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("event service returned status %d", resp.StatusCode)
	}

	// decode response
	var response struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Data    *model.EventModel `json:"data"`
	}

	// decode response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// check success
	if !response.Success {
		return nil, fmt.Errorf("event service error: %s", response.Message)
	}

	// return response
	return response.Data, nil
}
