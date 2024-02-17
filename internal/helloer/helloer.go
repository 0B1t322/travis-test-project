package helloer

import "fmt"

// Helloer ...
type Helloer struct {
	To string
}

// SayHello say hello to
func (h Helloer) SayHello() string {
	return fmt.Sprintf("say hello: %s", h.To)
}

// New create new helloer
func New(to string) *Helloer {
	return &Helloer{
		To: to,
	}
}
