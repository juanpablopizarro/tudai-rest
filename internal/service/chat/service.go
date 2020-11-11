package chat

import "github.com/juanpablopizarro/tudai-rest/internal/config"

// Message ...
type Message struct {
	ID   int64
	Text string
}

// ChatService ...
type ChatService interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}

type service struct {
	conf *config.Config
}

// New ...
func New(c *config.Config) (ChatService, error) {
	return service{c}, nil
}

func (s service) AddMessage(m Message) error {
	return nil
}

func (s service) FindByID(ID int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	list = append(list, &Message{0, "Hello World"})
	return list
}
