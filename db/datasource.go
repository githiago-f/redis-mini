package db

import (
	"sync"

	"github.com/githiago-f/redis-mini/protocol"
)

type Datasource struct {
	mu          sync.Mutex
	values      map[string]*protocol.Value
	expirations map[string]int
}

func New() *Datasource {
	return &Datasource{values: map[string]*protocol.Value{}}
}

func (db *Datasource) Lock() {
	db.mu.Lock()
}

func (db *Datasource) Unlock() {
	db.mu.Unlock()
}

func (db *Datasource) Get(key string) (*protocol.Value, bool) {
	val, exists := db.values[key]
	return val, exists
}

func (db *Datasource) Set(key string, val *protocol.Value) {
	db.values[key] = val
}

func (db *Datasource) SetTimeout(key string, seconds int) {
	db.expirations[key] = seconds
}

func (db *Datasource) Del(key string) {
	delete(db.values, key)
}
