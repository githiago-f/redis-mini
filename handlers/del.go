package handlers

import (
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
)

func DelHandler(db *db.Datasource, data []*any) ([]int, error) {
	core.Logger.Debugf("Deleting %v keys", len(data))

	count := 0
	// for _, key := range data {
	// 	// varName, err := key.(string)
	// 	// if err == nil {
	// 	// 	core.Logger.Debugf("Deleting key %v", varName)
	// 	// 	db.Values.Delete(varName)
	// 	// 	count++
	// 	// }
	// }

	return []int{count}, nil
}
