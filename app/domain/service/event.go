package service

import (
	"github.com/koind/calendar/app/domain/model"
	"github.com/koind/calendar/app/domain/repository"
)

// Интерфейс сервиса событий
type EventServiceInterface interface {
	// Ищет событие оп UUID
	FindByUUID(UUID int) (*model.Event, error)

	// Создает новое событие
	Create(event model.Event) (*model.Event, error)

	// Обновляет событие
	Update(UUID int, event model.Event) (*model.Event, error)

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
func (service *EventService) FindByUUID(UUID int) (*model.Event, error) {
	event, err := service.eventRepository.FindByUUID(UUID)
	if err != nil {
		return nil, nil
	}

	return event, nil
}

// Создает новое событие
func (service *EventService) Create(event model.Event) (*model.Event, error) {
	newEvent, err := service.eventRepository.Create(event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Обновляет событие
func (service *EventService) Update(UUID int, event model.Event) (*model.Event, error) {
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
