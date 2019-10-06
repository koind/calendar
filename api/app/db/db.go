package db

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/koind/calendar/app/config"
	"github.com/pkg/errors"
)

// Создает пул соединений и возвращает само подключение
func IntPostgres(ctx context.Context, options config.Postgres) (*sqlx.DB, error) {
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
