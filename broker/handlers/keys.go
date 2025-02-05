package handlers

import (
	"github.com/githiago-f/redis-mini/broker/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func KeysHandler(db *db.InMemory, data []*protocol.Arg) ([]*protocol.Value, error) {
	return nil, nil
}
