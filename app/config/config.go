package config

type Config struct {
	Host          string `env:"CALENDAR_HOST,required"`
	Port          int    `env:"CALENDAR_PORT" envDefault:"7777"`
	ClientTimeout int    `env:"CLIENT_TIMEOUT" envDefault:"400"`
	DSN           string `evn:"DSN"`
}
