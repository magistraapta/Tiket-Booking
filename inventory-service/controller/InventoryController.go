package controller

import (
	"encoding/json"
	"inventory-service/model"
	"inventory-service/service"
	"net/http"

	"github.com/gorilla/mux"
)

type InventoryController struct {
	inventoryService service.InventoryService
}

func NewInventoryController(inventoryService service.InventoryService) *InventoryController {
	return &InventoryController{
		inventoryService: inventoryService,
	}
}

func (c *InventoryController) CreateInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var request struct {
		EventID  string `json:"event_id"`
		Quantity int    `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := model.Response{
			Success: false,
			Message: "Invalid JSON format",
			Error:   "Request body must be valid JSON",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.EventID == "" {
		response := model.Response{
			Success: false,
			Message: "Event ID is required",
			Error:   "event_id field is mandatory",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Quantity <= 0 {
		response := model.Response{
			Success: false,
			Message: "Invalid quantity",
			Error:   "Quantity must be greater than 0",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	inventory, err := c.inventoryService.CreateInventory(request.EventID, request.Quantity)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Failed to create inventory",
			Error:   err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Inventory created successfully",
		Data:    inventory,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *InventoryController) GetInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	eventID := vars["event_id"]

	if eventID == "" {
		response := model.Response{
			Success: false,
			Message: "Event ID is required",
			Error:   "Event ID must be provided in the URL path",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	inventory, err := c.inventoryService.GetInventory(eventID)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Failed to get inventory",
			Error:   err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Inventory retrieved successfully",
		Data:    inventory,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *InventoryController) DecreaseInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	eventID := vars["event_id"]

	if eventID == "" {
		response := model.Response{
			Success: false,
			Message: "Event ID is required",
			Error:   "Event ID must be provided in the URL path",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Parse request body
	var request struct {
		Quantity int `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := model.Response{
			Success: false,
			Message: "Invalid JSON format",
			Error:   "Request body must be valid JSON",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Quantity <= 0 {
		response := model.Response{
			Success: false,
			Message: "Invalid quantity",
			Error:   "Quantity must be greater than 0",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	inventory, err := c.inventoryService.DecreaseInventory(eventID, request.Quantity)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Failed to decrease inventory",
			Error:   err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Inventory decreased successfully",
		Data:    inventory,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
