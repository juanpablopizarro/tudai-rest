package greeter

import "fmt"

// Greeter service interface
type Greeter interface {
	SayHello(string) string
}

type greeter struct{}

// NewGreeter ...
func NewGreeter() Greeter {
	return greeter{}
}

func (f greeter) SayHello(s string) string {
	return fmt.Sprintf("Hello %v!", s)
}
