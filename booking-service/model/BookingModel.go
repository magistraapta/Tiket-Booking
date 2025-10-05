package model

import "time"

type BookingModel struct {
	ID        string        `json:"id"`
	EventID   string        `json:"event_id"`
	UserID    string        `json:"user_id"`
	Status    BookingStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

// CreateBookingRequest represents the request body for creating a booking
type CreateBookingRequest struct {
	EventID string `json:"event_id" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
}

// BookingResponse represents the response for booking operations
type BookingResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *BookingModel `json:"data,omitempty"`
	Error   string        `json:"error,omitempty"`
}

// BookingWithEventResponse represents the response for booking with event details
type BookingWithEventResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *BookingData `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
}

// BookingData contains booking and event information
type BookingData struct {
	Booking *BookingModel `json:"booking"`
	Event   *EventModel   `json:"event"`
}

// EventModel represents an event from the event service
type EventModel struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Organizer   string    `json:"organizer"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
