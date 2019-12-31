package engine

import (
	"bytes"
	"fmt"

	"github.com/Guardian-Development/fastengine/client/fast/template"
	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

type FixMessage struct {
}

type FastEngine interface {
	Deserialise(message *bytes.Buffer) (FixMessage, error)
}

type fastEngine struct {
	templateStore template.Store
}

// Deserialise takes a FAST encoded FIX message in bytes, decodes and turns it into a FIX message
// Expected message format: (PMap (1+ bytes), templateId (1 + bytes), Message encoded from template with templateId)
func (engine fastEngine) Deserialise(message *bytes.Buffer) (FixMessage, error) {
	pMap, err := presencemap.New(message)
	if err != nil {
		return FixMessage{}, fmt.Errorf("Unable to create pMap for message, reason: %v", err)
	}

	if pMap.GetIsSetAndIncrement() {
		templateID, err := fast.ReadUInt32(message)
		if err != nil {
			return FixMessage{}, err
		}

		if template, exists := engine.templateStore.Templates[templateID]; exists {
			template.Deserialise(message)
			return FixMessage{}, nil
		}

		return FixMessage{}, fmt.Errorf("No template found in store to deserialise message with ID: %d", templateID)
	}

	return FixMessage{}, fmt.Errorf("Message not supported: message must have template ID encoded")
}

// New instance of a FAST engine, that can serialise/deserialise FAST messages using the templates provided
func New(templateStore template.Store) FastEngine {
	return fastEngine{
		templateStore: templateStore,
	}
}
