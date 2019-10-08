package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/koind/calendar/api/app/domain/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	querySelectEventByID = `SELECT * FROM events WHERE id=$1`
	queryInsertEvent     = `INSERT INTO events(title, datetime, duration, description, user_id, time_send_notify)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	queryUpdateEventByID = `UPDATE events SET title=$1 WHERE id=$2`
	queryDeleteEventByID = `DELETE FROM events WHERE id=$1`
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
func (r *EventRepository) FindByID(ctx context.Context, ID int) (*repository.Event, error) {
	if ctx.Err() == context.Canceled {
		r.logger.Info(
			"Поиск события по ID был прерван из-за отмены контекста",
			zap.Int("ID", ID),
		)

		return nil, errors.New("поиск события по ID был прерван из-за отмены контекста")
	}

	row := r.DB.QueryRowContext(context.Background(), querySelectEventByID, ID)

	event := new(repository.Event)
	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Datetime,
		&event.Duration,
		&event.Description,
		&event.UserId,
		&event.TimeSendNotify,
	)

	if err == sql.ErrNoRows {
		r.logger.Warn(
			"Не удалось найти событие по ID",
			zap.Int("ID", ID),
		)

		return nil, errors.Wrap(err, "не удалось найти событие по ID")
	} else if err != nil {
		r.logger.Warn(
			"Возникла ошибка при поиске события по ID",
			zap.Error(err),
			zap.Int("ID", ID),
		)

		return nil, errors.Wrap(err, "возникла ошибка при поиске события по ID")
	}

	return event, nil
}

// Создает новое событие
func (r *EventRepository) Create(ctx context.Context, event repository.Event) (*repository.Event, error) {
	if ctx.Err() == context.Canceled {
		r.logger.Info(
			"Создание нового события было прервано из-за отмены контекста",
			zap.Int("ID", event.ID),
		)

		return nil, errors.New("создание нового события было прервано из-за отмены контекста")
	}

	err := r.DB.QueryRowContext(
		context.Background(),
		queryInsertEvent,
		event.Title,
		event.Datetime,
		event.Duration,
		event.Description,
		event.UserId,
		event.TimeSendNotify,
	).Scan(&event.ID)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка во время создания нового события")
	}

	return &event, nil
}

// Обновляет событие
func (r *EventRepository) Update(ctx context.Context, ID int, event repository.Event) (*repository.Event, error) {
	if ctx.Err() == context.Canceled {
		r.logger.Info(
			"Обновление события было прервано из-за отмены контекста",
			zap.Int("ID", ID),
		)

		return nil, errors.New("обновление события было прервано из-за отмены контекста")
	}

	_, err := r.DB.ExecContext(context.Background(), queryUpdateEventByID, event.Title, ID)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка во время обновления события")
	}

	newEvent, err := r.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Удаляет событие
func (r *EventRepository) Delete(ctx context.Context, ID int) error {
	if ctx.Err() == context.Canceled {
		r.logger.Info(
			"Удаление события было прервано из-за отмены контекста",
			zap.Int("ID", ID),
		)

		return errors.New("Удаление события было прервано из-за отмены контекста")
	}

	_, err := r.DB.ExecContext(context.Background(), queryDeleteEventByID, ID)
	if err != nil {
		return errors.Wrap(err, "ошибка во время удаления события")
	}

	return nil
}
