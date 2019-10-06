package internal

import (
	"github.com/BurntSushi/toml"
	"log"
)

// Путь до конфигураций
var Path string

// Настройки микросервиса
type Options struct {
	RabbitMQ RabbitMQ
}

// Инициализирует конфигурации микросервиса
func Init(configPath string) Options {
	opt := Options{}

	if _, err := toml.DecodeFile(configPath, &opt); err != nil {
		log.Fatal("Не удалось загрузить конфиги микросервиса ", err)
	}

	return opt
}

// Настройки RabbitMQ
type RabbitMQ struct {
	URL string
}
