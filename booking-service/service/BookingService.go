package service

import (
	"booking-service/client"
	"booking-service/model"
	"fmt"
	"time"
)

var bookingDummy = []*model.BookingModel{}

type BookingService struct {
	eventClient client.EventClient
	userClient  client.UserClient
}

func NewBookingService(eventClient client.EventClient, userClient client.UserClient) *BookingService {
	return &BookingService{
		eventClient: eventClient,
		userClient:  userClient,
	}
}

func (s *BookingService) CreateBooking(bookingRequest *model.CreateBookingRequest) (*model.BookingModel, error) {
	event, err := s.eventClient.GetEvent(bookingRequest.EventID)
	user, err := s.userClient.GetUsersById(bookingRequest.UserID)

	if err != nil {
		return nil, fmt.Errorf("failed to validate event: %w", err)
	}

	booking := &model.BookingModel{
		EventID:   event.ID,
		UserID:    user.Data.ID,
		Status:    model.BookingStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return booking, nil

}

func (s *BookingService) GetBooking(bookingID string) (*model.BookingModel, error) {
	// In a real application, this would fetch from database
	for _, booking := range bookingDummy {
		if booking.ID == bookingID {
			return booking, nil
		}
	}
	return nil, fmt.Errorf("booking not found")
}

func (s *BookingService) GetBookingWithEventDetails(bookingID string) (*model.BookingModel, *model.EventModel, error) {
	booking, err := s.GetBooking(bookingID)
	if err != nil {
		return nil, nil, err
	}

	event, err := s.eventClient.GetEvent(booking.EventID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get event details: %w", err)
	}

	return booking, event, nil
}
