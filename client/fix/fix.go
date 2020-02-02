package fix

import (
	"fmt"
)

type Message struct {
	tags map[uint64]Value
}

func (message *Message) SetTag(tag uint64, value Value) {
	message.tags[tag] = value
}

func (message Message) GetTag(tag uint64) (interface{}, error) {
	if value, ok := message.tags[tag]; ok {
		switch t := value.(type) {
		case NullValue:
			return nil, nil
		case SequenceValue:
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

type Value interface {
	Get() interface{}
}

type RawValue struct {
	value interface{}
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

type NullValue struct {
}

func (nullValue NullValue) Get() interface{} {
	return nil
}

type SequenceValue struct {
	values []Message
}

func (sequenceValue SequenceValue) Get() interface{} {
	return sequenceValue.values
}

func (sequenceValue *SequenceValue) SetValue(index uint32, tag uint64, value Value) {
	sequenceValue.values[index].SetTag(tag, value)
}

func NewSequenceValue(sequenceSize uint32) SequenceValue {
	value := SequenceValue{values: make([]Message, sequenceSize)}
	for i := 0; i < len(value.values); i++ {
		value.values[i] = New()
	}

	return value
}
