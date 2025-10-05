package main

import (
	"booking-service/client"
	"booking-service/controller"
	"booking-service/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Get event service URL from environment variable or use default
	eventServiceURL := os.Getenv("EVENT_SERVICE_URL")
	if eventServiceURL == "" {
		eventServiceURL = "http://localhost:8080"
	}

	// Initialize HTTP client for event service
	eventClient := client.NewHTTPEventClient(eventServiceURL)

	// Initialize user service URL from environment variable or use default
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "http://localhost:8083"
	}

	// Initialize HTTP client for user service
	userClient := client.NewHTTPUserClient(userServiceURL)

	// Initialize booking service with event client
	bookingService := service.NewBookingService(eventClient, userClient)

	// Initialize booking controller
	bookingController := controller.NewBookingController(bookingService)

	// Define routes
	router.HandleFunc("/bookings", bookingController.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings/{id}", bookingController.GetBooking).Methods("GET")
	router.HandleFunc("/bookings/{id}/details", bookingController.GetBookingWithEventDetails).Methods("GET")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "booking-service"}`))
	}).Methods("GET")

	log.Printf("Booking service is running on port 8081")
	log.Printf("Event service URL: %s", eventServiceURL)
	log.Fatal(http.ListenAndServe(":8081", router))
}
