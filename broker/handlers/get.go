package handlers

import (
	"github.com/githiago-f/redis-mini/broker/db"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

func GetHandler(db *db.InMemory, args []*protocol.Arg) ([]*protocol.Value, error) {
	return MGetHandler(db, args[0:1])
}

func MGetHandler(db *db.InMemory, keys []*protocol.Arg) ([]*protocol.Value, error) {
	results := make([]*protocol.Value, len(keys))

	db.Lock()
	defer db.Unlock()
	for i, key := range keys {
		keyValue, err := key.AsID()
		if err != nil {
			continue
		}

		core.Logger.Debugf("Getting %v", keyValue)
		results[i], _ = db.Get(keyValue)
	}

	return results, nil
}
