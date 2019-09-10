package service

import (
	"github.com/koind/calendar/app/domain/repository"
)

// Интерфейс сервиса событий
type EventServiceInterface interface {
	// Ищет событие оп ID
	FindByID(ID int) (*repository.Event, error)

	// Создает новое событие
	Create(event repository.Event) (*repository.Event, error)

	// Обновляет событие
	Update(ID int, event repository.Event) (*repository.Event, error)

	// Удаляет событие
	Delete(ID int) error
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

// Ищет событие оп ID
func (service *EventService) FindByID(ID int) (*repository.Event, error) {
	event, err := service.eventRepository.FindByID(ID)
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
func (service *EventService) Update(ID int, event repository.Event) (*repository.Event, error) {
	newEvent, err := service.eventRepository.Update(ID, event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Удаляет событие
func (service *EventService) Delete(ID int) error {
	err := service.eventRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
