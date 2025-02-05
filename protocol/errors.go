package protocol

import (
	"errors"
	"fmt"
)

func BadCommand(command string) error {
	return fmt.Errorf("ERR unknown command '%v'", command)
}

func BadType() error {
	return errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
}
