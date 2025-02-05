package handlers

import (
	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

func IncrHandler(b *broker.Broker, data []*protocol.Arg) ([]*protocol.Value, error) {
	varName, err := data[0].AsID()
	if err != nil {
		return nil, err
	}

	b.Lock()
	defer b.Unlock()
	val, ok := b.Get(varName)
	if !ok {
		val = protocol.NewValue(float64(0))
	}

	switch num := val.Value.(type) {
	default:
		core.Logger.Debugf("Value accessed %v", num)
		return nil, protocol.BadType()
	case float64:
		num++
		val.Value = num
		b.Set(varName, val)
	}

	return val.Collect(), nil
}
