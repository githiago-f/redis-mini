package handlers

import (
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func DelHandler(db *db.Datasource, data []*protocol.Arg) ([]*protocol.Value, error) {
	core.Logger.Debugf("Deleting %v keys", len(data))

	db.Lock()
	defer db.Unlock()

	count := 0
	for _, key := range data {
		varName, err := key.AsID()
		if err == nil {
			core.Logger.Debugf("Deleting key %v", varName)
			db.Del(varName)
			count++
		}
	}

	return protocol.NewValue(count).Collect(), nil
}
