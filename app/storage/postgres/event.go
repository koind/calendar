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
	querySelectEventByUUID = "SELECT * FROM issues WHERE issue_id=$1"
)

// Конструктор репозитория событий
func NewEventRepository(db *sqlx.DB, logger zap.Logger, ctx context.Context) *EventRepository {
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

// Ищет событие оп UUID
func (r *EventRepository) FindByUUID(UUID int) (*repository.Event, error) {
	ctx, cancel := context.WithTimeout(r.ctx, time.Duration(30)*time.Millisecond)
	defer cancel()

	if ctx.Err() == context.Canceled {
		r.logger.Info(
			"Поиск события по UUID был прерван из-за отмены контекста",
			zap.Int("UUID", UUID),
		)

		return nil, errors.New("поиск события по UUID был прерван из-за отмены контекста")
	}

	row := r.DB.QueryRowContext(ctx, querySelectEventByUUID, UUID)

	event := repository.Event{}
	err := row.Scan(
		&event.UUID,
		&event.Title,
		&event.Datetime,
		&event.Duration,
		&event.Description,
		&event.UserId,
		&event.TimeSendNotify,
	)

	if err != sql.ErrNoRows {
		r.logger.Warn(
			"Не удалось найти событие по UUID",
			zap.Int("UUID", UUID),
		)

		return nil, errors.Wrap(err, "не удалось найти событие по UUID")
	} else {
		r.logger.Warn(
			"Возникла ошибка при поиске события по UUID",
			zap.Int("UUID", UUID),
		)

		return nil, errors.Wrap(err, "возникла ошибка при поиске события по UUID")
	}
}

// Создает новое событие
func (r *EventRepository) Create(event repository.Event) (*repository.Event, error) {
	return nil, nil
}

// Обновляет событие
func (r *EventRepository) Update(UUID int, event repository.Event) (*repository.Event, error) {
	return nil, nil
}

// Удаляет событие
func (r *EventRepository) Delete(UUID int) error {
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
