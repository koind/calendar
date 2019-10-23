package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/koind/calendar/api/app/domain/repository"
	"github.com/koind/calendar/api/app/domain/service"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

var (
	RPSCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "api",
		Name:      "rps",
		Help:      "Requests per second",
	})
	responseStatus = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "api12",
		Name:      "response_status",
		Help:      "Response status of endpoints.",
	}, []string{"method"})
	responseTimeSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "api",
		Name:      "response_time",
		Help:      "Response time of endpoints.",
	}, []string{"method"})
)

func init() {
	prometheus.MustRegister(RPSCounter)
	prometheus.MustRegister(responseTimeSummary)
	prometheus.MustRegister(responseStatus)
}

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
	start := time.Now()
	decoder := json.NewDecoder(r.Body)
	event := repository.Event{}

	err := decoder.Decode(&event)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		responseStatus.WithLabelValues("CreateHandle-400").Inc()

		return
	}

	newEvent, err := service.Create(r.Context(), event)
	if err != nil {
		service.logger.Error(
			"Во время создания события возникла ошибка",
			zap.Error(err),
		)

		w.Write([]byte(err.Error()))
		responseStatus.WithLabelValues("CreateHandle-500").Inc()
	} else {
		service.logger.Info(
			"Событие создано",
			zap.Any("event", newEvent),
		)

		json.NewEncoder(w).Encode(newEvent)
		responseStatus.WithLabelValues("CreateHandle-200").Inc()
	}

	elapsed := float64(time.Since(start)) / float64(time.Microsecond)
	responseTimeSummary.WithLabelValues("CreateHandle").Observe(elapsed)
	RPSCounter.Inc()
}

// Обработчик обновления данных события
func (service *EventService) UpdateHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		uuid, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			responseStatus.WithLabelValues("UpdateHandle-400").Inc()

			return
		}

		decoder := json.NewDecoder(r.Body)
		event := repository.Event{}

		err = decoder.Decode(&event)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			responseStatus.WithLabelValues("UpdateHandle-400").Inc()

			return
		}

		newEvent, err := service.Update(r.Context(), uuid, event)
		if err != nil {
			service.logger.Error(
				"Во время обновления данных события возникла ошибка",
				zap.Error(err),
			)

			w.Write([]byte(err.Error()))
			responseStatus.WithLabelValues("UpdateHandle-500").Inc()
		} else {
			service.logger.Info(
				"Данные события обновлены",
				zap.Int("ID", uuid),
				zap.Any("event", newEvent),
			)

			json.NewEncoder(w).Encode(newEvent)
			responseStatus.WithLabelValues("UpdateHandle-200").Inc()
		}
	}

	elapsed := float64(time.Since(start)) / float64(time.Microsecond)
	responseTimeSummary.WithLabelValues("UpdateHandle").Observe(elapsed)
	RPSCounter.Inc()
}

// Предоставляет список всех событий
func (service *EventService) FindAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	events, err := service.FindAll(r.Context())
	if err != nil {
		service.logger.Error(
			"Во время получения списка всех событий возникла ошибка",
			zap.Error(err),
		)

		w.Write([]byte(err.Error()))
		responseStatus.WithLabelValues("FindAllHandle-500").Inc()
	} else {
		json.NewEncoder(w).Encode(events)
		responseStatus.WithLabelValues("FindAllHandle-200").Inc()
	}

	elapsed := float64(time.Since(start)) / float64(time.Microsecond)
	responseTimeSummary.WithLabelValues("FindAllHandle").Observe(elapsed)
	RPSCounter.Inc()
}

// Обработчик удаления события
func (service *EventService) DeleteHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		uuid, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			responseStatus.WithLabelValues("DeleteHandle-400").Inc()

			return
		}

		err = service.Delete(r.Context(), uuid)
		if err != nil {
			service.logger.Error(
				"Во время удаления события возникла ошибка",
				zap.Error(err),
			)

			w.Write([]byte(err.Error()))
			responseStatus.WithLabelValues("DeleteHandle-500").Inc()
		} else {
			service.logger.Info(
				"Событие было удалено",
				zap.Int("ID", uuid),
			)

			w.Write([]byte("ok"))
			responseStatus.WithLabelValues("DeleteHandle-200").Inc()
		}
	}

	elapsed := float64(time.Since(start)) / float64(time.Microsecond)
	responseTimeSummary.WithLabelValues("DeleteHandle").Observe(elapsed)
	RPSCounter.Inc()
}
