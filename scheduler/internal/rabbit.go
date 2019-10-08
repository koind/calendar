package internal

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Отправляет сообщения в очередь
func SendMessagesToQueue(ctx context.Context, conn amqp.Connection, events []Event) error {
	if ctx.Err() == context.Canceled {
		return errors.New("отправка уведомлений было прервано из-за отмены контекста")
	}

	ch, err := conn.Channel()
	if err != nil {
		return errors.Wrap(err, "не удалось подключиться к каналу")
	}
	defer ch.Close()

	for _, event := range events {
		data, err := json.Marshal(event)
		if err != nil {
			return errors.Wrap(err, "не удалось закодировать в json")
		}

		err = ch.Publish(
			"notify_about_event",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)
	}

	return nil
}
