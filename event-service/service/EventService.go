package service

import (
	"event-service/model"
	"event-service/repository"
)

type EventService struct {
	eventRepository *repository.EventRepository
}

func NewEventService(eventRepository *repository.EventRepository) *EventService {
	return &EventService{eventRepository: eventRepository}
}

func (s *EventService) CreateEvent(event *model.EventModel) error {
	return s.eventRepository.CreateEvent(event)
}

func (s *EventService) GetEvents() []*model.EventModel {
	return s.eventRepository.GetEvents()
}

func (s *EventService) GetEventById(id string) (*model.EventModel, error) {
	return s.eventRepository.GetEventById(id)
}
