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
		"",
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

	for d := range msgs {
		log.Println(string(d.Body))
	}

	return nil
}
