package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/handlers"
	"github.com/githiago-f/redis-mini/protocol"
)

var eventBroker *broker.Broker

func init() {
	cache, err := db.Restore()
	if err != nil {
		core.Logger.Error(err)
	}

	eventBroker = broker.New(cache)

	eventBroker.Use("GET", handlers.GetHandler)
	eventBroker.Use("HELLO", handlers.HELLOHandler)
	eventBroker.Use("SET", handlers.SetHandler)
	// eventBroker.Use("DEL", handlers.DelHandler)
	eventBroker.Use("MGET", handlers.MGetHandler)
	// eventBroker.Use("KEYS", handlers.KeysHandler)
	// eventBroker.Use("INCR", handlers.IncrHandler)
	// eventBroker.Use("EXPIRE", handlers.ExpireHandler)
	// eventBroker.Use("EXISTS", handlers.ExistsHandler)

	go db.ScheduledSnapshot(cache)
}

func HandleConnection(con net.Conn) {
	network := con.RemoteAddr().Network()
	if network != "tcp" {
		con.Write([]byte("invalid connection type"))
		con.Close()
	}

	defer con.Close()

	reader := bufio.NewReader(con)

	args, err := protocol.DecodeLine(reader)
	if err != nil {
		con.Write([]byte("-" + err.Error()))
		return
	}

	core.Logger.Debugf("Args: %v", args)

	switch argsList := args.(type) {
	default:
		con.Write([]byte("-" + protocol.BadArgNumber(1).Error()))
		return
	case []any:
		res, err := eventBroker.Handle(argsList[0].(string), argsList[1:])
		if err != nil {
			core.Logger.Error(err)
			con.Write([]byte("-" + err.Error()))
			return
		}

		for _, result := range res {
			con.Write([]byte(fmt.Sprintf("%v\r\n", result)))
		}
	}
}
