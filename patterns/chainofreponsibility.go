package patterns

import "log"

type Handler interface {
	Handle(request Request)
	SetNextHandler(handler Handler)
}

type Request struct {
	number int
}

type MultiplyHandler struct {
	nextHandler *Handler
}

func (m *MultiplyHandler) Handle(request Request) {
	result := Request{number: request.number * 2}
	log.Print(result)
	if m.nextHandler != nil {
		(*m.nextHandler).Handle(result)
	}
}

func (m *MultiplyHandler) SetNextHandler(handler Handler) {
	m.nextHandler = &handler
}

type AdditionalHandler struct {
	nextHandler *Handler
}

func (a *AdditionalHandler) Handle(request Request) {
	result := Request{number: request.number + 2}
	log.Print(result)
	if a.nextHandler != nil {
		(*a.nextHandler).Handle(result)
	}
}

func (a *AdditionalHandler) SetNextHandler(handler Handler) {
	a.nextHandler = &handler
}

func ChainOfResponsibility() {
	m := MultiplyHandler{}
	nextHandler := AdditionalHandler{}
	m.SetNextHandler(&nextHandler)
	r := Request{number: 2}
	m.Handle(r)
}
