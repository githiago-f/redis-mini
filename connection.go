package main

import (
	"bufio"
	"io"
	"net"
	"os"

	"github.com/githiago-f/redis-mini/broker"
	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/db"
	"github.com/githiago-f/redis-mini/handlers"
)

var eventBroker *broker.Broker

func init() {
	cache, err := db.Restore()
	if err != nil {
		core.Logger.Error(err)
		os.Exit(1)
	}

	eventBroker = broker.New(cache)

	eventBroker.Use("GET", handlers.GetHandler)
	eventBroker.Use("SET", handlers.SetHandler)
	eventBroker.Use("DEL", handlers.DelHandler)
	eventBroker.Use("MGET", handlers.MGetHandler)
	eventBroker.Use("KEYS", handlers.KeysHandler)
	eventBroker.Use("INCR", handlers.IncrHandler)
	eventBroker.Use("EXPIRE", handlers.ExpireHandler)
	eventBroker.Use("EXISTS", handlers.ExistsHandler)

	go db.ScheduledSnapshot(cache)
}

func HandleConnection(con net.Conn) {
	core.Logger.Infof("%v connection received", con.RemoteAddr().Network())
	defer con.Close()

	body, err := bufio.NewReader(con).ReadString('\n')
	if err != nil && err != io.EOF {
		core.Logger.Error(err)
		return
	}

	result, err := eventBroker.Handle(body)
	if err != nil {
		core.Logger.Error(err)
		con.Write([]byte(err.Error()))
	} else if result == nil {
		con.Write([]byte("nil"))
	} else {
		size := len(result)
		for i, val := range result {
			con.Write(val.ToByteArray())
			if i <= size-2 {
				con.Write([]byte("\r\n"))
			}
		}
	}

	con.Write([]byte("\r\n"))
}
