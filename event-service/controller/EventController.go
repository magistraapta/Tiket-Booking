package controller

import (
	"encoding/json"
	"event-service/model"
	"event-service/service"
	"event-service/validator"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type EventController struct {
	eventService *service.EventService
}

func NewEventController(eventService *service.EventService) *EventController {
	return &EventController{eventService: eventService}
}

func (c *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	event := &model.EventModel{}

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		response := validator.ValidationResponse{
			Success: false,
			Message: "Invalid JSON format",
			Errors: []validator.ValidationError{
				{Field: "body", Message: "Request body must be valid JSON"},
			},
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Set timestamps
	now := time.Now()
	event.CreatedAt = now
	event.UpdatedAt = now

	validationErrors := validator.ValidateEvent(event)
	if len(validationErrors) > 0 {
		response := validator.CreateValidationResponse(validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err := c.eventService.CreateEvent(event)
	if err != nil {
		response := validator.ValidationResponse{
			Success: false,
			Message: "Failed to create event",
			Errors: []validator.ValidationError{
				{Field: "service", Message: err.Error()},
			},
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := validator.ValidationResponse{
		Success: true,
		Message: "Event created successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": response.Success,
		"message": response.Message,
		"data":    event,
	})
}

func (c *EventController) GetEvents(w http.ResponseWriter, r *http.Request) {
	events := c.eventService.GetEvents()
	json.NewEncoder(w).Encode(events)
}

func (c *EventController) GetEventById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		response := validator.ValidationResponse{
			Success: false,
			Message: "Event ID is required",
			Errors: []validator.ValidationError{
				{Field: "id", Message: "Event ID must be provided in the URL path"},
			},
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	event, err := c.eventService.GetEventById(id)
	if err != nil {
		response := validator.ValidationResponse{
			Success: false,
			Message: "Event not found",
			Errors: []validator.ValidationError{
				{Field: "id", Message: err.Error()},
			},
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := validator.ValidationResponse{
		Success: true,
		Message: "Event retrieved successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": response.Success,
		"message": response.Message,
		"data":    event,
	})
}
