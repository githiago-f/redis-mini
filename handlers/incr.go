package handlers

import (
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func IncrHandler(db *db.Datasource, data []*protocol.Arg) ([]*protocol.Value, error) {
	varName, err := data[0].AsID()
	if err != nil {
		return nil, err
	}

	val, ok := db.Get(varName)
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
		db.Set(varName, val)
	}

	return val.Collect(), nil
}
