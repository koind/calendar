package internal

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

// Создает пул соединений и возвращает само подключение
func IntPostgres(ctx context.Context, options Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", options.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "произошла ошибка при создании пула соединений")
	}

	db.SetMaxOpenConns(options.MaxOpenConns)
	db.SetMaxIdleConns(options.MaxIdleConns)
	db.SetConnMaxLifetime(options.ConnMaxLifetime)

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "произошла ошибка при подключении к базе данных")
	}

	return db, nil
}

// Интерфейс репозитория событий
type EventRepositoryInterface interface {
	// Ищет события по времени отправки уведомления
	FindByTimeOfSendNotify(ctx context.Context, t time.Time) ([]Event, error)

	// Удаляет старые события
	DeleteOldEvents(ctx context.Context, t time.Time) error
}

// Конструктор postgres репозитория событий
func NewPostgresEventRepository(db *sqlx.DB) *PostgresEventRepository {
	return &PostgresEventRepository{
		db: db,
	}
}

// Postgres репозитория событий
type PostgresEventRepository struct {
	db *sqlx.DB
}

// Ищет события по времени отправки уведомления
func (r *PostgresEventRepository) FindByTimeOfSendNotify(ctx context.Context, t time.Time) ([]Event, error) {
	if ctx.Err() == context.Canceled {
		return nil, errors.New("поиск событий по дате был прерван из-за отмены контекста")
	}

	events := make([]Event, 0)
	query := `select * from events where time_send_notify > $1`

	err := r.db.SelectContext(ctx, &events, query, t)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка во время поиска событий по дате")
	}

	return events, nil
}

// Удаляет старые события
func (r *PostgresEventRepository) DeleteOldEvents(ctx context.Context, t time.Time) error {
	if ctx.Err() == context.Canceled {
		return errors.New("удаление старых событий было прервано из-за отмены контекста")
	}

	query := `delete from events where datetime < $1 `

	_, err := r.db.ExecContext(ctx, query, t)
	if err != nil {
		return errors.Wrap(err, "ошибка во время удаления старых событий")
	}

	return nil
}

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
