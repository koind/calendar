package service

import (
	"context"
	"github.com/koind/calendar/api/app/domain/repository"
)

// Интерфейс сервиса событий
type EventServiceInterface interface {
	// Создает новое событие
	Create(ctx context.Context, event repository.Event) (*repository.Event, error)

	// Возвращает список всех событий
	FindAll(ctx context.Context) ([]*repository.Event, error)

	// Обновляет событие
	Update(ctx context.Context, ID int, event repository.Event) (*repository.Event, error)

	// Удаляет событие
	Delete(ctx context.Context, ID int) error
}

// Конструктор сервиса событий
func NewEventService(eventRepository repository.EventRepositoryInterface) *EventService {
	return &EventService{
		eventRepository: eventRepository,
	}
}

// Сервис событий
type EventService struct {
	eventRepository repository.EventRepositoryInterface
}

// Создает новое событие
func (service *EventService) Create(ctx context.Context, event repository.Event) (*repository.Event, error) {
	newEvent, err := service.eventRepository.Create(ctx, event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Возвращает список всех событий
func (service *EventService) FindAll(ctx context.Context) ([]*repository.Event, error) {
	events, err := service.eventRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}

// Обновляет событие
func (service *EventService) Update(ctx context.Context, ID int, event repository.Event) (*repository.Event, error) {
	newEvent, err := service.eventRepository.Update(ctx, ID, event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Удаляет событие
func (service *EventService) Delete(ctx context.Context, ID int) error {
	err := service.eventRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
