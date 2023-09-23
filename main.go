package main

import (
	"fmt"

	"github.com/yenonn/learn-patterns/patterns"
)

func printDivider() {
	fmt.Println("---")
}

func main() {
	fmt.Println("Hello design pattern in Golang!")

	printDivider()
	fmt.Println("Adapter pattern")
	patterns.Adapter()

	printDivider()
	fmt.Println("Bridge pattern")
	patterns.Bridge()

	printDivider()
	fmt.Println("Builder pattern")
	patterns.Builder()

	printDivider()
	fmt.Println("Chain of resposibilty pattern")
	patterns.ChainOfResponsibility()

	printDivider()
	fmt.Println("Composite pattern")
	patterns.Composite()

	printDivider()
	fmt.Println("Decorator pattern")
	patterns.Decorator()

	printDivider()
	fmt.Println("Decorator2 pattern")
	patterns.Decorator2()

	printDivider()
	fmt.Println("Facade pattern")
	patterns.Facade()

	printDivider()
	fmt.Println("Memento pattern")
	patterns.Memento()

	printDivider()
	fmt.Println("Observer pattern")
	patterns.Observer()

	printDivider()
	fmt.Println("Prototype pattern")
	patterns.Prototype()
}
