package main

import (
	"fmt"

	"github.com/yenonn/learn-patterns/patterns"
)

func main() {
	fmt.Print("Hello Adapter design pattern in Golang!\n")
	displayer := patterns.CelsiusTemperaturerDisplayer{}

	celsiusTemperatureProvider := patterns.CelsiusTemperatureProvider{}
	displayer.Display(celsiusTemperatureProvider)

	adapter := patterns.FahrenheitToCelsiusAdapter{
		TemperatureProvider: patterns.FahrenheitTemperatureProvider{},
	}

	displayer.Display(adapter)
}
