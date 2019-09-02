package memory

import (
	"github.com/koind/calendar/app/domain/repository"
	"testing"
	"time"
)

const evntUUID = 1

var eventRepository repository.EventRepositoryInterface

func init() {
	eventRepository = NewEventRepository()
}

func before() {
	event := repository.Event{
		UUID:           1,
		Title:          "Купить часы",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	eventRepository.Create(event)
}

func after() {
	eventRepository.Delete(evntUUID)
}

func TestEventRepository_Create(t *testing.T) {
	event := repository.Event{
		UUID:           1,
		Title:          "Купить часы",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	_, err := eventRepository.Create(event)
	if err != nil {
		t.Error("Не должно быть ошибки при создании")
	}

	after()
}

func TestEventRepository_Update(t *testing.T) {
	before()

	event := repository.Event{
		UUID:           1,
		Title:          "Купить Rolex",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex за $5000",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	_, err := eventRepository.Update(evntUUID, event)
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	_, err = eventRepository.Update(22, event)
	if err != repository.EventNotFountError {
		t.Errorf("Ошибки должны совподать: %v - %v", err, repository.EventNotFountError)
	}

	after()
}

func TestEventRepository_Delete(t *testing.T) {
	before()

	err := eventRepository.Delete(1)
	if err != nil {
		t.Error("Не должно быть ошибки при удалении")
	}

	_, err = eventRepository.FindByUUID(1)
	if err != repository.EventNotFountError {
		t.Errorf("Ошибки должны совподать: %v - %v", err, repository.EventNotFountError)
	}
}

func TestEventRepository_FindOnDay(t *testing.T) {
	before()

	eventList, err := eventRepository.FindOnDay(time.Now())
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	if len(eventList) != 1 {
		t.Errorf("Количество событий должно совпадать: %v - %v", len(eventList), 1)
	}

	after()
}

func TestEventRepository_FindOnWeek(t *testing.T) {
	before()

	eventList, err := eventRepository.FindOnWeek(time.Now().Weekday())
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	if len(eventList) != 1 {
		t.Errorf("Количество событий должно совпадать: %v - %v", len(eventList), 1)
	}

	after()
}

func TestEventRepository_FindOnMonth(t *testing.T) {
	before()

	eventList, err := eventRepository.FindOnMonth(time.Now().Month())
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	if len(eventList) != 1 {
		t.Errorf("Количество событий должно совпадать: %v - %v", len(eventList), 1)
	}

	after()
}
