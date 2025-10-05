package client

import (
	"booking-service/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type EventClient interface {
	GetEvent(eventID string) (*model.EventModel, error)
	GetEvents() []*model.EventModel
}

type HTTPEventClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPEventClient(baseURL string) *HTTPEventClient {
	return &HTTPEventClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *HTTPEventClient) GetEvent(eventID string) (*model.EventModel, error) {
	url := fmt.Sprintf("%s/events/%s", c.baseURL, eventID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("event service returned status %d", resp.StatusCode)
	}

	var response struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Data    *model.EventModel `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("event service error: %s", response.Message)
	}

	return response.Data, nil
}

func (c *HTTPEventClient) GetEvents() []*model.EventModel {
	url := fmt.Sprintf("%s/events", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Event service returned status %d", resp.StatusCode)
		return nil
	}

	var events []*model.EventModel
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return nil
	}

	return events
}
