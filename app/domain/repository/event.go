package repository

import (
	"errors"
	"time"
)

var EventNotFountError = errors.New("событие не найдено")

// Модель событий
type Event struct {
	ID             int           `json:"id"`                         //ID - уникальный идентификатор события
	Title          string        `json:"title"`                      //Заголовок - короткий текст
	Datetime       time.Time     `json:"datetime"`                   //Дата и время события
	Duration       time.Duration `json:"duration"`                   //Длительность события
	Description    string        `json:"description,omitempty"`      //Описание события, опционально
	UserId         int           `json:"user_id"`                    //Пользователь, владелец события
	TimeSendNotify time.Time     `json:"time_send_notify,omitempty"` //За сколько времени высылать уведомление, опционально
}

// Интерфейс репозитория событий
type EventRepositoryInterface interface {
	// Ищет событие оп ID
	FindByID(ID int) (*Event, error)

	// Создает новое событие
	Create(event Event) (*Event, error)

	// Обновляет событие
	Update(ID int, event Event) (*Event, error)

	// Удаляет событие
	Delete(ID int) error
}

// Интерфейс поиска событий
type EventFindRepositoryInterface interface {
	// Ищет события на день
	FindOnDay(day time.Time) ([]Event, error)

	// Ищет события на неделю
	FindOnWeek(week time.Weekday) ([]Event, error)

	// Ищет события на месяц
	FindOnMonth(month time.Month) ([]Event, error)
}
