package service

import (
	"fmt"
	"inventory-service/client"
	"inventory-service/model"
	"time"

	"github.com/google/uuid"
)

type InventoryService interface {
	GetInventory(eventID string) (*model.InventoryModel, error)
	CreateInventory(eventID string, quantity int) (*model.InventoryModel, error)
	DecreaseInventory(eventID string, quantity int) (*model.InventoryModel, error)
}

type InventoryServiceImpl struct {
	eventClient client.EventClient
	inventories map[string]*model.InventoryModel // In-memory storage
}

func NewInventoryService(eventClient client.EventClient) *InventoryServiceImpl {
	return &InventoryServiceImpl{
		eventClient: eventClient,
		inventories: make(map[string]*model.InventoryModel),
	}
}

func (s *InventoryServiceImpl) CreateInventory(eventID string, quantity int) (*model.InventoryModel, error) {
	// Don't validate event existence during creation - the event service calls this
	// during event creation, so the event might not be fully created yet

	inventory := &model.InventoryModel{
		ID:        uuid.New().String(), // Generate unique inventory ID
		EventID:   eventID,
		Quantity:  quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Store the inventory
	s.inventories[eventID] = inventory

	return inventory, nil
}

func (s *InventoryServiceImpl) GetInventory(eventID string) (*model.InventoryModel, error) {
	// Check if inventory exists in storage
	if inventory, exists := s.inventories[eventID]; exists {
		return inventory, nil
	}

	// If not found, validate event exists and return default inventory
	event, err := s.eventClient.GetEventById(eventID)
	if err != nil {
		return nil, err
	}

	// Create default inventory for existing event
	defaultInventory := &model.InventoryModel{
		ID:        uuid.New().String(),
		EventID:   event.ID,
		Quantity:  100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Store the default inventory
	s.inventories[eventID] = defaultInventory

	return defaultInventory, nil
}

func (s *InventoryServiceImpl) DecreaseInventory(eventID string, requestedQuantity int) (*model.InventoryModel, error) {
	// Get the stored inventory
	inventory, err := s.GetInventory(eventID)
	if err != nil {
		return nil, err
	}

	// Check if enough inventory is available
	if inventory.Quantity < requestedQuantity {
		return nil, fmt.Errorf("insufficient inventory: requested %d, available %d", requestedQuantity, inventory.Quantity)
	}

	// Decrease the quantity
	inventory.Quantity = inventory.Quantity - requestedQuantity
	inventory.UpdatedAt = time.Now()

	// Update the stored inventory
	s.inventories[eventID] = inventory

	return inventory, nil
}
