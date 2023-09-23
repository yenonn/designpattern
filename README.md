# Learning Golang Design Pattern
* pattern [Adapter](/patterns/adapter.go)
* pattern [Bridge](/patterns/bridge.go)
* pattern [Builder](/patterns/builder.go)
* pattern [Chain of Responsibility](/patterns/chainofresponsibility.go)
* pattern [Composite](/patterns/composite.go)
* pattern [Decorator](/patterns/decorator.go)
* pattern [Decorator 2](/patterns/decorator_example_2.go)
* pattern [Facade](/patterns/facade.go)
* pattern [Memento](/patterns/memento.go)
* pattern [Observer](/patterns/observer.go) 
* pattern [Prototype](/patterns/prototype.go) 

## output of the main.go
```bash
❯ go run main.go
Hello design pattern in Golang!
Adapter pattern
2023/09/22 21:19:55 Temperature in Celsius: 25
2023/09/22 21:19:55 Temperature in Celsius: 19
Bridge pattern
0.5519999999999999
0.06400000000000002
Builder pattern
2023/09/22 21:19:55 {24 180 brown}
Chain of resposibilty pattern
2023/09/22 21:19:55 {4}
2023/09/22 21:19:55 {6}
2023/09/22 21:19:55 {4}
Composite pattern
2023/09/22 21:19:55 b1-> l0-> l1-> b0-> l2-> l3
Decorator pattern
2023/09/22 21:19:55 Soldier stats: Attack(1) Defense(1)
2023/09/22 21:19:55 Soldier stats: Attack(11) Defense(11)
Decorator2 pattern
Total calculate price: $155
Facade pattern
2023/09/22 21:19:55 true
2023/09/22 21:19:55 false
Memento pattern
2023/09/22 21:19:55 display life point: 100
2023/09/22 21:19:55 reducing: 75
2023/09/22 21:19:55 reduced life point: 25
2023/09/22 21:19:55 reducing: 50
2023/09/22 21:19:55 reduced life point: -25
2023/09/22 21:19:55 display life point: -25
2023/09/22 21:19:55 Restore the saved person.
2023/09/22 21:19:55 display life point: 25
2023/09/22 21:19:55 Restore the saved person.
2023/09/22 21:19:55 display life point: 100
Observer pattern
2023/09/22 21:19:55 ID: 123 received message: hello world
2023/09/22 21:19:55 ID: 234 received message: hello world
2023/09/22 21:19:55 Random ID: 56 received message: hello autogen
2023/09/22 21:19:55 Random ID: 7 received message: hello autogen
2023/09/22 21:19:55 Random ID: 53 received message: hello autogen
2023/09/22 21:19:55 Random ID: 42 received message: hello autogen
2023/09/22 21:19:55 Random ID: 79 received message: hello autogen
2023/09/22 21:19:55 Random ID: 36 received message: hello autogen
2023/09/22 21:19:55 Random ID: 43 received message: hello autogen
2023/09/22 21:19:55 Random ID: 0 received message: hello autogen
2023/09/22 21:19:55 Random ID: 58 received message: hello autogen
2023/09/22 21:19:55 Random ID: 35 received message: hello autogen
```
