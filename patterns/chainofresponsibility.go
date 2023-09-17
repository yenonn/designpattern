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

type SubstractionHandler struct {
	nextHandler *Handler
}

func (s *SubstractionHandler) Handle(request Request) {
	result := Request{number: request.number - 2}
	log.Print(result)
	if s.nextHandler != nil {
		(*s.nextHandler).Handle(result)
	}
}

func (s *SubstractionHandler) SetNextHandler(handler Handler) {
	s.nextHandler = &handler
}

func ChainOfResponsibility() {
	multiplyHandler := MultiplyHandler{}
	additionHanddler := AdditionalHandler{}
	subtractionHandler := SubstractionHandler{}
	// chaining the responsibility
	multiplyHandler.SetNextHandler(&additionHanddler)
	additionHanddler.SetNextHandler(&subtractionHandler)

	r := Request{number: 2}
	multiplyHandler.Handle(r)
}
