package main

import (
	"fmt"

	"github.com/yenonn/learn-patterns/patterns"
)

func main() {
	fmt.Println("Hello design pattern in Golang!")
	fmt.Println("Adapter pattern")
	patterns.Adapter()

	fmt.Println("Bridge pattern")
	patterns.Bridge()

	fmt.Println("Builder pattern")
	patterns.Builder()

	fmt.Println("Chain of resposibilty pattern")
	patterns.ChainOfResponsibility()

	fmt.Println("Composite pattern")
	patterns.Composite()

	fmt.Println("Decorator pattern")
	patterns.Decorator()

	fmt.Println("Decorator2 pattern")
	patterns.Decorator2()

	fmt.Println("Facade pattern")
	patterns.Facade()

	fmt.Println("Memento pattern")
	patterns.Memento()

	fmt.Println("Observer pattern")
	patterns.Observer()
}
