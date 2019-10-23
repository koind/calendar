package internal

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

// Путь до конфигураций
var Path string

// Настройки микросервиса
type Options struct {
	Interval   int
	Postgres   Postgres
	RabbitMQ   RabbitMQ
	Prometheus Prometheus
}

// Инициализирует конфигурации микросервиса
func Init(configPath string) Options {
	opt := Options{}

	if _, err := toml.DecodeFile(configPath, &opt); err != nil {
		log.Fatal("Не удалось загрузить конфиги микросервиса ", err)
	}

	return opt
}

// Настройки postgres
type Postgres struct {
	DSN             string
	PingTimeout     int
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Настройки RabbitMQ
type RabbitMQ struct {
	URL string
}

// Настройки Prometheus
type Prometheus struct {
	Port int
}

// Возвращает port
func (c Prometheus) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
