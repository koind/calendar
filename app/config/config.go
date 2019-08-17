package config

type Config struct {
	Host string `env:"CALENDAR_HOST,required"`
	Port int    `env:"CALENDAR_PORT" envDefault:"7777"`
}
