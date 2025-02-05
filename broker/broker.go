package broker

import (
	"sync"

	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

type Broker struct {
	mu     sync.Mutex
	values map[string]*protocol.Value

	handlers map[string]Handler
}

type Handler func(mediator *Broker, args []*protocol.Arg) ([]*protocol.Value, error)

func New() *Broker {
	return &Broker{
		values:   map[string]*protocol.Value{},
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

	return commandHandler(broker, args)
}

func (broker *Broker) Use(command string, fn Handler) {
	broker.handlers[command] = fn
}

func (b *Broker) Lock() {
	b.mu.Lock()
}

func (b *Broker) Unlock() {
	b.mu.Unlock()
}

func (b *Broker) Get(key string) (*protocol.Value, bool) {
	val, exists := b.values[key]
	return val, exists
}

func (b *Broker) Set(key string, val *protocol.Value) {
	b.values[key] = val
}

func (b *Broker) Del(key string) {
	delete(b.values, key)
}
