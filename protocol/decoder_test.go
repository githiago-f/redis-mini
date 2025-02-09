package protocol_test

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/githiago-f/redis-mini/protocol"
)

func makeBuffer(message []byte) ([]byte, *bufio.Reader) {
	reader := bufio.NewReader(bytes.NewReader(message))
	slice, err := reader.ReadBytes('\r')
	if err != nil && err != io.EOF {
		return nil, nil
	}

	return slice, reader
}

func TestDecodingInteger(t *testing.T) {
	msg, reader := makeBuffer([]byte(":-150\r\n"))
	want := -150

	res, err := protocol.DecodeAtom(msg, reader)

	if want != res.(int) || err != nil {
		t.Fatalf("Decode(':-150\\r\\n') = %d, %v; want match for %d, nil", res.(int), err, want)
	}

	msg, reader = makeBuffer([]byte(":150\r\n"))
	want = 150

	res, err = protocol.DecodeAtom(msg, reader)

	if want != res.(int) || err != nil {
		t.Fatalf("Decode(':-150\\r\\n') = %d, %v; want match for %d, nil", res.(int), err, want)
	}

	checkLastByte(t, reader)
}

func TestDecodingSimpleString(t *testing.T) {
	input, reader := makeBuffer([]byte("+HELLO\r\n"))
	want := "HELLO"

	res, err := protocol.DecodeAtom(input, reader)

	if want != res.(string) || err != nil {
		t.Fatalf("Decode('+HELLO\\r\\n') = %s, %v; want match for %s, nil", res.(string), err, want)
	}

	checkLastByte(t, reader)
}

func TestDecodingBulkString(t *testing.T) {
	msg, reader := makeBuffer([]byte("$10\r\nmy message\r\n"))
	want := "my message"

	res, err := protocol.DecodeAtom(msg, reader)

	if want != res.(string) || err != nil {
		t.Fatalf("Decode('$10\\r\\nmy message\\r\\n') = %s, %v; want match for %s, nil", res.(string), err, want)
	}

	msg, reader = makeBuffer([]byte("$0\r\n"))
	want = ""

	res, err = protocol.DecodeAtom(msg, reader)
	if want != res.(string) || err != nil {
		t.Fatalf("Decode('$0\\r\\n') = %s, %v; want match for %s, nil", res.(string), err, want)
	}

	checkLastByte(t, reader)
}

func TestDecodingBool(t *testing.T) {
	msg, reader := makeBuffer([]byte("#t\r\n"))
	want := true

	res, err := protocol.DecodeAtom(msg, reader)

	if want != res.(bool) || err != nil {
		t.Fatalf("Decode('#t\\r\\n') = %v, %v; want match for %v, nil", res.(bool), err, want)
	}

	msg, reader = makeBuffer([]byte("#f\r\n"))
	want = false

	res, err = protocol.DecodeAtom(msg, reader)

	if want != res.(bool) || err != nil {
		t.Fatalf("Decode('#f\\r\\n') = %v, %v; want match for %v, nil", res.(bool), err, want)
	}

	checkLastByte(t, reader)
}

func TestDecodingDouble(t *testing.T) {
	msg, reader := makeBuffer([]byte(",13.23\r\n"))
	want := 13.23

	res, err := protocol.DecodeAtom(msg, reader)

	if want != res.(float64) || err != nil {
		t.Fatalf("Decode(',13.23f\\r\\n') = %v, %v; want match for %v, nil", res.(float64), err, want)
	}

	checkLastByte(t, reader)

	msg, reader = makeBuffer([]byte(",0.00003\r\n"))
	want = 0.00003

	res, err = protocol.DecodeAtom(msg, reader)
	if want != res.(float64) || err != nil {
		t.Fatalf("Decode(',0.00003\\r\\n') = %v, %v; want match for %v, nil", res.(float64), err, want)
	}

	checkLastByte(t, reader)
}

func TestDecodingArray(t *testing.T) {
	msg, reader := makeBuffer([]byte("*2\r\n$3\r\nstr\r\n:3\r\n"))
	want1 := "str"
	want2 := 3

	res, err := protocol.DecodeAtom(msg, reader)

	if err != nil || res.([]any)[0].(string) != want1 || res.([]any)[1].(int) != want2 {
		t.Fatalf(
			"Decode('*2\\r\\n$3\\r\\nstr\\r\\n:3\\r\\n') = %v, %v; want match for %v, nil",
			res.([]any),
			err,
			[]any{want1, want2},
		)
	}

	checkLastByte(t, reader)
}

func TestDecondingHash(t *testing.T) {
	msg, reader := makeBuffer([]byte("%2\r\n+myKey\r\n:123\r\n+otherKey\r\n#f\r\n"))
	want := map[string]any{"myKey": 123, "otherKey": false}

	res, err := protocol.DecodeAtom(msg, reader)

	decodeCall := "Decode(%%2\\r\\n+myKey\\r\\n:123\\r\\n+otherKey\\r\\n#f\\r\\n)"

	if err != nil {
		t.Fatalf(
			"%s = nil, %v; wanted match for %v, nil",
			decodeCall,
			err,
			res,
		)
	}

	resMap := res.(map[string]any)
	if resMap["myKey"] == nil || resMap["myKey"] != want["myKey"] {
		t.Fatalf(
			"%s = {myKey: %v}, nil; wanted match for {myKey: %v}, nil",
			decodeCall,
			resMap["myKey"],
			want["myKey"],
		)
	}

	if resMap["otherKey"] == nil || resMap["otherKey"] != want["otherKey"] {
		t.Fatalf(
			"%s = {otherKey: %v}, nil; wanted match for {otherKey: %v}, nil",
			decodeCall,
			resMap["otherKey"],
			want["otherKey"],
		)
	}
}

func checkLastByte(t *testing.T, reader *bufio.Reader) {
	readed, _ := reader.Peek(1)
	if len(readed) != 0 {
		t.Fatalf("Decode has left %b byte(s) to read", readed)
	}
}
