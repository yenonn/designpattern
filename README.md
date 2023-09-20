# Learning Golang Design Pattern
* pattern [Adapter](/patterns/adapter.go)
* pattern [Bridge](/patterns/bridge.go)
* pattern [Builder](/patterns/builder.go)
* pattern [Chain of Responsibility](/patterns/chainofresponsibility.go)
* pattern [Composite](/patterns/composite.go)
* pattern [Decorator](/patterns/decorator.go)
* pattern [Facade](/patterns/facade.go)
* pattern [Memento](/patterns/memento.go)
* pattern [Observer](/patterns/observer.go) 

## output of the main.go
```
❯ go run main.go
Hello Adapter design pattern in Golang!
Adapter pattern
2023/09/20 21:45:06 Temperature in Celsius: 25
2023/09/20 21:45:06 Temperature in Celsius: 25
Bridge pattern
0.5519999999999999
0.06400000000000002
Builder pattern
2023/09/20 21:45:06 {24 180 brown}
Chain of resposibilty pattern
2023/09/20 21:45:06 {4}
2023/09/20 21:45:06 {6}
2023/09/20 21:45:06 {4}
Composite pattern
2023/09/20 21:45:06 b1-> l0-> l1-> b0-> l2-> l3
Decorator pattern
2023/09/20 21:45:06 Soldier stats: Attack(1) Defense(1)
2023/09/20 21:45:06 Soldier stats: Attack(11) Defense(11)
Facade pattern
2023/09/20 21:45:06 true
2023/09/20 21:45:06 false
Memento pattern
2023/09/20 21:45:06 display life point: 100
2023/09/20 21:45:06 reducing: 75
2023/09/20 21:45:06 reduced life point: 25
2023/09/20 21:45:06 reducing: 50
2023/09/20 21:45:06 reduced life point: -25
2023/09/20 21:45:06 display life point: -25
2023/09/20 21:45:06 Restore the saved person.
2023/09/20 21:45:06 display life point: 25
2023/09/20 21:45:06 Restore the saved person.
2023/09/20 21:45:06 display life point: 100
Observer pattern
2023/09/20 21:45:06 ID: 123 received message: hello world
2023/09/20 21:45:06 ID: 234 received message: hello world
```
