package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	OpenWeatherMap OpenWeatherMap
	Bot            Bot
}

type OpenWeatherMap struct {
	APIKey string `env:"OPENWEATHERMAP_API_KEY" env-required:"true"`
}

type Bot struct {
	Debug    bool `env:"BOT_DEBUG"`
	Telegram Telegram
}

type Telegram struct {
	Token string `env:"BOT_TELEGRAM_TOKEN" env-required:"true"`
}

func New() Config {
	var conf Config
	err := cleanenv.ReadConfig(".env", &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
