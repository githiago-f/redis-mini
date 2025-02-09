package handlers

import (
	"strconv"

	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func IncrByHandler(db *db.Datasource, data []any) ([]any, error) {
	if len(data) > 2 || len(data) < 1 {
		return nil, protocol.BadArgNumber(1)
	}

	memoryKey, isString := data[0].(string)

	if !isString {
		return nil, protocol.BadType()
	}

	var incrementBy float64 = 1
	if len(data) > 1 {
		switch v := data[1].(type) {
		case int:
			incrementBy = float64(v)
		case float64:
			incrementBy = v
		default:
			return nil, protocol.BadType()
		}
	}

	localValue, _ := db.Values.LoadOrStore(memoryKey, float64(0.0))

	switch v := localValue.(type) {
	default:
		return nil, protocol.BadType()
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		db.Values.Store(memoryKey, f+incrementBy)
		return []any{f + 1}, nil
	case int:
		f := float64(v) + incrementBy
		db.Values.Store(memoryKey, f)
		return []any{f}, nil
	case float64:
		f := v + incrementBy
		db.Values.Store(memoryKey, f)
		return []any{f}, nil
	}
}
