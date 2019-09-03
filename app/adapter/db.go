package adapter

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

// Настройки БД
type DBOptions struct {
	DSN             string
	PingTimeout     int
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Создает пул соединений и возвращает само подключение
func IntDB(options DBOptions) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", options.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "произошла ошибка при создании пула соединений")
	}

	db.SetMaxOpenConns(options.MaxOpenConns)
	db.SetMaxIdleConns(options.MaxIdleConns)
	db.SetConnMaxLifetime(options.ConnMaxLifetime)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(options.PingTimeout)*time.Millisecond,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "произошла ошибка при подключении к базе данных")
	}

	return db, nil
}
