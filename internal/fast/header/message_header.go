package header

import (
	"bytes"
	"fmt"

	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

type MessageHeader struct {
	PMap       *presencemap.PresenceMap
	TemplateID uint32
}

func New(message *bytes.Buffer) (MessageHeader, error) {
	pMap, err := presencemap.New(message)
	if err != nil {
		return MessageHeader{}, fmt.Errorf("unable to create pMap for message, reason: %v", err)
	}

	if pMap.GetIsSetAndIncrement() {
		templateID, err := fast.ReadUInt32(message)
		if err != nil {
			return MessageHeader{}, err
		}

		return MessageHeader{PMap: &pMap, TemplateID: templateID.Value}, nil
	}

	return MessageHeader{}, fmt.Errorf("message not supported: message must have template ID encoded")
}
