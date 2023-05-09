package bot

import (
	"SimaYagodkaBot/internal/bot/handler"
	"SimaYagodkaBot/internal/config"
	"SimaYagodkaBot/internal/service"
	"github.com/and3rson/telemux/v2"
	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	api  *tgbotapi.BotAPI
	conf config.Config
}

func New(conf config.Config) *Bot {
	api, _ := tgbotapi.NewBotAPI(conf.Bot.Telegram.Token)

	if conf.Bot.Debug {
		api.Debug = true
	}

	return &Bot{
		api:  api,
		conf: conf,
	}
}

func (bot *Bot) Run() {
	mux := telemux.NewMux().
		AddHandler(telemux.NewCommandHandler(
			"start",
			telemux.Any(),
			func(update *telemux.Update) {
				bot.api.Send(tgbotapi.NewMessage(
					update.EffectiveChat().ID,
					`
Привет, я бот SimaYagodkaBot

Если не знаешь что делать, пиши /help
`,
				))
			},
		)).
		AddHandler(telemux.NewCommandHandler(
			"help",
			telemux.Any(),
			func(update *telemux.Update) {
				bot.api.Send(tgbotapi.NewMessage(
					update.EffectiveChat().ID,
					`
Напишите название города, чтобы узнать его текущую погоду!
`,
				))
			},
		)).
		SetRecover(func(update *telemux.Update, err error, s string) {
			chat := update.EffectiveChat()
			if chat != nil {
				bot.api.Send(tgbotapi.NewMessage(
					chat.ID,
					"Упс, произошла ошибка!",
				))

				log.Printf("Warning! An error occurred: %s", err)
			}
		})

	currentWeatherData, err := openweathermap.NewCurrent("C", "ru", bot.conf.OpenWeatherMap.APIKey)
	if err != nil {
		log.Fatalln(err)
	}

	weatherService := service.NewWeatherService(currentWeatherData)

	handler.NewWeatherHandler(bot.api, weatherService).
		Register(mux)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updatesChannel := bot.api.GetUpdatesChan(updateConfig)

	for update := range updatesChannel {
		mux.Dispatch(bot.api, update)
	}
}
