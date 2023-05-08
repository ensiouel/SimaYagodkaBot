package handler

import (
	"SimaYagodkaBot/internal/bot/template"
	"SimaYagodkaBot/internal/service"
	"SimaYagodkaBot/pkg/owmutils"
	"bytes"
	"errors"
	"github.com/and3rson/telemux/v2"
	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var WeekdayReplacer = strings.NewReplacer(
	"Monday", "Понедельник",
	"Tuesday", "Вторник",
	"Wednesday", "Среда",
	"Thursday", "Четверг",
	"Friday", "Пятница",
	"Saturday", "Суббота",
	"Sunday", "Воскресенье")

func FirstTitle(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}

type WeatherHandler struct {
	api            *tgbotapi.BotAPI
	weatherService service.WeatherService
}

func NewWeatherHandler(api *tgbotapi.BotAPI, weatherService service.WeatherService) *WeatherHandler {
	return &WeatherHandler{api: api, weatherService: weatherService}
}

func (handler *WeatherHandler) Register(mux *telemux.Mux) {
	mux.AddHandler(
		telemux.NewMessageHandler(
			telemux.HasText(),
			handler.GetWeather,
		),
	)
}

func (handler *WeatherHandler) GetWeather(update *telemux.Update) {
	chatID := update.EffectiveChat().ID
	text := update.EffectiveMessage().Text

	var message tgbotapi.MessageConfig

	weather, err := handler.weatherService.GetCurrent(text)
	if err != nil {
		message = tgbotapi.NewMessage(chatID, `Не удалось найти город с названием <b>`+text+`</b>`)
	} else {
		message, err = renderMessage(chatID, weather)
		if err != nil {
			panic(err)
		}
	}

	message.ParseMode = tgbotapi.ModeHTML

	if _, err = handler.api.Send(message); err != nil {
		log.Println(err)
	}
}

func renderMessage(chatID int64, weather *openweathermap.CurrentWeatherData) (tgbotapi.MessageConfig, error) {
	if len(weather.Weather) == 0 {
		return tgbotapi.MessageConfig{}, errors.New("no weather info")
	}

	buffer := &bytes.Buffer{}
	err := template.GetWeather.Execute(buffer, map[string]any{
		"date":        WeekdayReplacer.Replace(time.Unix(int64(weather.Dt+weather.Timezone), 0).UTC().Format("Monday, 15:04")),
		"city_name":   weather.Name,
		"country":     weather.Sys.Country,
		"icon":        owmutils.GetIcon(weather.Weather[0].ID),
		"temp":        strconv.Itoa(int(math.Ceil(weather.Main.Temp))),
		"feels_like":  strconv.Itoa(int(math.Ceil(weather.Main.FeelsLike))),
		"description": FirstTitle(weather.Weather[0].Description),
	})
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	message := tgbotapi.NewMessage(
		chatID,
		buffer.String(),
	)

	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ветер "+strconv.FormatFloat(weather.Wind.Speed, 'f', 2, 64)+" м/с", "nil"),
			tgbotapi.NewInlineKeyboardButtonData("Давление "+strconv.FormatFloat(weather.Main.Pressure*0.750064, 'f', 2, 64)+" мм рт.ст", "nil"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Влажность "+strconv.Itoa(weather.Main.Humidity)+"%", "nil"),
			tgbotapi.NewInlineKeyboardButtonData("Видимость "+strconv.Itoa(weather.Visibility/1000)+" км", "nil"),
		),
	)

	return message, nil
}
