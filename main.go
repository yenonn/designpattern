package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/yenonn/learn-patterns/patterns"
)

func printDivider() {
	fmt.Println("---")
}

func main() {
	fmt.Println("Hello design pattern in Golang!")

	patterns := []func(){
		patterns.Adapter,
		patterns.Bridge,
		patterns.Builder,
		patterns.ChainOfResponsibility,
		patterns.Composite,
		patterns.Decorator,
		patterns.Decorator2,
		patterns.Facade,
		patterns.Memento,
		patterns.Observer,
		patterns.Prototype,
		patterns.Proxy,
		patterns.Singleton,
		patterns.Configuration,
	}

	for _, pattern := range patterns {
		printDivider()
		fmt.Println(runtime.FuncForPC(reflect.ValueOf(pattern).Pointer()).Name())
		pattern()
	}
}
