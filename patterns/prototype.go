package patterns

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type (
	IPrototype interface {
		Clone() interface{}
	}

	ThePublisher interface {
		IPrototype
		Publish(message string)
	}
)

type MyPublisher struct {
	id string
}

func NewMyPublisher() *MyPublisher {
	return &MyPublisher{id: uuid.New().String()}
}

func (pub *MyPublisher) Clone() interface{} {
	return &MyPublisher{id: pub.id}
}

func (pub *MyPublisher) Publish(message string) {
	fullMessage := fmt.Sprint("Publisher ", pub.id, " > ", message)
	log.Println(fullMessage)
}

func Prototype() {
	publisher1 := NewMyPublisher()
	run(publisher1, "hello")

	publisher2 := publisher1.Clone().(ThePublisher)
	run(publisher2, "hello from clone")
}

func run(publisher ThePublisher, message string) {
	publisher.Publish(message)
}
