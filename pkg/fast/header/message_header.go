package header

import (
	"bytes"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

type MessageHeader struct {
	PMap       *presencemap.PresenceMap
	TemplateID uint32
}

func New(message *bytes.Buffer, dict *dictionary.Dictionary) (MessageHeader, error) {
	pMap, err := presencemap.New(message)
	if err != nil {
		return MessageHeader{}, fmt.Errorf("unable to create pMap for message, reason: %v", err)
	}

	templateIdAttribute := fielduint32.NewCopyOperation(properties.New(0, "TemplateId", true))
	templateId, err := templateIdAttribute.Deserialise(message, &pMap, dict)
	if err != nil {
		return MessageHeader{}, err
	}

	switch t := templateId.(type) {
	case fix.RawValue:
		return MessageHeader{PMap: &pMap, TemplateID: t.Get().(uint32)}, nil
	}

	return MessageHeader{}, fmt.Errorf("message not supported: message must have template ID encoded")
}
