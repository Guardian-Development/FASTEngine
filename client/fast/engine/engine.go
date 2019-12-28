package engine

import (
	"bytes"
	"fmt"

	"github.com/Guardian-Development/fastengine/client/fast/template"
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
	_, err := presencemap.New(message)
	if err != nil {
		return FixMessage{}, fmt.Errorf("Unable to create pMap for message, reason: %v", err)
	}
	return FixMessage{}, nil
}

// New instance of a FAST engine, that can seralise/deserialise FAST messages using the templates provided
func New(templateStore template.Store) FastEngine {
	return fastEngine{
		templateStore: templateStore,
	}
}
