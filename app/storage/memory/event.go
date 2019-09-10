package memory

import (
	"github.com/koind/calendar/app/domain/repository"
	"sync"
	"time"
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
func (r *EventRepository) FindByID(ID int) (*repository.Event, error) {
	r.RLock()
	defer r.RUnlock()

	event, has := r.DB[ID]
	if !has {
		return nil, repository.EventNotFountError
	}

	return &event, nil
}

// Создает новое событие
func (r *EventRepository) Create(event repository.Event) (*repository.Event, error) {
	r.Lock()
	defer r.Unlock()

	r.DB[event.ID] = event

	return &event, nil
}

// Обновляет событие
func (r *EventRepository) Update(ID int, event repository.Event) (*repository.Event, error) {
	_, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	r.Lock()
	defer r.Unlock()

	r.DB[ID] = event

	return &event, nil
}

// Удаляет событие
func (r *EventRepository) Delete(ID int) error {
	_, err := r.FindByID(ID)
	if err != nil {
		return err
	}

	r.Lock()
	defer r.Unlock()

	delete(r.DB, ID)

	return nil
}

// Ищет события на день
func (r *EventRepository) FindOnDay(day time.Time) ([]repository.Event, error) {
	r.RLock()
	defer r.RUnlock()

	eventList := make([]repository.Event, 0, 0)
	for _, even := range r.DB {
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
func (r *EventRepository) FindOnWeek(week time.Weekday) ([]repository.Event, error) {
	r.RLock()
	defer r.RUnlock()

	eventList := make([]repository.Event, 0, 0)
	for _, even := range r.DB {
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
func (r *EventRepository) FindOnMonth(month time.Month) ([]repository.Event, error) {
	r.RLock()
	defer r.RUnlock()

	eventList := make([]repository.Event, 0, 0)
	for _, even := range r.DB {
		if even.Datetime.Month() == month {
			eventList = append(eventList, even)
		}
	}

	if len(eventList) > 0 {
		return eventList, nil
	}

	return nil, nil
}
