package handlers

import (
	"fmt"

	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func HELLOHandler(db *db.Datasource, data []any) ([]any, error) {
	if len(data) < 1 {
		return nil, protocol.BadArgNumber(1)
	}

	switch version := data[0].(type) {
	default:
		return nil, protocol.BadSyntax()
	case int:
		if version == core.RESP {
			return []any{"+OK"}, nil
		}
		return []any{fmt.Sprintf("-BAD Version, only supports %v", core.RESP)}, nil
	}
}
