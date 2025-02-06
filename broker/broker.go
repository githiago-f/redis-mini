package broker

import (
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

type Broker struct {
	db *db.Datasource

	handlers map[string]Handler
}

type Handler func(mediator *db.Datasource, args []*protocol.Arg) ([]*protocol.Value, error)

func New(db *db.Datasource) *Broker {
	return &Broker{
		db:       db,
		handlers: map[string]Handler{},
	}
}

func (broker *Broker) Handle(data string) ([]*protocol.Value, error) {
	command, args := protocol.Lex(data)
	core.Logger.Debugf("Command %v", command)

	commandHandler, exists := broker.handlers[command]
	if !exists {
		return nil, protocol.BadCommand(data)
	}

	return commandHandler(broker.db, args)
}

func (broker *Broker) Use(command string, fn Handler) {
	broker.handlers[command] = fn
}
