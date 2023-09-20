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
2023/09/20 22:00:59 Temperature in Celsius: 25
2023/09/20 22:00:59 Temperature in Celsius: 19
Bridge pattern
0.5519999999999999
0.06400000000000002
Builder pattern
2023/09/20 22:00:59 {24 180 brown}
Chain of resposibilty pattern
2023/09/20 22:00:59 {4}
2023/09/20 22:00:59 {6}
2023/09/20 22:00:59 {4}
Composite pattern
2023/09/20 22:00:59 b1-> l0-> l1-> b0-> l2-> l3
Decorator pattern
2023/09/20 22:00:59 Soldier stats: Attack(1) Defense(1)
2023/09/20 22:00:59 Soldier stats: Attack(11) Defense(11)
Facade pattern
2023/09/20 22:00:59 true
2023/09/20 22:00:59 false
Memento pattern
2023/09/20 22:00:59 display life point: 100
2023/09/20 22:00:59 reducing: 75
2023/09/20 22:00:59 reduced life point: 25
2023/09/20 22:00:59 reducing: 50
2023/09/20 22:00:59 reduced life point: -25
2023/09/20 22:00:59 display life point: -25
2023/09/20 22:00:59 Restore the saved person.
2023/09/20 22:00:59 display life point: 25
2023/09/20 22:00:59 Restore the saved person.
2023/09/20 22:00:59 display life point: 100
Observer pattern
2023/09/20 22:00:59 ID: 123 received message: hello world
2023/09/20 22:00:59 ID: 234 received message: hello world
2023/09/20 22:00:59 Random ID: 94 received message: hello autogen
2023/09/20 22:00:59 Random ID: 78 received message: hello autogen
2023/09/20 22:00:59 Random ID: 97 received message: hello autogen
2023/09/20 22:00:59 Random ID: 55 received message: hello autogen
2023/09/20 22:00:59 Random ID: 21 received message: hello autogen
2023/09/20 22:00:59 Random ID: 86 received message: hello autogen
2023/09/20 22:00:59 Random ID: 59 received message: hello autogen
2023/09/20 22:00:59 Random ID: 98 received message: hello autogen
2023/09/20 22:00:59 Random ID: 51 received message: hello autogen
2023/09/20 22:00:59 Random ID: 35 received message: hello autogen
```
