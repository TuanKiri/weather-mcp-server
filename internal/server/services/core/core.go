package core

import "github.com/TuanKiri/weather-mcp-server/internal/server/services"

type CoreServices struct {
	renderer   services.TemplateRenderer
	weatherAPI services.WeatherAPIProvider

	weatherService *WeatherService
}

func New(renderer services.TemplateRenderer, weatherAPI services.WeatherAPIProvider) *CoreServices {
	return &CoreServices{
		renderer:   renderer,
		weatherAPI: weatherAPI,
	}
}

func (cs *CoreServices) Weather() services.WeatherService {
	if cs.weatherService == nil {
		cs.weatherService = &WeatherService{CoreServices: cs}
	}

	return cs.weatherService
}
