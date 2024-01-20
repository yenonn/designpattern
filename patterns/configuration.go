package patterns

import (
	"fmt"
)

type Server struct {
	opts Opts
}

type Opts struct {
	tls            bool
	id             string
	maxConnections int
}

type OptsFunc func(opts *Opts)

func WithTLS() OptsFunc {
	return func(opts *Opts) {
		opts.tls = true
	}
}

func WithID(id string) OptsFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func WithMaxConnections(maxConnections int) OptsFunc {
	return func(opts *Opts) {
		opts.maxConnections = maxConnections
	}
}

func NewServer(opts ...OptsFunc) *Server {
	defaultOpts := Opts{tls: false, id: "default", maxConnections: 1}
	for _, opt := range opts {
		opt(&defaultOpts)
	}
	return &Server{opts: defaultOpts}
}

func Configuration() {
	fmt.Println("Configuration")
	defaultServer := NewServer()
	fmt.Println(defaultServer)

	tlsServer := NewServer(WithTLS())
	fmt.Println(tlsServer)

	idServer := NewServer(WithID("server-0001"))
	fmt.Println(idServer)

	maxServer := NewServer(WithMaxConnections(10))
	fmt.Println(maxServer)

	customServer := NewServer(WithTLS(), WithID("server-0001"), WithMaxConnections(10))
	fmt.Println(customServer)
}
