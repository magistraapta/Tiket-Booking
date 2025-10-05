package repository

import (
	"errors"
	"event-service/client"
	"event-service/model"
	"time"

	"log"

	"github.com/google/uuid"
)

var eventList = []*model.EventModel{
	{
		ID:          "1",
		Name:        "Event 1",
		Organizer:   "Organizer 1",
		Description: "Description 1",
		StartDate:   time.Date(2025, 12, 25, 10, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2025, 12, 25, 18, 0, 0, 0, time.UTC),
		Location:    "Location 1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          "2",
		Name:        "Event 2",
		Organizer:   "Organizer 2",
		Description: "Description 2",
		StartDate:   time.Date(2025, 12, 31, 10, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2025, 12, 31, 18, 0, 0, 0, time.UTC),
		Location:    "Location 2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

type EventRepository struct {
	inventoryClient *client.InventoryClient
}

func NewEventRepository(eventModel []*model.EventModel, inventoryClient *client.InventoryClient) *EventRepository {
	if eventModel != nil {
		eventList = eventModel
	}
	return &EventRepository{inventoryClient: inventoryClient}
}

func (r *EventRepository) CreateEvent(event *model.EventModel) error {
	// Generate event ID first
	event.ID = uuid.New().String()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	// Create inventory for the event
	_, err := r.inventoryClient.CreateInventory(event.ID, 100)
	if err != nil {
		log.Printf("Failed to create inventory for event %s: %v", event.ID, err)
		return err
	}

	// Add event to list
	eventList = append(eventList, event)
	log.Println("Event created: ", event.ID)
	return nil
}

func (r *EventRepository) GetEvents() []*model.EventModel {
	log.Println("Getting all events size:", len(eventList))
	if len(eventList) == 0 {
		log.Println("No events found")
		return nil
	}
	return eventList
}

func (r *EventRepository) GetEventById(id string) (*model.EventModel, error) {
	log.Println("Getting event by id:", id)
	for _, event := range eventList {
		if event.ID == id {
			log.Println("Event found:", event)
			return event, nil
		}
	}
	log.Println("Event not found")
	return nil, errors.New("event not found")
}
