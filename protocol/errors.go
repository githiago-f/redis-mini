package protocol

import (
	"errors"
	"fmt"
)

func BadCommand(command string) error {
	return fmt.Errorf("ERR unknown command '%v'", command)
}

func BadArgNumber(argsNumber int) error {
	return fmt.Errorf("ERR invalid number of args %v expected 2", argsNumber)
}

func BadType() error {
	return errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
}
