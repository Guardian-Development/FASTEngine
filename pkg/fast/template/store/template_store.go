package store

import (
	"bytes"
	"log"

	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// Store represents a loaded set of Templates that can be used to Serialise/Deserialise FAST messages
type Store struct {
	Templates map[uint32]Template
	logger    *log.Logger
}

// Template represents an ordered List of operations needed to Serialise/Deserialise a FAST message
type Template struct {
	TemplateUnits []Unit
}

// Unit represents an element within a FAST Template, with the ability to Serialise/Deserialise a part of a FAST message
type Unit interface {
	Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error)
	GetTagId() uint64
	RequiresPmap() bool
}

func (template Template) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (*fix.Message, error) {
	fixMessage := fix.New()
	for _, unit := range template.TemplateUnits {
		value, err := unit.Deserialise(inputSource, pMap, dictionary)
		if err != nil {
			return &fixMessage, err
		}
		tagID := unit.GetTagId()
		fixMessage.SetTag(tagID, value)
	}

	return &fixMessage, nil
}
