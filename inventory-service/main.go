package main

import (
	"inventory-service/client"
	"inventory-service/controller"
	"inventory-service/service"
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

	// Initialize inventory service with event client
	inventoryService := service.NewInventoryService(eventClient)

	// Initialize inventory controller
	inventoryController := controller.NewInventoryController(inventoryService)

	// Define routes
	router.HandleFunc("/inventory", inventoryController.CreateInventory).Methods("POST")
	router.HandleFunc("/inventory/{event_id}", inventoryController.GetInventory).Methods("GET")
	router.HandleFunc("/inventory/{event_id}/decrease", inventoryController.DecreaseInventory).Methods("POST")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "inventory-service"}`))
	}).Methods("GET")

	log.Printf("Inventory service is running on port 8082")
	log.Printf("Event service URL: %s", eventServiceURL)
	log.Fatal(http.ListenAndServe(":8082", router))
}
