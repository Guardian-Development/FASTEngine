package fix

import "fmt"

type Message struct {
	tags map[uint64]*interface{}
}

func (message *Message) SetTag(tag uint64, value interface{}) {
	message.tags[tag] = &value
}

func (message Message) GetTag(tag uint64) (interface{}, error) {
	if value, ok := message.tags[tag]; ok {
		return *value, nil
	}

	return nil, fmt.Errorf("No tag in message with value: %d", tag)
}

func New() Message {
	message := Message{
		tags: make(map[uint64]*interface{}),
	}

	return message
}
