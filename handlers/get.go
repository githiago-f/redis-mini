package handlers

import "github.com/githiago-f/redis-mini/db"

func GetHandler(db *db.Datasource, args []any) ([]any, error) {
	return MGetHandler(db, args[0:1])
}

func MGetHandler(db *db.Datasource, keys []any) ([]any, error) {
	results := make([]any, len(keys))

	for i, key := range keys {
		switch k := key.(type) {
		case string:
			results[i], _ = db.Values.Load(k)
		}
	}

	return results, nil
}
