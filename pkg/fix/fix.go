package fix

import (
	"fmt"
	"strconv"
	"strings"
)

// Message is a fix tag, value message
type Message struct {
	Tags        map[uint64]Value
	tagsInOrder []uint64
}

// SetTag with value
func (message *Message) SetTag(tag uint64, value Value) {
	message.Tags[tag] = value
	message.tagsInOrder = append(message.tagsInOrder, tag)
}

// GetTag returns the value associated with the tag.
// This can be: nil, int32, uint32, int64, uint64, []byte, string, []Message (for sequences)
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

// String representation of a fix message
func (message Message) String() string {
	stringBuilder := strings.Builder{}
	for _, tag := range message.tagsInOrder {
		stringBuilder.WriteString(strconv.FormatUint(tag, 10))
		stringBuilder.WriteString("=")
		stringBuilder.WriteString(message.Tags[tag].String())
	}
	return stringBuilder.String()
}

// New empty fix message
func New() Message {
	message := Message{
		Tags:        make(map[uint64]Value),
		tagsInOrder: make([]uint64, 0),
	}

	return message
}

// Value of a tag within a fix message
type Value interface {
	Get() interface{}
	String() string
}

// RawValue is a raw go type that can be used
type RawValue struct {
	value interface{}
}

// Get raw value
func (rawValue RawValue) Get() interface{} {
	return rawValue.value
}

// String is a value with a pipe seperator
func (rawValue RawValue) String() string {
	return fmt.Sprintf("%v|", rawValue.value)
}

// NewRawValue of the given type, if null we return a null fix representation explicitly
func NewRawValue(value interface{}) Value {
	if value == nil {
		return NullValue{}
	}

	return RawValue{value: value}
}

// NullValue is an explicit null value
type NullValue struct {
}

// Get nil value
func (nullValue NullValue) Get() interface{} {
	return nil
}

// String nil value with pipe seperator
func (nullValue NullValue) String() string {
	return "nil|"
}

// SequenceValue represents a repeating group
type SequenceValue struct {
	Values []Message
}

// Get returns an array of repeating groups, which are of the overall fix message type
func (sequenceValue SequenceValue) Get() interface{} {
	return sequenceValue.Values
}

// SetValue of tag within the given repeating group index
func (sequenceValue *SequenceValue) SetValue(index uint32, tag uint64, value Value) {
	sequenceValue.Values[index].SetTag(tag, value)
}

// String representation of the repeating group
func (sequenceValue SequenceValue) String() string {
	stringBuilder := strings.Builder{}
	// write length of this tag
	stringBuilder.WriteString(fmt.Sprint(len(sequenceValue.Values)))
	stringBuilder.WriteString("|")
	for _, value := range sequenceValue.Values {
		stringBuilder.WriteString(value.String())
	}
	return stringBuilder.String()
}

// NewSequenceValue with the given number of repeating groups
func NewSequenceValue(sequenceSize uint32) SequenceValue {
	value := SequenceValue{Values: make([]Message, sequenceSize)}
	for i := 0; i < len(value.Values); i++ {
		value.Values[i] = New()
	}

	return value
}
