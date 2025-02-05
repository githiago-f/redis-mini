package db

import (
	"sync"

	"github.com/githiago-f/redis-mini/protocol"
)

type InMemory struct {
	mu     sync.Mutex
	values map[string]*protocol.Value
}

func New() *InMemory {
	return &InMemory{values: map[string]*protocol.Value{}}
}

func (db *InMemory) Lock() {
	db.mu.Lock()
}

func (db *InMemory) Unlock() {
	db.mu.Unlock()
}

func (db *InMemory) Get(key string) (*protocol.Value, bool) {
	val, exists := db.values[key]
	return val, exists
}

func (db *InMemory) Set(key string, val *protocol.Value) {
	db.values[key] = val
}

func (db *InMemory) Del(key string) {
	delete(db.values, key)
}
