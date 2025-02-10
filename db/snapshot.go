package db

import (
	"bufio"
	"os"
	"sync"
	"time"

	"github.com/githiago-f/redis-mini/core"
	"github.com/githiago-f/redis-mini/protocol"
)

func Persist(db *Datasource) error {
	fi, err := os.Create("snapshot.mrds")
	if err != nil {
		return err
	}

	defer fi.Close()

	b, err := protocol.Encode(db.Values)
	if err != nil {
		return err
	}
	fi.Write(b)

	return nil
}

func Restore() (*Datasource, error) {
	cache := New()
	fi, err := os.Open("snapshot.mrds")

	if !os.IsNotExist(err) {
		reader := bufio.NewReader(fi)

		memory, err := protocol.DecodeLine(reader)
		if err != nil {
			core.Logger.Infof("Bad memory serialization :: %v", err)
			os.Exit(1)
		}

		switch mem := memory.(type) {
		default:
			core.Logger.Infof("Bad memory allocation :: %v", mem)
		case *sync.Map:
			cache.Values = mem
		}
	}

	core.Logger.Infof("Restored memory %v", cache.Values)

	return cache, nil
}

func ScheduledSnapshot(db *Datasource) {
	for {
		time.Sleep(5 * time.Second)
		Persist(db)
	}
}
