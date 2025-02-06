package handlers

import (
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func SetHandler(db *db.Datasource, data []*protocol.Arg) ([]*protocol.Value, error) {
	variable, err := data[0].AsID()
	value := data[1].Value

	if err != nil {
		return nil, err
	}

	core.Logger.Debugf("Setting %v = '%v'", variable, value)

	val, exists := db.Get(variable)

	if !exists {
		val = protocol.NewValue(nil)
	}

	val.Value = value
	db.Set(variable, val)

	return protocol.OkValue().Collect(), nil
}
