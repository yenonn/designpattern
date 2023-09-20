package patterns

import (
	"fmt"
	"log"
	"math/rand"
)

type IPublisher interface {
	addSubscriber(subscriber ISubscriber)
	removeSubscriber(subscriberId string)
	broadcast(message string)
}

type ISubscriber interface {
	id() string
	react(msg string)
}

// Implementing IPublisher
type publisher struct {
	subscribers map[string]ISubscriber
}

func NewPublisher() *publisher {
	return &publisher{subscribers: make(map[string]ISubscriber)}
}

func (p *publisher) addSubscriber(subscriber ISubscriber) {
	p.subscribers[subscriber.id()] = subscriber
}

func (p *publisher) removeSubscriber(subscriberId string) {
	delete(p.subscribers, subscriberId)
}

func (p *publisher) broadcast(message string) {
	for _, sub := range p.subscribers {
		sub.react(message)
	}
}

type subscriber struct {
	subId string
}

func NewSubscriber(subId string) *subscriber {
	return &subscriber{subId: subId}
}

func (s *subscriber) id() string {
	return s.subId
}

func (s *subscriber) react(message string) {
	log.Printf("ID: %s received message: %s", s.subId, message)
}

type autoGenerateSubcriberId struct {
	subId string
}

func NewAutoGenerateSubscriberId() *autoGenerateSubcriberId {
	randomSubId := fmt.Sprint(rand.Intn(100))
	return &autoGenerateSubcriberId{subId: randomSubId}
}

func (a *autoGenerateSubcriberId) id() string {
	return a.subId
}

func (a *autoGenerateSubcriberId) react(message string) {
	log.Printf("Random ID: %s received message: %s", a.subId, message)
}

func Observer() {
	p := NewPublisher()
	s0 := NewSubscriber("123")
	s1 := NewSubscriber("234")
	p.addSubscriber(s0)
	p.addSubscriber(s1)
	p.broadcast("hello world")

	p1 := NewPublisher()
	for i := 0; i < 10; i++ {
		p1.addSubscriber(NewAutoGenerateSubscriberId())
	}
	p1.broadcast("hello autogen")
}
