package internal

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	"time"
)

func Run(ctx context.Context, db *sqlx.DB, conn amqp.Connection) error {
	eventRepository := NewPostgresEventRepository(db)

	events, err := eventRepository.FindByTimeOfSendNotify(ctx, time.Now())
	if err != nil {
		return err
	}

	if len(events) < 0 {
		return nil
	}

	if err := SendNotify(ctx, conn, events); err != nil {
		return err
	}

	return nil
}
