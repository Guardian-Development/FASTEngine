package loader

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
	"github.com/Guardian-Development/fastengine/internal/converter"
	"github.com/Guardian-Development/fastengine/internal/fast/field"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	tokenxml "github.com/Guardian-Development/fastengine/internal/xml"
)

const templatesTag = "templates"
const templateTag = "template"
const stringTag = "string"
const uInt32Tag = "uInt32"
const int32Tag = "int32"
const uInt64Tag = "uInt64"
const int64Tag = "int64"
const decimalTag = "decimal"
const exponentTag = "exponent"
const mantissaTag = "mantissa"

const constantOperation = "constant"

type valueConverter func(string) (interface{}, error)

// Load instance of the Store from the given FAST Templates XML file
func Load(templateFile *os.File) (store.Store, error) {
	decoder := xml.NewDecoder(templateFile)
	xmlTags, err := tokenxml.LoadTagsFrom(decoder)

	if err != nil {
		return store.Store{}, err
	}

	if xmlTags.Type != templatesTag {
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
			return store.Store{}, fmt.Errorf("Could not parse template ID, make sure it is present and uint: %v", err)
		}
		if _, exists := templateStore.Templates[uint32(templateID)]; exists {
			return store.Store{}, fmt.Errorf("Template with ID %d, has already been loaded", templateID)
		}

		templateStore.Templates[uint32(templateID)] = template
	}

	return templateStore, nil
}

func createTemplate(templateRoot *tokenxml.Tag) (store.Template, error) {
	if templateRoot.Type != templateTag {
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
	fieldDetails, err := createFieldDetails(tagInTemplate)
	if err != nil {
		return nil, err
	}

	switch tagInTemplate.Type {
	case stringTag:
		operation, err := getOperation(tagInTemplate, converter.ToString)
		if err != nil {
			return nil, err
		}
		return field.String{FieldDetails: fieldDetails, Operation: operation}, nil
	case uInt32Tag:
		operation, err := getOperation(tagInTemplate, converter.ToUInt32)
		if err != nil {
			return nil, err
		}
		return field.UInt32{FieldDetails: fieldDetails, Operation: operation}, nil
	case int32Tag:
		operation, err := getOperation(tagInTemplate, converter.ToInt32)
		if err != nil {
			return nil, err
		}
		return field.Int32{FieldDetails: fieldDetails, Operation: operation}, nil
	case uInt64Tag:
		operation, err := getOperation(tagInTemplate, converter.ToUInt64)
		if err != nil {
			return nil, err
		}
		return field.UInt64{FieldDetails: fieldDetails, Operation: operation}, nil
	case int64Tag:
		operation, err := getOperation(tagInTemplate, converter.ToInt64)
		if err != nil {
			return nil, err
		}
		return field.Int64{FieldDetails: fieldDetails, Operation: operation}, nil
	case decimalTag:
		if len(tagInTemplate.NestedTags) < 2 {
			exponentOperation, err := getOperation(tagInTemplate, converter.ToExponent)
			if err != nil {
				return nil, err
			}
			exponentField := field.Int32{FieldDetails: fieldDetails, Operation: exponentOperation}
			mantissaOperation, err := getOperation(tagInTemplate, converter.ToMantissa)
			if err != nil {
				return nil, err
			}
			mantissaFieldFieldDetails := fieldDetails
			mantissaFieldFieldDetails.Required = true
			mantissaField := field.Int64{FieldDetails: mantissaFieldFieldDetails, Operation: mantissaOperation}

			return field.Decimal{FieldDetails: fieldDetails, ExponentField: exponentField, MantissaField: mantissaField}, nil
		}
		if len(tagInTemplate.NestedTags) == 2 {
			exponentTag := tagInTemplate.NestedTags[0]
			exponentOperation, err := getOperation(&exponentTag, converter.ToInt32)
			if err != nil {
				return nil, err
			}
			exponentField := field.Int32{FieldDetails: fieldDetails, Operation: exponentOperation}

			mantissaTag := tagInTemplate.NestedTags[1]
			mantissaOperation, err := getOperation(&mantissaTag, converter.ToInt64)
			if err != nil {
				return nil, err
			}
			mantissaFieldFieldDetails := fieldDetails
			mantissaFieldFieldDetails.Required = true
			mantissaField := field.Int64{FieldDetails: mantissaFieldFieldDetails, Operation: mantissaOperation}

			return field.Decimal{FieldDetails: fieldDetails, ExponentField: exponentField, MantissaField: mantissaField}, nil
		}
		return nil, fmt.Errorf("decimal must be declared with either no operation (empty), or with <exponent/> and <mantissa/>")
	default:
		return nil, fmt.Errorf("Unsupported tag type: %s", tagInTemplate.Type)
	}
}

func getOperation(tagInTemplate *tokenxml.Tag, converter valueConverter) (operation.Operation, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return operation.None{}, nil
	}

	operationTag := tagInTemplate.NestedTags[0]

	switch operationTag.Type {
	case constantOperation:
		operation := operation.Constant{}
		constant := operationTag.Attributes["value"]
		if constant == "" {
			return nil, fmt.Errorf("No value specified for constant operation")
		}
		constantAsCorrectValue, err := converter(constant)
		if err != nil {
			return nil, err
		}
		operation.ConstantValue = constantAsCorrectValue
		return operation, nil
	default:
		return nil, fmt.Errorf("Unsupported operation type: %s", operationTag)
	}
}
