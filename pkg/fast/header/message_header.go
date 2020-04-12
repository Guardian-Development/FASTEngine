package header

import (
	"bytes"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
	"log"
)

type MessageHeader struct {
	PMap       *presencemap.PresenceMap
	TemplateID uint32
}

func New(message *bytes.Buffer, dict *dictionary.Dictionary, logger *log.Logger) (MessageHeader, error) {
	pMap, err := presencemap.New(message)
	if err != nil {
		logger.Printf("could not deserialise presence map from byte buffer, reason: %s", err)
		return MessageHeader{}, fmt.Errorf("unable to create presence map for message")
	}

	templateIdAttribute := fielduint32.NewCopyOperation(properties.New(0, "TemplateId", true, logger))
	templateId, err := templateIdAttribute.Deserialise(message, &pMap, dict)
	if err != nil {
		logger.Printf("could not deserialise template id from byte buffer, reason: %v", err)
		return MessageHeader{}, fmt.Errorf("could not deserialise template id from byte buffer")
	}

	switch t := templateId.(type) {
	case fix.RawValue:
		return MessageHeader{PMap: &pMap, TemplateID: t.Get().(uint32)}, nil
	}

	logger.Printf("no template id was found in the byte buffer, unable to calculate format of message, reason: %s", err)
	return MessageHeader{}, fmt.Errorf("message not supported: message must have template id encoded")
}
