package handlers

import (
	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/protocol"
)

func ExistsHandler(b *broker.Broker, data []*protocol.Arg) ([]*protocol.Value, error) {
	varName, err := data[0].AsID()
	if err != nil {
		return nil, err
	}

	_, exists := b.Get(varName)

	return protocol.NewValue(exists).Collect(), nil
}
