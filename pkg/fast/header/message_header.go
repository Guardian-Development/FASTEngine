package header

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// MessageHeader represents the beginning of every fast message. It contains the presence map for the message and the template id to use when decoding the message.
type MessageHeader struct {
	PMap       *presencemap.PresenceMap
	TemplateID uint32
}

// New MessageHeader read from the byte buffer
func New(message *bytes.Buffer, dict *dictionary.Dictionary, logger *log.Logger) (MessageHeader, error) {
	pMap, err := presencemap.New(message)
	if err != nil {
		logger.Printf("could not deserialise presence map from byte buffer, reason: %s", err)
		return MessageHeader{}, fmt.Errorf("unable to create presence map for message")
	}

	templateIDAttribute := fielduint32.NewCopyOperation(properties.New(0, "TemplateId", true, logger))
	templateID, err := templateIDAttribute.Deserialise(message, &pMap, dict)
	if err != nil {
		logger.Printf("could not deserialise template id from byte buffer, reason: %v", err)
		return MessageHeader{}, fmt.Errorf("could not deserialise template id from byte buffer")
	}

	switch t := templateID.(type) {
	case fix.RawValue:
		return MessageHeader{PMap: &pMap, TemplateID: t.Get().(uint32)}, nil
	}

	logger.Printf("no template id was found in the byte buffer, unable to calculate format of message, reason: %s", err)
	return MessageHeader{}, fmt.Errorf("message not supported: message must have template id encoded")
}
