package main

import (
	"event-service/client"
	"event-service/controller"
	"event-service/repository"
	"event-service/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	inventoryClient := client.NewInventoryClient(os.Getenv("INVENTORY_SERVICE_URL"))

	eventRepository := repository.NewEventRepository(nil, inventoryClient)
	eventService := service.NewEventService(eventRepository)
	eventController := controller.NewEventController(eventService)

	router.HandleFunc("/events", eventController.CreateEvent).Methods("POST")
	router.HandleFunc("/events", eventController.GetEvents).Methods("GET")
	router.HandleFunc("/events/{id}", eventController.GetEventById).Methods("GET")

	log.Println("Server is running on port 8084")
	http.ListenAndServe(":8084", router)
}
