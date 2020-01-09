package template

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Guardian-Development/fastengine/client/fast/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	tokenxml "github.com/Guardian-Development/fastengine/internal/xml"
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
	Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error
}

func (template Template) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (*fix.Message, error) {
	fixMessage := fix.New()
	for _, unit := range template.TemplateUnits {
		err := unit.Deserialise(inputSource, pMap, &fixMessage)
		if err != nil {
			return nil, err
		}
	}

	return &fixMessage, nil
}

// New instance of the Store from the given FAST Templates XML file
func New(templateFile *os.File) (Store, error) {
	decoder := xml.NewDecoder(templateFile)
	xmlTags, err := tokenxml.LoadTagsFrom(decoder)

	if err != nil {
		return Store{}, err
	}

	if xmlTags.Type != templatesTag {
		return Store{}, fmt.Errorf("expected the root level of tag of the templateFile to be of type <templates> but was: %s", xmlTags.Type)
	}

	return loadStoreFromXML(xmlTags)
}
