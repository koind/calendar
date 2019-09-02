package service

import (
	"github.com/koind/calendar/app/domain/repository"
)

// Интерфейс сервиса событий
type EventServiceInterface interface {
	// Ищет событие оп UUID
	FindByUUID(UUID int) (*repository.Event, error)

	// Создает новое событие
	Create(event repository.Event) (*repository.Event, error)

	// Обновляет событие
	Update(UUID int, event repository.Event) (*repository.Event, error)

	// Удаляет событие
	Delete(UUID int) error
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

// Ищет событие оп UUID
func (service *EventService) FindByUUID(UUID int) (*repository.Event, error) {
	event, err := service.eventRepository.FindByUUID(UUID)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// Создает новое событие
func (service *EventService) Create(event repository.Event) (*repository.Event, error) {
	newEvent, err := service.eventRepository.Create(event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Обновляет событие
func (service *EventService) Update(UUID int, event repository.Event) (*repository.Event, error) {
	newEvent, err := service.eventRepository.Update(UUID, event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Удаляет событие
func (service *EventService) Delete(UUID int) error {
	err := service.eventRepository.Delete(UUID)
	if err != nil {
		return err
	}

	return nil
}
