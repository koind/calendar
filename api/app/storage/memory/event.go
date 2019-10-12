package memory

import (
	"context"
	"github.com/koind/calendar/api/app/domain/repository"
	"sync"
)

func NewEventRepository() *EventRepository {
	return &EventRepository{
		DB: make(map[int]repository.Event),
	}
}

// Фиктивный репозиторий событий
type EventRepository struct {
	sync.RWMutex
	DB map[int]repository.Event
}

// Ищет событие оп ID
func (r *EventRepository) FindByID(ctx context.Context, ID int) (*repository.Event, error) {
	r.RLock()
	defer r.RUnlock()

	event, has := r.DB[ID]
	if !has {
		return nil, repository.EventNotFountError
	}

	return &event, nil
}

// Возвращает список всех событий
func (r *EventRepository) FindAll(ctx context.Context) ([]*repository.Event, error) {
	r.RLock()
	defer r.RUnlock()

	if len(r.DB) <= 0 {
		return nil, nil
	}

	events := make([]*repository.Event, 0, len(r.DB))

	for _, event := range r.DB {
		events = append(events, &event)
	}

	return events, nil
}

// Создает новое событие
func (r *EventRepository) Create(ctx context.Context, event repository.Event) (*repository.Event, error) {
	r.Lock()
	defer r.Unlock()

	r.DB[event.ID] = event

	return &event, nil
}

// Обновляет событие
func (r *EventRepository) Update(ctx context.Context, ID int, event repository.Event) (*repository.Event, error) {
	_, err := r.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	r.Lock()
	defer r.Unlock()

	r.DB[ID] = event

	return &event, nil
}

// Удаляет событие
func (r *EventRepository) Delete(ctx context.Context, ID int) error {
	_, err := r.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	r.Lock()
	defer r.Unlock()

	delete(r.DB, ID)

	return nil
}
