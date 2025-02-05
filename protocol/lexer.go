package protocol

import (
	"strings"

	"github.com/githiago-f/redis-mini/core"
)

func Lex(str string) (string, []*Arg) {
	t := []rune(strings.TrimSpace(str))
	core.Logger.Debugf("Lexing %v", t)

	parts := GetParts(t)
	argsLen := len(parts)

	core.Logger.Infof("Parts %v", parts)
	core.Logger.Infof("Args len: %v", argsLen-1)

	var args []*Arg = make([]*Arg, argsLen-1)
	for i := 1; i < argsLen; i++ {
		args[i-1] = ParseArg(parts[i])
	}

	return parts[0], args
}

func GetParts(t []rune) []string {
	n := len(t)
	core.Logger.Debugf("Tokens size %v", n)

	isClosed := true
	parts := make([]string, 0, n)
	builder := strings.Builder{}
	for _, char := range t {
		if isClosed && char == 32 {
			part := builder.String()
			core.Logger.Debugf("Appending part %v", part)
			parts = append(parts, part)

			if isClosed {
				builder.Reset()
			}

			continue
		}
		if char == 34 {
			isClosed = !isClosed
		}
		builder.WriteRune(char)
	}

	part := builder.String()
	core.Logger.Debugf("Appending part %v", part)
	parts = append(parts, part)

	return parts
}
