package chat

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanpablopizarro/tudai-rest/internal/config"
)

// Message ...
type Message struct {
	ID   int64
	Text string
}

// Service ...
type Service interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(m Message) error {
	return nil
}

func (s service) FindByID(ID int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	if err := s.db.Select(&list, "SELECT * FROM messages"); err != nil {
		panic(err)
	}
	return list
}
