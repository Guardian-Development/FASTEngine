package engine

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/Guardian-Development/fastengine/pkg/fast/errors"

	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/header"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/loader"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// FastEngine capable of deserialising a fast encoded message from the given byte buffer.
// This is not thread safe, and should only be called from a single threaded context, due to the fast engine making
// use of a dictionary of previous values
type FastEngine interface {
	Deserialise(message *bytes.Buffer) (*fix.Message, error)
}

type fastEngine struct {
	templateStore    store.Store
	globalDictionary dictionary.Dictionary

	logger *log.Logger
}

// Deserialise takes a FAST encoded FIX message in bytes, decodes and turns it into a FIX message
// Expected message format: (PMap (1+ bytes), templateId (1 + bytes), Message encoded from template with templateId)
func (engine fastEngine) Deserialise(message *bytes.Buffer) (*fix.Message, error) {
	engine.globalDictionary.Reset()

	messageHeader, err := header.New(message, &engine.globalDictionary, engine.logger)
	if err != nil {
		engine.logger.Printf("unable to deserialise header of message: %v", err)
		return nil, fmt.Errorf("unable to parse message, reason: %v", err)
	}

	if template, exists := engine.templateStore.Templates[messageHeader.TemplateID]; exists {
		return template.Deserialise(message, messageHeader.PMap, &engine.globalDictionary)
	}

	engine.logger.Println("no template exists for id", messageHeader.TemplateID)
	return nil, fmt.Errorf("%s: id %d", errors.D9, messageHeader.TemplateID)
}

// New instance of a FAST engine, that can serialise/deserialise FAST messages using the template store provided
func New(templateStore store.Store, logger *log.Logger) FastEngine {
	return fastEngine{
		templateStore:    templateStore,
		globalDictionary: dictionary.New(),
		logger:           logger,
	}
}

// NewFromTemplateFile of a FAST engine, that can serialise/deserialise FAST messages using the template file provided.
// This file should be xml, if we are unable to find the file or parse it, an error is returned
func NewFromTemplateFile(templateFile string, logger *log.Logger) (FastEngine, error) {
	file, err := os.Open(templateFile)
	defer file.Close()

	if err != nil {
		logger.Println("unable to open template file")
		return nil, fmt.Errorf("unable to open template file: %s", err)
	}
	templateStore, err := loader.Load(file, logger)
	if err != nil {
		logger.Println("unable to load template store")
		return nil, fmt.Errorf("unable to load template file: %s", err)
	}
	fastEngine := New(templateStore, logger)
	return fastEngine, nil
}
