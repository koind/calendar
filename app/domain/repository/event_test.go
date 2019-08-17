package repository

import (
	"github.com/koind/calendar/app/domain/model"
	"testing"
	"time"
)

const evntUUID = 1

var dummyEventRepository EventRepositoryInterface

func init() {
	dummyEventRepository = NewDummyEventRepository()
}

func before() {
	event := model.Event{
		UUID:           1,
		Title:          "Купить часы",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	dummyEventRepository.Create(event)
}

func after() {
	dummyEventRepository.Delete(evntUUID)
}

func TestDummyEventRepository_Create(t *testing.T) {
	event := model.Event{
		UUID:           1,
		Title:          "Купить часы",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	_, err := dummyEventRepository.Create(event)
	if err != nil {
		t.Error("Не должно быть ошибки при создании")
	}

	after()
}

func TestDummyEventRepository_Update(t *testing.T) {
	before()

	event := model.Event{
		UUID:           1,
		Title:          "Купить Rolex",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex за $5000",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	_, err := dummyEventRepository.Update(evntUUID, event)
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	_, err = dummyEventRepository.Update(22, event)
	if err != EventNotFountError {
		t.Errorf("Ошибки должны совподать: %v - %v", err, EventNotFountError)
	}

	after()
}

func TestDummyEventRepository_Delete(t *testing.T) {
	before()

	err := dummyEventRepository.Delete(1)
	if err != nil {
		t.Error("Не должно быть ошибки при удалении")
	}

	_, err = dummyEventRepository.FindByUUID(1)
	if err != EventNotFountError {
		t.Errorf("Ошибки должны совподать: %v - %v", err, EventNotFountError)
	}
}

func TestDummyEventRepository_FindOnDay(t *testing.T) {
	before()

	eventList, err := dummyEventRepository.FindOnDay(time.Now())
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	if len(eventList) != 1 {
		t.Errorf("Количество событий должно совпадать: %v - %v", len(eventList), 1)
	}

	after()
}

func TestDummyEventRepository_FindOnWeek(t *testing.T) {
	before()

	eventList, err := dummyEventRepository.FindOnWeek(time.Now().Weekday())
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	if len(eventList) != 1 {
		t.Errorf("Количество событий должно совпадать: %v - %v", len(eventList), 1)
	}

	after()
}

func TestDummyEventRepository_FindOnMonth(t *testing.T) {
	before()

	eventList, err := dummyEventRepository.FindOnMonth(time.Now().Month())
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	if len(eventList) != 1 {
		t.Errorf("Количество событий должно совпадать: %v - %v", len(eventList), 1)
	}

	after()
}
