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

	query := `select * from events where time_send_notify > $1 `

	rows, err := r.db.NamedQueryContext(ctx, query, t)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка во время поиска событий по дате")
	}

	events := make([]Event, 0)

	for rows.Next() {
		var event Event
		err := rows.StructScan(&event)
		if err != nil {
			return nil, errors.Wrap(err, "ошибка во время сканировании результатов структуры")
		}

		events = append(events, event)
	}

	return events, nil
}

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
