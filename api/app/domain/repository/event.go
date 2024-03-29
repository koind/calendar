package repository

import (
	"context"
	"errors"
	"time"
)

var EventNotFountError = errors.New("событие не найдено")

// Модель событий
type Event struct {
	ID             int           `json:"id" db:"id"`                                       //ID - уникальный идентификатор события
	Title          string        `json:"title" db:"title"`                                 //Заголовок - короткий текст
	Datetime       time.Time     `json:"datetime" db:"datetime"`                           //Дата и время события
	Duration       time.Duration `json:"duration" db:"duration"`                           //Длительность события
	Description    string        `json:"description,omitempty" db:"description"`           //Описание события, опционально
	UserId         int           `json:"user_id" db:"user_id"`                             //Пользователь, владелец события
	TimeSendNotify time.Time     `json:"time_send_notify,omitempty" db:"time_send_notify"` //За сколько времени высылать уведомление, опционально
}

// Интерфейс репозитория событий
type EventRepositoryInterface interface {
	// Ищет событие оп ID
	FindByID(ctx context.Context, ID int) (*Event, error)

	// Возвращает список всех событий
	FindAll(ctx context.Context) ([]*Event, error)

	// Создает новое событие
	Create(ctx context.Context, event Event) (*Event, error)

	// Обновляет событие
	Update(ctx context.Context, ID int, event Event) (*Event, error)

	// Удаляет событие
	Delete(ctx context.Context, ID int) error
}
