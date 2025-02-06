package protocol

import "fmt"

type Value struct {
	Value interface{}
}

func (v *Value) ToString() string {
	if v == nil || v.Value == nil {
		return "nil"
	}
	switch val := v.Value.(type) {
	default:
		return "nil"
	case int, float64, bool:
		return fmt.Sprintf(":%v", val)
	case string:
		return "$" + val
	}
}

func (v *Value) ToByteArray() []byte {
	return []byte(v.ToString())
}

func (v *Value) Collect() []*Value {
	return []*Value{v}
}

func NewValue(val interface{}) *Value {
	return &Value{
		Value: val,
	}
}

func OkValue() *Value {
	return &Value{
		Value: "OK",
	}
}
