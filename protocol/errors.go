package protocol

import (
	"errors"
	"fmt"
)

func BadCommand(command string) error {
	return fmt.Errorf("ERR unknown command '%v'", command)
}

func BadArgNumber(argsNumber int) error {
	return fmt.Errorf("ERR invalid number of args, expected %v", argsNumber)
}

func BadType() error {
	return errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func BadSyntax() error {
	return errors.New("SYNTAX invalid syntax")
}
