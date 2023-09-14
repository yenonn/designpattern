package main

import (
	"fmt"

	"github.com/yenonn/learn-patterns/patterns"
)

func main() {
	fmt.Print("Hello Adapter design pattern in Golang!\n")
	displayer := patterns.CelsiusTemperaturerDisplayer{}

	staticTemperatureProvider := patterns.StaticTemperatureProvider{}
	displayer.Display(staticTemperatureProvider)

	randomTemperatureProvider := patterns.RandomTemperatureProvider{}
	displayer.Display(randomTemperatureProvider)
}
