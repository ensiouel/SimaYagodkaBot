package service

import (
	"github.com/briandowns/openweathermap"
)

type WeatherService interface {
	GetCurrent(cityName string) (*openweathermap.CurrentWeatherData, error)
}

type weatherService struct {
	currentWeatherData *openweathermap.CurrentWeatherData
}

func NewWeatherService(currentWeatherData *openweathermap.CurrentWeatherData) WeatherService {
	return &weatherService{currentWeatherData: currentWeatherData}
}

func (service *weatherService) GetCurrent(cityName string) (*openweathermap.CurrentWeatherData, error) {
	err := service.currentWeatherData.CurrentByName(cityName)
	if err != nil {
		return nil, err
	}

	return service.currentWeatherData, nil
}
