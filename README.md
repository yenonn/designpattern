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
* pattern [Proxy](/patterns/proxy.go) 
* pattern [Singleton](/patterns/singleton.go) 

## output of the main.go
```bash
❯ go run main.go
Hello design pattern in Golang!
---
github.com/yenonn/learn-patterns/patterns.Adapter
2023/09/23 18:00:28 Temperature in Celsius: 25
2023/09/23 18:00:28 Temperature in Celsius: 20
---
github.com/yenonn/learn-patterns/patterns.Bridge
0.5519999999999999
0.06400000000000002
---
github.com/yenonn/learn-patterns/patterns.Builder
2023/09/23 18:00:28 {24 180 brown}
---
github.com/yenonn/learn-patterns/patterns.ChainOfResponsibility
2023/09/23 18:00:28 {4}
2023/09/23 18:00:28 {6}
2023/09/23 18:00:28 {4}
---
github.com/yenonn/learn-patterns/patterns.Composite
2023/09/23 18:00:28 b1-> l0-> l1-> b0-> l2-> l3
---
github.com/yenonn/learn-patterns/patterns.Decorator
2023/09/23 18:00:28 Soldier stats: Attack(1) Defense(1)
2023/09/23 18:00:28 Soldier stats: Attack(11) Defense(11)
---
github.com/yenonn/learn-patterns/patterns.Decorator2
Total calculate price: $155
---
github.com/yenonn/learn-patterns/patterns.Facade
2023/09/23 18:00:28 true
2023/09/23 18:00:28 false
---
github.com/yenonn/learn-patterns/patterns.Memento
2023/09/23 18:00:28 display life point: 100
2023/09/23 18:00:28 reducing: 75
2023/09/23 18:00:28 reduced life point: 25
2023/09/23 18:00:28 reducing: 50
2023/09/23 18:00:28 reduced life point: -25
2023/09/23 18:00:28 display life point: -25
2023/09/23 18:00:28 Restore the saved person.
2023/09/23 18:00:28 display life point: 25
2023/09/23 18:00:28 Restore the saved person.
2023/09/23 18:00:28 display life point: 100
---
github.com/yenonn/learn-patterns/patterns.Observer
2023/09/23 18:00:28 ID: 123 received message: hello world
2023/09/23 18:00:28 ID: 234 received message: hello world
2023/09/23 18:00:28 Random ID: 36 received message: hello autogen
2023/09/23 18:00:28 Random ID: 27 received message: hello autogen
2023/09/23 18:00:28 Random ID: 96 received message: hello autogen
2023/09/23 18:00:28 Random ID: 28 received message: hello autogen
2023/09/23 18:00:28 Random ID: 98 received message: hello autogen
2023/09/23 18:00:28 Random ID: 47 received message: hello autogen
2023/09/23 18:00:28 Random ID: 74 received message: hello autogen
2023/09/23 18:00:28 Random ID: 39 received message: hello autogen
2023/09/23 18:00:28 Random ID: 25 received message: hello autogen
2023/09/23 18:00:28 Random ID: 52 received message: hello autogen
---
github.com/yenonn/learn-patterns/patterns.Prototype
2023/09/23 18:00:28 Publisher 902536f4-fb8c-4a74-b1e7-33687ee52b18 > hello
2023/09/23 18:00:28 Publisher 902536f4-fb8c-4a74-b1e7-33687ee52b18 > hello from clone
```
