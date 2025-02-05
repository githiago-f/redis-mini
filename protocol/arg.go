package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	STR = "STR"
	NUM = "NUM"
	ID  = "ID"
)

type Arg struct {
	Type  string
	Value interface{}
}

func (a *Arg) AsID() (string, error) {
	if a.Type == ID {
		return a.Value.(string), nil
	}
	return "", fmt.Errorf("%v is not a valid identifier", a.Type)
}

func ParseArg(value string) *Arg {
	if strings.HasPrefix(value, "\"") {
		return &Arg{Type: STR, Value: value[1 : len(value)-1]}
	}

	if val, err := strconv.ParseFloat(value, 64); err == nil {
		return &Arg{Type: NUM, Value: val}
	}

	return &Arg{Type: ID, Value: value}
}
