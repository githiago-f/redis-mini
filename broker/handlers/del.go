package handlers

import (
	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

func DelHandler(b *broker.Broker, data []*protocol.Arg) ([]*protocol.Value, error) {
	core.Logger.Debugf("Deleting %v keys", len(data))

	b.Lock()
	defer b.Unlock()

	count := 0
	for _, key := range data {
		varName, err := key.AsID()
		if err == nil {
			core.Logger.Debugf("Deleting key %v", varName)
			b.Del(varName)
			count++
		}
	}

	return protocol.NewValue(count).Collect(), nil
}
