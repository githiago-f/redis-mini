package db

import (
	"fmt"
	"os"
	"time"
)

func Persist(db *Datasource) error {
	fi, err := os.Create("snapshot.mrds")
	if err != nil {
		return err
	}

	defer fi.Close()

	db.Values.Range(func(key any, value any) bool {
		fi.Write([]byte(fmt.Sprintf("+%v\r\n", key)))

		switch v := value.(type) {
		default:
		case int:
			fi.Write([]byte(fmt.Sprintf(":%v\r\n", v)))
		case float64:
			fi.Write([]byte(fmt.Sprintf(",%v\r\n", v)))
		case bool:
			fi.Write([]byte(fmt.Sprintf("#%v\r\n", v)))
		case string:
			fi.Write([]byte(fmt.Sprintf("$%d\r\n%v\r\n", len(v), v)))
		}

		return true
	})

	return nil
}

func Restore() (*Datasource, error) {
	cache := New()

	// fi, openErr := os.Open("snapshot.mrds")
	// if openErr != nil {
	// 	return nil, openErr
	// }
	// defer fi.Close()

	// scanner := bufio.NewScanner(fi)
	// for scanner.Scan() {
	// 	// key := scanner.Text()
	// 	if !scanner.Scan() {
	// 		return nil, errors.New("serialization error: incomplete snapshot")
	// 	}
	// 	// value := scanner.Text()
	// 	// TODO simplify serialization and deserialization of values
	// 	// cache.Set(key, protocol.NewValue(value))
	// }

	// if scanErr := scanner.Err(); scanErr != nil {
	// 	return nil, scanErr
	// }

	return cache, nil
}

func ScheduledSnapshot(db *Datasource) {
	for {
		time.Sleep(5 * time.Second)
		Persist(db)
	}
}
