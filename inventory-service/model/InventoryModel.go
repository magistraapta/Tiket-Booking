package model

import "time"

type InventoryModel struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type CreateInventoryRequest struct {
	Quantity int `json:"quantity"`
}
