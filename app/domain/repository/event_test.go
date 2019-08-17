package repository

import (
	"github.com/koind/calendar/app/domain/model"
	"testing"
	"time"
)

var dummyEventRepository EventRepositoryInterface

func init() {
	dummyEventRepository = &DummyEventRepository{
		DB: make(map[int]model.Event),
	}
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

	err := dummyEventRepository.Create(event)
	if err != nil {
		t.Error("Не должно быть ошибки при создании")
	}
}

func TestDummyEventRepository_Update(t *testing.T) {
	event := model.Event{
		UUID:           1,
		Title:          "Купить Rolex",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex за $5000",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	err := dummyEventRepository.Update(1, event)
	if err != nil {
		t.Error("Не должно быть ошибки при обновлении")
	}

	err = dummyEventRepository.Update(22, event)
	if err != EventNotFountError {
		t.Errorf("Ошибки должны совподать: %v - %v", err, EventNotFountError)
	}
}

func TestDummyEventRepository_Delete(t *testing.T) {
	err := dummyEventRepository.Delete(1)
	if err != nil {
		t.Error("Не должно быть ошибки при удалении")
	}

	_, err = dummyEventRepository.FindByUUID(1)
	if err != EventNotFountError {
		t.Errorf("Ошибки должны совподать: %v - %v", err, EventNotFountError)
	}
}
