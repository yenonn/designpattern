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

	IBroadcaster interface {
		IPrototype
		Broadcast(message string)
	}
)

type MyBroadcaster struct {
	id string
}

func NewMyBroadcaster() *MyBroadcaster {
	return &MyBroadcaster{id: uuid.New().String()}
}

func (b *MyBroadcaster) Clone() interface{} {
	return &MyBroadcaster{id: b.id}
}

func (b *MyBroadcaster) Broadcast(message string) {
	fullMessage := fmt.Sprint("Publisher ", b.id, " > ", message)
	log.Println(fullMessage)
}

func Prototype() {
	b1 := NewMyBroadcaster()
	run(b1, "hello")

	b2 := b1.Clone().(IBroadcaster)
	run(b2, "hello from clone")
}

func run(broadcaster IBroadcaster, message string) {
	broadcaster.Broadcast(message)
}
