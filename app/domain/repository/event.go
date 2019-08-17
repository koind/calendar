package repository

import (
	"errors"
	"github.com/koind/calendar/app/domain/model"
	"sync"
)

// Интерфейс репозитория событий
type EventRepositoryInterface interface {
	// Ищет событие оп UUID
	FindByUUID(UUID int) (model.Event, error)

	// Создает новое событие
	Create(event model.Event) error

	// Обновляет событие
	Update(UUID int, event model.Event) error

	// Удаляет событие
	Delete(UUID int) error
}

var EventNotFountError = errors.New("событие не найдено")

// Фиктивный репозиторий событий
type DummyEventRepository struct {
	sync.RWMutex
	DB map[int]model.Event
}

// Ищет событие оп UUID
func (repository *DummyEventRepository) FindByUUID(UUID int) (model.Event, error) {
	repository.RLock()
	defer repository.RUnlock()

	event, has := repository.DB[UUID]
	if !has {
		return model.Event{}, EventNotFountError
	}

	return event, nil
}

// Создает новое событие
func (repository *DummyEventRepository) Create(event model.Event) error {
	repository.Lock()
	defer repository.Unlock()

	repository.DB[event.UUID] = event

	return nil
}

// Обновляет событие
func (repository *DummyEventRepository) Update(UUID int, event model.Event) error {
	_, err := repository.FindByUUID(UUID)
	if err != nil {
		return err
	}

	repository.Lock()
	defer repository.Unlock()

	repository.DB[UUID] = event

	return nil
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
