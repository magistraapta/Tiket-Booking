package model

import (
	"time"
)

type EventModel struct {
	ID          string    `json:"id" validate:"omitempty,uuid4"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Organizer   string    `json:"organizer" validate:"required,min=2,max=50"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Location    string    `json:"location" validate:"required,min=5,max=100"`
	CreatedAt   time.Time `json:"created_at" validate:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at" validate:"omitempty"`
}

// GetStartDate returns the start date for validation
func (e *EventModel) GetStartDate() time.Time {
	return e.StartDate
}

// GetEndDate returns the end date for validation
func (e *EventModel) GetEndDate() time.Time {
	return e.EndDate
}
