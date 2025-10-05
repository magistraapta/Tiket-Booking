package validator

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidationError represents a validation error with field and message
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResponse represents the response structure for validation errors
type ValidationResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

// EventValidator interface for event validation
type EventValidator interface {
	GetStartDate() time.Time
	GetEndDate() time.Time
}

// ValidateEvent validates an event model and returns validation errors
func ValidateEvent(event interface{}) []ValidationError {
	var validationErrors []ValidationError

	// Basic field validation
	err := validate.Struct(event)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field:   err.Field(),
				Message: getErrorMessage(err),
			})
		}
	}

	// Custom business logic validation
	if customErrors := validateEventBusinessRules(event); len(customErrors) > 0 {
		validationErrors = append(validationErrors, customErrors...)
	}

	return validationErrors
}

// validateEventBusinessRules performs custom business logic validation
func validateEventBusinessRules(event interface{}) []ValidationError {
	var errors []ValidationError

	// Use reflection to access StartDate and EndDate fields
	if validator, ok := event.(EventValidator); ok {
		startDate := validator.GetStartDate()
		endDate := validator.GetEndDate()

		// Check if end date is after start date
		if !endDate.IsZero() && !startDate.IsZero() && endDate.Before(startDate) {
			errors = append(errors, ValidationError{
				Field:   "end_date",
				Message: "End date must be after start date",
			})
		}

		// Check if start date is in the future
		if !startDate.IsZero() && startDate.Before(time.Now()) {
			errors = append(errors, ValidationError{
				Field:   "start_date",
				Message: "Start date must be in the future",
			})
		}
	}

	return errors
}

// getErrorMessage returns a user-friendly error message for validation errors
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
	case "uuid4":
		return fmt.Sprintf("%s must be a valid UUID", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}

// CreateValidationResponse creates a standardized validation response
func CreateValidationResponse(errors []ValidationError) ValidationResponse {
	if len(errors) == 0 {
		return ValidationResponse{
			Success: true,
			Message: "Validation passed",
		}
	}

	return ValidationResponse{
		Success: false,
		Message: "Validation failed",
		Errors:  errors,
	}
}
