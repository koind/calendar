package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/koind/calendar/api/app/domain/repository"
	"github.com/koind/calendar/api/app/domain/service"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Http сервис событий
type EventService struct {
	service.EventServiceInterface
	logger *zap.Logger
}

// Конструктор Http сервиса событий
func NewEventService(event service.EventServiceInterface, logger *zap.Logger) *EventService {
	return &EventService{
		EventServiceInterface: event,
		logger:                logger,
	}
}

// Обработчик создания события
func (service *EventService) CreateHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	event := repository.Event{}

	err := decoder.Decode(&event)
	if err != nil {
		w.WriteHeader(400)
	}

	newEvent, err := service.Create(r.Context(), event)
	if err != nil {
		service.logger.Error(
			"Во время создания события возникла ошибка",
			zap.Error(err),
		)

		w.Write([]byte(err.Error()))
	} else {
		service.logger.Info(
			"Событие создано",
			zap.Any("event", newEvent),
		)

		json.NewEncoder(w).Encode(newEvent)
	}
}

// Обработчик обновления данных события
func (service *EventService) UpdateHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		uuid, err := strconv.Atoi(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		decoder := json.NewDecoder(r.Body)
		event := repository.Event{}

		err = decoder.Decode(&event)
		if err != nil {
			w.WriteHeader(400)
		}

		newEvent, err := service.Update(r.Context(), uuid, event)
		if err != nil {
			service.logger.Error(
				"Во время обновления данных события возникла ошибка",
				zap.Error(err),
			)

			w.Write([]byte(err.Error()))
		} else {
			service.logger.Info(
				"Данные события обновлены",
				zap.Int("ID", uuid),
				zap.Any("event", newEvent),
			)

			json.NewEncoder(w).Encode(newEvent)
		}
	} else {
		w.WriteHeader(400)
	}
}

// Предоставляет список всех событий
func (service *EventService) FindAllHandle(w http.ResponseWriter, r *http.Request) {
	events, err := service.FindAll(r.Context())
	if err != nil {
		service.logger.Error(
			"Во время получения списка всех событий возникла ошибка",
			zap.Error(err),
		)

		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

// Обработчик удаления события
func (service *EventService) DeleteHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		uuid, err := strconv.Atoi(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		err = service.Delete(r.Context(), uuid)
		if err != nil {
			service.logger.Error(
				"Во время удаления события возникла ошибка",
				zap.Error(err),
			)

			w.Write([]byte(err.Error()))
		} else {
			service.logger.Info(
				"Событие было удалено",
				zap.Int("ID", uuid),
			)

			w.Write([]byte("ok"))
		}
	} else {
		w.WriteHeader(400)
	}
}
