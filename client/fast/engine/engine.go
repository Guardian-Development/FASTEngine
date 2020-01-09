package engine

import (
	"bytes"
	"fmt"

	"github.com/Guardian-Development/fastengine/client/fast/fix"
	"github.com/Guardian-Development/fastengine/client/fast/template"
	"github.com/Guardian-Development/fastengine/internal/fast/header"
)

type FastEngine interface {
	Deserialise(message *bytes.Buffer) (*fix.Message, error)
}

type fastEngine struct {
	templateStore template.Store
}

// Deserialise takes a FAST encoded FIX message in bytes, decodes and turns it into a FIX message
// Expected message format: (PMap (1+ bytes), templateId (1 + bytes), Message encoded from template with templateId)
func (engine fastEngine) Deserialise(message *bytes.Buffer) (*fix.Message, error) {
	messageHeader, err := header.New(message)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse message, reason: %v", err)
	}

	if template, exists := engine.templateStore.Templates[messageHeader.TemplateID]; exists {
		return template.Deserialise(message, messageHeader.PMap)
	}

	return nil, fmt.Errorf("No template found in store to deserialise message with ID: %d", messageHeader.TemplateID)
}

// New instance of a FAST engine, that can serialise/deserialise FAST messages using the templates provided
func New(templateStore template.Store) FastEngine {
	return fastEngine{
		templateStore: templateStore,
	}
}
