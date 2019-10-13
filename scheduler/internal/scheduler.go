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

	if err := SendMessagesToQueue(ctx, conn, events); err != nil {
		return err
	}

	err = eventRepository.DeleteOldEvents(ctx, time.Now().AddDate(-1, 0, 0))
	if err != nil {
		return err
	}

	return nil
}
