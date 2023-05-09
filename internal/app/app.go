package app

import (
	"SimaYagodkaBot/internal/bot"
	"SimaYagodkaBot/internal/config"
)

type App struct {
	conf config.Config
}

func New() *App {
	conf := config.New()

	return &App{
		conf: conf,
	}
}

func (app *App) Run() {
	bot.New(app.conf).
		Run()
}
