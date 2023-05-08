package app

import (
	"SimaYagodkaBot/internal/bot"
	"SimaYagodkaBot/internal/config"
)

type App struct{}

func New() *App {
	return &App{}
}

func (app *App) Run() {
	conf := config.New()

	bot.New(conf).
		Run()
}
