package handlers

import (
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func SetHandler(db *db.Datasource, data []any) ([]any, error) {
	if len(data) != 2 {
		return nil, protocol.BadArgNumber(2)
	}

	variable := data[0].(string)
	value := data[1]

	core.Logger.Debugf("Setting %v = %v", variable, value)

	db.Values.Store(variable, value)

	return []any{"+OK"}, nil
}
