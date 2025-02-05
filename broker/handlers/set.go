package handlers

import (
	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

func SetHandler(b *broker.Broker, data []*protocol.Arg) ([]*protocol.Value, error) {
	variable, err := data[0].AsID()
	value := data[1].Value

	if err != nil {
		return nil, err
	}

	core.Logger.Debugf("Setting %v = '%v'", variable, value)

	b.Lock()
	defer b.Unlock()
	val, exists := b.Get(variable)

	if !exists {
		val = protocol.NewValue(nil)
	}

	val.Value = value
	b.Set(variable, val)

	return protocol.OkValue().Collect(), nil
}
