package patterns

import "log"

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

func Observer() {
	p := NewPublisher()
	s0 := NewSubscriber("123")
	s1 := NewSubscriber("234")
	p.addSubscriber(s0)
	p.addSubscriber(s1)
	p.broadcast("hello world")
}
