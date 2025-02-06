package db

import (
	"bufio"
	"errors"
	"os"
	"time"

	"github.com/githiago-f/redis-mini/protocol"
)

func Persist(db *Datasource) error {
	fi, err := os.Create("snapshot.mrds")
	if err != nil {
		return err
	}

	defer fi.Close()

	db.values.Range(func(key any, value any) bool {
		fi.Write([]byte(key.(string)))
		fi.Write([]byte("\n"))
		fi.Write(value.(*protocol.Value).ToByteArray())
		fi.Write([]byte("\n"))
		return true
	})

	return nil
}

func Restore() (*Datasource, error) {
	cache := New()

	fi, openErr := os.Open("snapshot.gedis")
	if openErr != nil {
		return nil, openErr
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		key := scanner.Text()
		if !scanner.Scan() {
			return nil, errors.New("serialization error: incomplete snapshot")
		}
		value := scanner.Text()
		// TODO simplify serialization and deserialization of values
		cache.Set(key, protocol.NewValue(value))
	}

	if scanErr := scanner.Err(); scanErr != nil {
		return nil, scanErr
	}

	return cache, nil
}

func ScheduledSnapshot(db *Datasource) {
	for {
		time.Sleep(5 * time.Second)
		Persist(db)
	}
}
