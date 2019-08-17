package repository

import (
	"errors"
	"github.com/koind/calendar/app/domain/model"
	"sync"
	"time"
)

// Интерфейс репозитория событий
type EventRepositoryInterface interface {
	// Ищет событие оп UUID
	FindByUUID(UUID int) (*model.Event, error)

	// Создает новое событие
	Create(event model.Event) (*model.Event, error)

	// Обновляет событие
	Update(UUID int, event model.Event) (*model.Event, error)

	// Удаляет событие
	Delete(UUID int) error

	// Ищет события на день
	FindOnDay(day time.Time) ([]model.Event, error)

	// Ищет события на неделю
	FindOnWeek(week time.Weekday) ([]model.Event, error)

	// Ищет события на месяц
	FindOnMonth(month time.Month) ([]model.Event, error)
}

var EventNotFountError = errors.New("событие не найдено")

func NewDummyEventRepository() *DummyEventRepository {
	return &DummyEventRepository{
		DB: make(map[int]model.Event),
	}
}

// Фиктивный репозиторий событий
type DummyEventRepository struct {
	sync.RWMutex
	DB map[int]model.Event
}

// Ищет событие оп UUID
func (repository *DummyEventRepository) FindByUUID(UUID int) (*model.Event, error) {
	repository.RLock()
	defer repository.RUnlock()

	event, has := repository.DB[UUID]
	if !has {
		return nil, EventNotFountError
	}

	return &event, nil
}

// Создает новое событие
func (repository *DummyEventRepository) Create(event model.Event) (*model.Event, error) {
	repository.Lock()
	defer repository.Unlock()

	repository.DB[event.UUID] = event

	return &event, nil
}

// Обновляет событие
func (repository *DummyEventRepository) Update(UUID int, event model.Event) (*model.Event, error) {
	_, err := repository.FindByUUID(UUID)
	if err != nil {
		return nil, err
	}

	repository.Lock()
	defer repository.Unlock()

	repository.DB[UUID] = event

	return &event, nil
}

// Удаляет событие
func (repository *DummyEventRepository) Delete(UUID int) error {
	_, err := repository.FindByUUID(UUID)
	if err != nil {
		return err
	}

	repository.Lock()
	defer repository.Unlock()

	delete(repository.DB, UUID)

	return nil
}

// Ищет события на день
func (repository *DummyEventRepository) FindOnDay(day time.Time) ([]model.Event, error) {
	repository.RLock()
	defer repository.RUnlock()

	eventList := make([]model.Event, 0, 15)
	for _, even := range repository.DB {
		if even.Datetime.Day() == day.Day() {
			eventList = append(eventList, even)
		}
	}

	if len(eventList) > 0 {
		return eventList, nil
	}

	return nil, nil
}

// Ищет события на неделю
func (repository *DummyEventRepository) FindOnWeek(week time.Weekday) ([]model.Event, error) {
	repository.RLock()
	defer repository.RUnlock()

	eventList := make([]model.Event, 0, 15)
	for _, even := range repository.DB {
		if even.Datetime.Weekday() == week {
			eventList = append(eventList, even)
		}
	}

	if len(eventList) > 0 {
		return eventList, nil
	}

	return nil, nil
}

// Ищет события на месяц
func (repository *DummyEventRepository) FindOnMonth(month time.Month) ([]model.Event, error) {
	repository.RLock()
	defer repository.RUnlock()

	eventList := make([]model.Event, 0, 15)
	for _, even := range repository.DB {
		if even.Datetime.Month() == month {
			eventList = append(eventList, even)
		}
	}

	if len(eventList) > 0 {
		return eventList, nil
	}

	return nil, nil
}
