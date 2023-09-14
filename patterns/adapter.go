package patterns

import (
	"log"
	"math/rand"
)

type (
	Celsius                     int
	Fahrenheit                  int
	CelsiusTemperatureDisplayer struct{}
)

func (displayer CelsiusTemperatureDisplayer) Display(temperatureProvider InterfaceTemperatureProvider) {
	log.Printf("Temperature in Celsius: %v", temperatureProvider.GetCelsius())
}

type InterfaceTemperatureProvider interface {
	GetCelsius() Celsius
}

type CelsiusTemperatureProvider struct{}

func (temp CelsiusTemperatureProvider) GetCelsius() Celsius {
	return 25
}

// completely different provider with another implementation
type FahrenheitTemperatureProvider struct{}

func (temp FahrenheitTemperatureProvider) GetFahrenheit() Fahrenheit {
	// random temperature
	min := 50
	max := 65
	return Fahrenheit(rand.Intn(max-min) + min)
}

// adapter implementation from fahrenheit to celsius
// implementing GetCelsius
type FahrenheitToCelsiusAdapter struct{ TemperatureProvider FahrenheitTemperatureProvider }

func (adapter FahrenheitToCelsiusAdapter) GetCelsius() Celsius {
	celsius := Celsius(adapter.TemperatureProvider.GetFahrenheit()) - 32
	return celsius
}

// main adapter usecase
func Adapter() {
	displayer := CelsiusTemperatureDisplayer{}

	celsiusTemperatureProvider := CelsiusTemperatureProvider{}
	displayer.Display(celsiusTemperatureProvider)

	adapter := FahrenheitToCelsiusAdapter{
		TemperatureProvider: FahrenheitTemperatureProvider{},
	}

	displayer.Display(adapter)
}
