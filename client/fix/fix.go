package fix

import (
	"fmt"
)

type Message struct {
	tags map[uint64]Value
}

type Value interface {
	Get() interface{}
}

type RawValue struct {
	value interface{}
}

type NullValue struct {
}

func (nullValue NullValue) Get() interface{} {
	return nil
}

func (rawValue RawValue) Get() interface{} {
	return rawValue.value
}

func NewRawValue(value interface{}) Value {
	if value == nil {
		return NullValue{}
	}

	return RawValue{value: value}
}

func (message *Message) SetTag(tag uint64, value Value) {
	message.tags[tag] = value
}

func (message Message) GetTag(tag uint64) (interface{}, error) {
	if value, ok := message.tags[tag]; ok {
		switch t := value.(type) {
		case NullValue:
			return nil, nil
		case RawValue:
			return t.Get(), nil
		default:
			return nil, fmt.Errorf("Unsupported type of tag: %s", t)
		}
	}

	return nil, fmt.Errorf("No tag in message with value: %d", tag)
}

func New() Message {
	message := Message{
		tags: make(map[uint64]Value),
	}

	return message
}
