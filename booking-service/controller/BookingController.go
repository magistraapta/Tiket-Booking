package controller

import (
	"booking-service/model"
	"booking-service/service"
	"booking-service/validator"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type BookingController struct {
	bookingService *service.BookingService
}

func NewBookingController(bookingService *service.BookingService) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

func (c *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := &model.CreateBookingRequest{}

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response := validator.CreateErrorResponse("Invalid JSON format", errors.New("Request body must be valid JSON"))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate request
	validationErrors := validator.ValidateCreateBookingRequest(req)
	if len(validationErrors) > 0 {
		response := validator.CreateValidationResponse(validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Convert request to booking model
	createdBooking, err := c.bookingService.CreateBooking(req)
	if err != nil {
		response := validator.CreateErrorResponse("Failed to create booking", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.BookingResponse{
		Success: true,
		Message: "Booking created successfully",
		Data:    createdBooking,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *BookingController) GetBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookingID := vars["id"]

	if bookingID == "" {
		response := validator.CreateErrorResponse("Booking ID is required", errors.New("Booking ID must be provided in the URL path"))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	booking, err := c.bookingService.GetBooking(bookingID)
	if err != nil {
		response := validator.CreateErrorResponse("Booking not found", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.BookingResponse{
		Success: true,
		Message: "Booking retrieved successfully",
		Data:    booking,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *BookingController) GetBookingWithEventDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookingID := vars["id"]

	if bookingID == "" {
		response := validator.CreateErrorResponse("Booking ID is required", errors.New("Booking ID must be provided in the URL path"))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	booking, event, err := c.bookingService.GetBookingWithEventDetails(bookingID)
	if err != nil {
		response := validator.CreateErrorResponse("Failed to retrieve booking with event details", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.BookingWithEventResponse{
		Success: true,
		Message: "Booking with event details retrieved successfully",
		Data: &model.BookingData{
			Booking: booking,
			Event:   event,
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
