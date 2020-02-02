package store

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

// Store represents a loaded set of Templates that can be used to Serialise/Deserialise FAST messages
type Store struct {
	Templates map[uint32]Template
}

// Template represents an ordered List of operations needed to Serialise/Deserialise a FAST message
type Template struct {
	TemplateUnits []Unit
}

// Unit represents an element within a FAST Template, with the ability to Serialise/Deserialise a part of a FAST message
type Unit interface {
	Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error)
	GetTagId() uint64
	RequiresPmap() bool
}

func (template Template) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (*fix.Message, error) {
	fixMessage := fix.New()
	for _, unit := range template.TemplateUnits {
		value, err := unit.Deserialise(inputSource, pMap)
		if err != nil {
			return nil, err
		}
		tagID := unit.GetTagId()
		fixMessage.SetTag(tagID, value)
	}

	return &fixMessage, nil
}
