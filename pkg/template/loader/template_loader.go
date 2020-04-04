package loader

import (
	"encoding/xml"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldsequence"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loadasciistring"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loadbytevector"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loaddecimal"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loadint32"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loadint64"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loadproperties"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loaduint32"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loaduint64"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/loadunicodestring"
	"github.com/Guardian-Development/fastengine/pkg/template/structure"
	"os"
	"strconv"

	tokenxml "github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/template/store"
)

// Load instance of the Store from the given FAST Templates XML file
func Load(templateFile *os.File) (store.Store, error) {
	decoder := xml.NewDecoder(templateFile)
	xmlTags, err := tokenxml.LoadTagsFrom(decoder)

	if err != nil {
		return store.Store{}, err
	}

	if xmlTags.Type != structure.TemplatesTag {
		return store.Store{}, fmt.Errorf("expected the root level of tag of the templateFile to be of type <templates> but was: %s", xmlTags.Type)
	}

	return loadStoreFromXML(xmlTags)
}

func loadStoreFromXML(xmlTags tokenxml.Tag) (store.Store, error) {
	templateStore := store.Store{
		Templates: make(map[uint32]store.Template),
	}

	for _, templateXMLElement := range xmlTags.NestedTags {
		template, err := createTemplate(&templateXMLElement)
		if err != nil {
			return store.Store{}, err
		}
		templateID, err := strconv.ParseUint(templateXMLElement.Attributes["id"], 10, 32)

		if err != nil {
			return store.Store{}, fmt.Errorf("could not parse template ID, make sure it is present and uint: %v", err)
		}
		if _, exists := templateStore.Templates[uint32(templateID)]; exists {
			return store.Store{}, fmt.Errorf("template with ID %d, has already been loaded", templateID)
		}

		templateStore.Templates[uint32(templateID)] = template
	}

	return templateStore, nil
}

func createTemplate(templateRoot *tokenxml.Tag) (store.Template, error) {
	if templateRoot.Type != structure.TemplateTag {
		return store.Template{}, fmt.Errorf("expected to find template tag, but found %s", templateRoot.Type)
	}

	template := store.Template{
		TemplateUnits: make([]store.Unit, len(templateRoot.NestedTags)),
	}

	for unitNumber, tagInTemplate := range templateRoot.NestedTags {
		templateUnit, err := createTemplateUnit(&tagInTemplate)

		if err != nil {
			return store.Template{}, err
		}

		template.TemplateUnits[unitNumber] = templateUnit
	}

	return template, nil
}

func createTemplateUnit(tagInTemplate *tokenxml.Tag) (store.Unit, error) {
	fieldDetails, err := loadproperties.Load(tagInTemplate)
	if err != nil {
		return nil, err
	}

	switch tagInTemplate.Type {
	case structure.StringTag:
		if tagInTemplate.Attributes["charset"] == structure.UnicodeStringLabel {
			return loadunicodestring.Load(tagInTemplate, fieldDetails)
		}
		return loadasciistring.Load(tagInTemplate, fieldDetails)
	case structure.UInt32Tag, structure.LengthTag:
		return loaduint32.Load(tagInTemplate, fieldDetails)
	case structure.Int32Tag:
		return loadint32.Load(tagInTemplate, fieldDetails)
	case structure.UInt64Tag:
		return loaduint64.Load(tagInTemplate, fieldDetails)
	case structure.Int64Tag:
		return loadint64.Load(tagInTemplate, fieldDetails)
	case structure.DecimalTag:
		return loaddecimal.Load(tagInTemplate, fieldDetails)
	case structure.ByteVectorTag:
		return loadbytevector.Load(tagInTemplate, fieldDetails)
	case structure.SequenceTag:
		return loadSequence(tagInTemplate, fieldDetails)
	default:
		return nil, fmt.Errorf("unsupported tag type: %s", tagInTemplate.Type)
	}
}

func loadSequence(tagInTemplate *tokenxml.Tag, fieldDetails properties.Properties) (fieldsequence.FieldSequence, error) {
	fields := make([]store.Unit, 0)
	for _, tagInTemplate := range tagInTemplate.NestedTags {
		if tagInTemplate.Type == structure.LengthTag {
			continue
		}
		templateUnit, err := createTemplateUnit(&tagInTemplate)
		if err != nil {
			return fieldsequence.FieldSequence{}, err
		}

		fields = append(fields, templateUnit)
	}

	if tagInTemplate.NestedTags[0].Type == structure.LengthTag {
		lengthProperties, err := loadproperties.Load(&tagInTemplate.NestedTags[0])
		if err != nil {
			return fieldsequence.FieldSequence{}, err
		}
		length, err := loaduint32.Load(&tagInTemplate.NestedTags[0], lengthProperties)
		if err != nil {
			return fieldsequence.FieldSequence{}, err
		}
		if structure.IsNullString(length.FieldDetails.Name) {
			length.FieldDetails.Name = fieldDetails.Name
		}
		length.FieldDetails.Required = fieldDetails.Required
		return fieldsequence.New(fieldDetails, length, fields), nil
	} else {
		length := fielduint32.New(properties.New(0, fieldDetails.Name, fieldDetails.Required))
		return fieldsequence.New(fieldDetails, length, fields), nil
	}
}
