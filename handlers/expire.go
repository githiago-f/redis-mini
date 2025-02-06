package handlers

import (
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/protocol"
)

func ExpireHandler(db *db.Datasource, args []*protocol.Arg) ([]*protocol.Value, error) {
	arsNumber := len(args)
	if arsNumber != 2 {
		return nil, protocol.BadArgNumber(arsNumber)
	}

	varName, err := args[0].AsID()
	if err != nil {
		return nil, err
	}

	expiresIn, err := args[1].AsNumber()
	if err != nil {
		return nil, err
	}

	db.SetTimeout(varName, int(expiresIn))

	return nil, nil
}
