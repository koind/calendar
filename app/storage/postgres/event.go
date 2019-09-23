package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/koind/calendar/app/domain/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

const (
	querySelectEventByID = "SELECT * FROM events WHERE id=$1"
	queryInsertEvent     = `INSERT INTO events(title, datetime, duration, description, user_id, time_send_notify)
		VALUES (:title, :datetime, :duration, :description, :user_id, :time_send_notify)`
)

// Конструктор репозитория событий
func NewEventRepository(ctx context.Context, db *sqlx.DB, logger zap.Logger) *EventRepository {
	return &EventRepository{
		DB:     db,
		logger: logger,
		ctx:    ctx,
	}
}

// PostgreSQL репозиторий события
type EventRepository struct {
	DB     *sqlx.DB
	logger zap.Logger
	ctx    context.Context
}

// Ищет событие оп ID
func (r *EventRepository) FindByID(ID int) (*repository.Event, error) {
	ctx, cancel := context.WithTimeout(r.ctx, time.Duration(30)*time.Millisecond)
	defer cancel()

	if ctx.Err() == context.Canceled {
		r.logger.Info(
			"Поиск события по ID был прерван из-за отмены контекста",
			zap.Int("ID", ID),
		)

		return nil, errors.New("поиск события по ID был прерван из-за отмены контекста")
	}

	row := r.DB.QueryRowContext(ctx, querySelectEventByID, ID)

	event := repository.Event{}
	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Datetime,
		&event.Duration,
		&event.Description,
		&event.UserId,
		&event.TimeSendNotify,
	)

	if err != sql.ErrNoRows {
		r.logger.Warn(
			"Не удалось найти событие по ID",
			zap.Int("ID", ID),
		)

		return nil, errors.Wrap(err, "не удалось найти событие по ID")
	} else {
		r.logger.Warn(
			"Возникла ошибка при поиске события по ID",
			zap.Int("ID", ID),
		)

		return nil, errors.Wrap(err, "возникла ошибка при поиске события по ID")
	}
}

// Создает новое событие
func (r *EventRepository) Create(event repository.Event) (*repository.Event, error) {
	res, err := r.DB.NamedExecContext(context.Background(), queryInsertEvent, map[string]interface{}{
		":title":            event.Title,
		":datetime":         event.Datetime,
		":duration":         event.Duration,
		":description":      event.Description,
		":user_id":          event.UserId,
		":time_send_notify": event.TimeSendNotify,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ошибка во время создания нового события")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "не удалось получить первичный ключ")
	}

	event.ID = int(id)

	return &event, nil
}

// Обновляет событие
func (r *EventRepository) Update(ID int, event repository.Event) (*repository.Event, error) {
	return nil, nil
}

// Удаляет событие
func (r *EventRepository) Delete(ID int) error {
	return nil
}

// Ищет события на день
func (r *EventRepository) FindOnDay(day time.Time) ([]repository.Event, error) {
	return nil, nil
}

// Ищет события на неделю
func (r *EventRepository) FindOnWeek(week time.Weekday) ([]repository.Event, error) {
	return nil, nil
}

// Ищет события на месяц
func (r *EventRepository) FindOnMonth(month time.Month) ([]repository.Event, error) {
	return nil, nil
}
