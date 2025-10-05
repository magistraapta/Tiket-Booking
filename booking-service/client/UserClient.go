package client

import (
	"booking-service/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserClient interface {
	GetUsersById(id string) (*model.UserResponse, error)
}

type HTTPUserClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPUserClient(baseURL string) *HTTPUserClient {
	return &HTTPUserClient{baseURL: baseURL, httpClient: &http.Client{Timeout: 30 * time.Second}}
}

func (c *HTTPUserClient) GetUsersById(id string) (*model.UserResponse, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, id)

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
		return nil, fmt.Errorf("user service returned status %d", resp.StatusCode)
	}

	var response model.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
