package presencemap

import "bytes"

type PresenceMap struct {
}

func New(message *bytes.Buffer) (PresenceMap, error) {
	b, err := message.ReadByte()

	if err != nil {
		return PresenceMap{}, err
	}

	result := b & 128
	if result == 128 {
	}

	return PresenceMap{}, nil
}
