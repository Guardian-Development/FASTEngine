package template

import (
	"encoding/xml"
	"fmt"
	"os"

	tokenxml "github.com/Guardian-Development/fastengine/internal/xml"
)

const templatesTag = "templates"

// Store represents a loaded set of Templates that can be used to Serialise/Deserialise FAST messages
type Store struct {
	Templates []Template
}

// Create an instance of the Store from the given FAST Templates XML file
func Create(templateFile *os.File) (Store, error) {
	decoder := xml.NewDecoder(templateFile)
	xmlTags, err := tokenxml.LoadTagsFrom(decoder)

	if err != nil {
		return Store{}, err
	}

	if xmlTags.Type != templatesTag {
		return Store{}, fmt.Errorf("expected the root level of tag of the templateFile to be of type <templates> but was: %s", xmlTags.Type)
	}

	templateStore := Store{
		Templates: make([]Template, len(xmlTags.NestedTags)),
	}

	for templateNumber, templateXMLElement := range xmlTags.NestedTags {
		template, err := createTemplate(&templateXMLElement)
		if err != nil {
			return Store{}, err
		}

		templateStore.Templates[templateNumber] = template
	}

	return templateStore, nil
}
