package template

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

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

func (template Template) Deserialise(inputSource *bytes.Buffer) {
	//TODO: create FIX message, then go through each unit and deserialise into FIX message
}

// Unit represents an element within a FAST Template, with the ability to Serialise/Deserialise a part of a FAST message
type Unit interface {
	Deserialise(inputSource *bytes.Buffer)
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
