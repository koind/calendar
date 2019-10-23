package config

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
	Postgres   Postgres
	GRPCServer GRPCServer
	HTTPServer HTTPServer
	GRPCClient GRPCClient
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

// Настройки gRPC сервера
type GRPCServer struct {
	Host string
	Port int
}

// Возвращает домен
func (s GRPCServer) GetDomain() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Настройки HTTP сервера
type HTTPServer struct {
	Host string
	Port int
}

// Возвращает домен
func (s HTTPServer) GetDomain() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Настройки gRPC клиента
type GRPCClient struct {
	Host    string
	Port    int
	Timeout int
}

// Возвращает домен
func (c GRPCClient) GetDomain() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Возвращает таймаут
func (c GRPCClient) GetTimeout() time.Duration {
	return time.Duration(c.Timeout) * time.Millisecond
}

// Настройки Prometheus
type Prometheus struct {
	Port int
}

// Возвращает port
func (c Prometheus) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
