package fix

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/internal/fix"
)

type Message struct {
	tags map[uint64]fix.Value
}

func (message *Message) SetTag(tag uint64, value fix.Value) {
	message.tags[tag] = value
}

func (message Message) GetTag(tag uint64) (interface{}, error) {
	if value, ok := message.tags[tag]; ok {
		switch t := value.(type) {
		case fix.NullValue:
			return nil, nil
		case fix.RawValue:
			return t.Get(), nil
		default:
			return nil, fmt.Errorf("Unsupported type of tag: %s", t)
		}
	}

	return nil, fmt.Errorf("No tag in message with value: %d", tag)
}

func New() Message {
	message := Message{
		tags: make(map[uint64]fix.Value),
	}

	return message
}
