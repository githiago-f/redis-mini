package protocol

import (
	"bufio"
	"bytes"
	"strconv"

	"github.com/githiago-f/redis-mini/core"
)

var DELIM = []byte("\r")

func parseInt(msg []byte, reader *bufio.Reader) (int, error) {
	// skip \n for next reader
	reader.ReadByte()

	skip := 0
	if msg[0] == '-' || msg[0] == '+' {
		skip = 1
	}
	num, _, _ := bytes.Cut(msg[skip:], DELIM)

	i, err := strconv.Atoi(string(num))
	if err != nil {
		return 0, err
	}

	if msg[0] == '-' {
		return -i, nil
	}
	return i, nil
}

func parseStr(msg []byte, reader *bufio.Reader) (string, error) {
	size, err := parseInt(msg[1:], reader)
	if err != nil {
		return "", err
	}

	str := make([]byte, size)
	i, err := reader.Read(str)
	// skip \r\n
	reader.Read(make([]byte, 2))

	core.Logger.Infof("Readed %v bytes", i)

	return string(str), err
}

func parseArray(msg []byte, reader *bufio.Reader) ([]any, error) {
	numberOfLines, err := parseInt(msg[1:], reader)
	if err != nil {
		core.Logger.Error(err)
		return nil, BadType()
	}
	items := make([]any, 0)

	for i := 0; i < numberOfLines; i++ {
		line, err := reader.ReadBytes(DELIM[0])
		if err != nil {
			core.Logger.Error(err)
			return nil, err
		}
		item, err := DecodeAtom(line, reader)
		if err != nil {
			core.Logger.Error(err)
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}

func parseHash(msg []byte, reader *bufio.Reader) (map[string]any, error) {
	numberOfEntries, err := parseInt(msg[1:], reader)
	if err != nil {
		core.Logger.Error(err)
		return nil, BadType()
	}

	entries := map[string]any{}
	for i := 0; i < numberOfEntries; i++ {
		key, err := DecodeLine(reader)
		if err != nil {
			return nil, err
		}
		val, err := DecodeLine(reader)
		if err != nil {
			return nil, err
		}

		entries[key.(string)] = val
	}

	return entries, nil
}

func parseDouble(msg []byte, reader *bufio.Reader) (any, error) {
	reader.ReadByte()

	skip := 0
	if msg[0] == '-' || msg[0] == '+' {
		skip = 1
	}

	num, _, _ := bytes.Cut(msg[skip:], DELIM)

	f, err := strconv.ParseFloat(string(num), 64)
	if err != nil {
		return nil, err
	}

	if msg[0] == '-' {
		return -f, nil
	}

	return f, nil
}

/*
Decode and parse only one atom of the reader.
(e.g.: A simple string will be decoded from +simpleString to "simpleString")
*/
func DecodeAtom(msg []byte, reader *bufio.Reader) (any, error) {
	leadByte := msg[0]

	switch leadByte {
	default:
		return nil, nil
	case ':':
		return parseInt(msg[1:], reader)
	case '+':
		b, _, _ := bytes.Cut(msg[1:], DELIM)
		reader.ReadByte()
		return string(b), nil
	case '$':
		return parseStr(msg, reader)
	case '*':
		return parseArray(msg, reader)
	case '#':
		reader.ReadByte()
		return msg[1] == 't', nil
	case ',':
		return parseDouble(msg[1:], reader)
	case '%':
		return parseHash(msg, reader)
	case '_':
		return nil, nil
	}
}

func DecodeLine(reader *bufio.Reader) (any, error) {
	line, err := reader.ReadBytes(DELIM[0])
	if err != nil {
		core.Logger.Error(err)
		return nil, err
	}
	key, err := DecodeAtom(line, reader)
	if err != nil {
		core.Logger.Error(err)
		return nil, err
	}
	return key, nil
}
