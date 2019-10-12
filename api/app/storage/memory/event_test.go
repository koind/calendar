package memory

import (
	"context"
	"github.com/koind/calendar/api/app/domain/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const evntID = 1

var eventRepository repository.EventRepositoryInterface

func init() {
	eventRepository = NewEventRepository()
}

func before() {
	event := repository.Event{
		ID:             evntID,
		Title:          "Купить часы",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	eventRepository.Create(context.Background(), event)
}

func after() {
	eventRepository.Delete(context.Background(), evntID)
}

func TestEventRepository_Create(t *testing.T) {
	event := repository.Event{
		ID:             evntID,
		Title:          "Купить часы",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	_, err := eventRepository.Create(context.Background(), event)

	assert.Nil(t, err, "не должно быть ошибки при создании")

	after()
}

func TestEventRepository_Update(t *testing.T) {
	before()

	event := repository.Event{
		ID:             evntID,
		Title:          "Купить Rolex",
		Datetime:       time.Now(),
		Duration:       time.Second * 5,
		Description:    "Купить Rolex за $5000",
		UserId:         123,
		TimeSendNotify: time.Now(),
	}

	_, err := eventRepository.Update(context.Background(), evntID, event)
	assert.Nil(t, err, "не должно быть ошибки при обновлении")

	_, err = eventRepository.Update(context.Background(), 22, event)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, repository.EventNotFountError.Error(), "ошибки должны совподать")
	}

	after()
}

func TestEventRepository_Delete(t *testing.T) {
	before()

	err := eventRepository.Delete(context.Background(), evntID)
	assert.Nil(t, err, "не должно быть ошибки при удалении")

	_, err = eventRepository.FindByID(context.Background(), evntID)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, repository.EventNotFountError.Error(), "ошибки должны совподать")
	}
}

func TestEventRepository_FindByID(t *testing.T) {
	before()

	_, err := eventRepository.FindByID(context.Background(), evntID)
	assert.Nil(t, err, "не должно быть ошибки при удалении")

	after()

	_, err = eventRepository.FindByID(context.Background(), evntID)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, repository.EventNotFountError.Error(), "ошибки должны совподать")
	}
}

func TestEventRepository_FindAll(t *testing.T) {
	before()

	events, err := eventRepository.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, events, 1)

	after()

	events, _ = eventRepository.FindAll(context.Background())
	assert.Len(t, events, 0)
}
