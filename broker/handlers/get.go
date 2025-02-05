package handlers

import (
	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

func GetHandler(b *broker.Broker, keys []*protocol.Arg) ([]*protocol.Value, error) {
	results := make([]*protocol.Value, len(keys))

	b.Lock()
	defer b.Unlock()
	for i, key := range keys {
		keyValue, err := key.AsID()
		if err != nil {
			continue
		}

		core.Logger.Debugf("Getting %v", keyValue)
		results[i], _ = b.Get(keyValue)
	}

	return results, nil
}
