package fix

import (
	"fmt"
	"strconv"
	"strings"
)

type Message struct {
	Tags        map[uint64]Value
	tagsInOrder []uint64
}

func (message *Message) SetTag(tag uint64, value Value) {
	message.Tags[tag] = value
	message.tagsInOrder = append(message.tagsInOrder, tag)
}

func (message Message) GetTag(tag uint64) (interface{}, error) {
	if value, ok := message.Tags[tag]; ok {
		switch t := value.(type) {
		case NullValue:
			return nil, nil
		case SequenceValue:
			return t.Get(), nil
		case RawValue:
			return t.Get(), nil
		default:
			return nil, fmt.Errorf("unsupported type of tag: %s", t)
		}
	}

	return nil, fmt.Errorf("no tag in message with value: %d", tag)
}

func (message Message) String() string {
	stringBuilder := strings.Builder{}
	for _, tag := range message.tagsInOrder {
		stringBuilder.WriteString(strconv.FormatUint(tag, 10))
		stringBuilder.WriteString("=")
		stringBuilder.WriteString(message.Tags[tag].String())
		stringBuilder.WriteString("|")
	}
	return stringBuilder.String()
}

func New() Message {
	message := Message{
		Tags:        make(map[uint64]Value),
		tagsInOrder: make([]uint64, 0),
	}

	return message
}

type Value interface {
	Get() interface{}
	String() string
}

type RawValue struct {
	value interface{}
}

func (rawValue RawValue) Get() interface{} {
	return rawValue.value
}

func (rawValue RawValue) String() string {
	return fmt.Sprint(rawValue.value)
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

func (nullValue NullValue) String() string {
	return "nil"
}

type SequenceValue struct {
	Values []Message
}

func (sequenceValue SequenceValue) Get() interface{} {
	return sequenceValue.Values
}

func (sequenceValue *SequenceValue) SetValue(index uint32, tag uint64, value Value) {
	sequenceValue.Values[index].SetTag(tag, value)
}

func (sequenceValue SequenceValue) String() string {
	stringBuilder := strings.Builder{}
	// write length of this tag
	stringBuilder.WriteString(string(len(sequenceValue.Values)))
	stringBuilder.WriteString("|")
	for _, value := range sequenceValue.Values {
		stringBuilder.WriteString(value.String())
	}
	return stringBuilder.String()
}

func NewSequenceValue(sequenceSize uint32) SequenceValue {
	value := SequenceValue{Values: make([]Message, sequenceSize)}
	for i := 0; i < len(value.Values); i++ {
		value.Values[i] = New()
	}

	return value
}
