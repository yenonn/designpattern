package patterns

import (
	"log"
	"math/rand"
)

type (
	Celsius                      int
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

type RandomTemperatureProvider struct{}

func (temp RandomTemperatureProvider) GetCelsius() Celsius {
	min := 20
	max := 35
	return Celsius(rand.Intn(max-min) + min)
}
