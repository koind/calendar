package model

import "time"

type Event struct {
	UUID           int           `json:"uuid"`                       //UUID - уникальный идентификатор события
	Title          string        `json:"title"`                      //Заголовок - короткий текст
	Datetime       time.Time     `json:"datetime"`                   //Дата и время события
	Duration       time.Duration `json:"duration"`                   //Длительность события
	Description    string        `json:"description,omitempty"`      //Описание события, опционально
	UserId         int           `json:"user_id"`                    //Пользователь, владелец события
	TimeSendNotify time.Time     `json:"time_send_notify,omitempty"` //За сколько времени высылать уведомление, опционально
}
