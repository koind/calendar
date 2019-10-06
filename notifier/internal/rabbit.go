package internal

import (
	"context"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
)

// Обрабатывает сообщения из очереди
func ProcessMessagesFromQueue(ctx context.Context, conn amqp.Connection) error {
	if ctx.Err() == context.Canceled {
		return errors.New("обработка сообщений  было прервано из-за отмены контекста")
	}

	ch, err := conn.Channel()
	if err != nil {
		return errors.Wrap(err, "не удалось подключиться к каналу")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notify_about_event", // name
		"direct",             // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		return errors.Wrap(err, "не удалось объявить exchange")
	}

	q, err := ch.QueueDeclare(
		"events", // name
		true,     // durable
		false,    // delete when unused
		true,     // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return errors.Wrap(err, "не удалось объявить очередь")
	}

	err = ch.QueueBind(
		q.Name,
		"events",
		"notify_about_event",
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "не удалось привязать очередь к обменнику")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		return errors.Wrap(err, "ошибка при приеме сообщений")
	}

	go func() {
		for d := range msgs {
			log.Println(string(d.Body))
		}
	}()

	return nil
}
