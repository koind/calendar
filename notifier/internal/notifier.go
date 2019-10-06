package internal

import (
	"context"
	"github.com/streadway/amqp"
)

func Run(ctx context.Context, conn amqp.Connection) error {
	if err := ProcessMessagesFromQueue(ctx, conn); err != nil {
		return err
	}

	return nil
}
