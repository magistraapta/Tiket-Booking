package validator

import (
	"booking-service/model"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResponse represents the response for validation errors
type ValidationResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

// ValidateCreateBookingRequest validates the create booking request
func ValidateCreateBookingRequest(req *model.CreateBookingRequest) []ValidationError {
	var errors []ValidationError

	// Validate EventID
	if strings.TrimSpace(req.EventID) == "" {
		errors = append(errors, ValidationError{
			Field:   "event_id",
			Message: "Event ID is required",
		})
	}

	// Validate UserID
	if strings.TrimSpace(req.UserID) == "" {
		errors = append(errors, ValidationError{
			Field:   "user_id",
			Message: "User ID is required",
		})
	}

	return errors
}

// CreateValidationResponse creates a validation response from errors
func CreateValidationResponse(errors []ValidationError) ValidationResponse {
	return ValidationResponse{
		Success: false,
		Message: "Validation failed",
		Errors:  errors,
	}
}

// CreateSuccessResponse creates a success response
func CreateSuccessResponse(message string) ValidationResponse {
	return ValidationResponse{
		Success: true,
		Message: message,
	}
}

// CreateErrorResponse creates an error response
func CreateErrorResponse(message string, err error) ValidationResponse {
	response := ValidationResponse{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Errors = []ValidationError{
			{
				Field:   "service",
				Message: err.Error(),
			},
		}
	}

	return response
}
