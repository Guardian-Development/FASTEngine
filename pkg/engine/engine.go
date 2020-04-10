package engine

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Guardian-Development/fastengine/pkg/fast/errors"

	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/header"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/loader"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

type FastEngine interface {
	Deserialise(message *bytes.Buffer) (*fix.Message, error)
}

type fastEngine struct {
	templateStore    store.Store
	globalDictionary dictionary.Dictionary
}

// Deserialise takes a FAST encoded FIX message in bytes, decodes and turns it into a FIX message
// Expected message format: (PMap (1+ bytes), templateId (1 + bytes), Message encoded from template with templateId)
func (engine fastEngine) Deserialise(message *bytes.Buffer) (*fix.Message, error) {
	engine.globalDictionary.Reset()

	messageHeader, err := header.New(message, &engine.globalDictionary)
	if err != nil {
		return nil, fmt.Errorf("unable to parse message, reason: %v", err)
	}

	if template, exists := engine.templateStore.Templates[messageHeader.TemplateID]; exists {
		return template.Deserialise(message, messageHeader.PMap, &engine.globalDictionary)
	}

	return nil, fmt.Errorf("%s: id %d", errors.D9, messageHeader.TemplateID)
}

// New instance of a FAST engine, that can serialise/deserialise FAST messages using the template store provided
func New(templateStore store.Store) FastEngine {
	return fastEngine{
		templateStore:    templateStore,
		globalDictionary: dictionary.New(),
	}
}

// NewFromTemplateFile of a FAST engine, that can serialise/deserialise FAST messages using the template file provided.
// This file should be xml, if we are unable to find the file or parse it, an error is returned
func NewFromTemplateFile(templateFile string) (FastEngine, error) {
	file, err := os.Open(templateFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open template file: %s", err)
	}
	templateStore, err := loader.Load(file)
	if err != nil {
		return nil, fmt.Errorf("unable to load template file: %s", err)
	}
	fastEngine := New(templateStore)
	return fastEngine, nil
}
