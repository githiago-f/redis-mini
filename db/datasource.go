package db

import (
	"sync"

	"github.com/githiago-f/redis-mini/protocol"
)

type Datasource struct {
	values      sync.Map
	expirations sync.Map
}

func New() *Datasource {
	return &Datasource{
		values:      sync.Map{},
		expirations: sync.Map{},
	}
}

func (db *Datasource) Get(key string) (*protocol.Value, bool) {
	val, exists := db.values.Load(key)
	return val.(*protocol.Value), exists
}

func (db *Datasource) Set(key string, val *protocol.Value) {
	db.values.Store(key, val)
}

func (db *Datasource) SetTimeout(key string, seconds int) {
	db.expirations.Store(key, seconds)
}

func (db *Datasource) ClearTimeout(key string) {
	db.expirations.Delete(key)
}

func (db *Datasource) Del(key string) {
	db.values.Delete(key)
}
