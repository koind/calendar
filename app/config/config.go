package config

type Config struct {
	Host          string `env:"CALENDAR_HOST" envDefault:"0.0.0.0"`
	Port          int    `env:"CALENDAR_PORT" envDefault:"7777"`
	ClientTimeout int    `env:"CLIENT_TIMEOUT" envDefault:"400"`
	DSN           string `evn:"DSN" envDefault:"400"`
}
