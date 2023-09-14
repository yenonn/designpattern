package patterns

import (
	"log"
	"math/rand"
)

type (
	Celsius                      int
	Fahrenheit                   int
	CelsiusTemperaturerDisplayer struct{}
)

func (displayer CelsiusTemperaturerDisplayer) Display(temperatureProvider InterfaceTemperatureProvider) {
	log.Printf("Temperaturer in Celsius: %v", temperatureProvider.GetCelsius())
}

type InterfaceTemperatureProvider interface {
	GetCelsius() Celsius
}

type StaticTemperatureProvider struct{}

func (temp StaticTemperatureProvider) GetCelsius() Celsius {
	return 25
}

// completely different provider with another implementation
type RandomTemperatureProvider struct{}

func (temp RandomTemperatureProvider) GetFahrenheit() Fahrenheit {
	min := 50
	max := 65
	return Fahrenheit(rand.Intn(max-min) + min)
}

// adapter implementation from fahrenheit to celsius
type FahrenheitToCelsiusAdapter struct{ TemperatureProvider RandomTemperatureProvider }

func (adapter FahrenheitToCelsiusAdapter) GetCelsius() Celsius {
	celsius := Celsius(adapter.TemperatureProvider.GetFahrenheit()) - 32
	return celsius
}
