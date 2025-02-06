package handlers

import (
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func ExistsHandler(b *db.Datasource, data []*protocol.Arg) ([]*protocol.Value, error) {
	varName, err := data[0].AsID()
	if err != nil {
		return nil, err
	}

	_, exists := b.Get(varName)

	return protocol.NewValue(exists).Collect(), nil
}
