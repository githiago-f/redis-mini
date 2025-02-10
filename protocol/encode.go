package protocol

import (
	"bytes"
	"fmt"
	"sync"
)

func Encode(obj any) ([]byte, error) {
	switch v := obj.(type) {
	default:
		return nil, BadType()
	case int:
		return []byte(fmt.Sprintf(":%v\r\n", v)), nil
	case float64:
		return []byte(fmt.Sprintf(",%v\r\n", v)), nil
	case string:
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)), nil
	case error:
		return []byte(fmt.Sprintf("-%s\r\n", v.Error())), nil
	case []any:
		size := len(v)
		buffer := bytes.NewBuffer([]byte(fmt.Sprintf("*%d", size)))

		for _, i := range v {
			otherVal, err := Encode(i)
			if err != nil {
				return nil, err
			}
			buffer.Write(otherVal)
		}

		return buffer.Bytes(), nil
	case *sync.Map:
		size := 0
		buffer := bytes.NewBuffer([]byte{})

		v.Range(func(key, value any) bool {
			size++

			k := fmt.Sprintf("+%v\r\n", key)
			buffer.Write([]byte(k))

			val, err := Encode(value)
			if err != nil {
				return false
			}
			buffer.Write(val)

			return true
		})

		resultBuff := bytes.NewBuffer([]byte(fmt.Sprintf("%%%d\r\n", size)))
		resultBuff.Write(buffer.Bytes())

		return resultBuff.Bytes(), nil
	}
}
