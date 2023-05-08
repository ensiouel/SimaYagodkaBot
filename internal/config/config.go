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
	APIKey string `env:"OPENWEATHERMAP_API_KEY"`
}

type Bot struct {
	Debug    bool `env:"BOT_DEBUG"`
	Telegram Telegram
}

type Telegram struct {
	Token string `env:"BOT_TELEGRAM_TOKEN"`
}

func New() Config {
	var conf Config
	err := cleanenv.ReadConfig(".env", &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
