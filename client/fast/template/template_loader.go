package template

import (
	"fmt"
	"strconv"

	"github.com/Guardian-Development/fastengine/internal/fast/field"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/xml"
)

const templatesTag = "templates"
const templateTag = "template"

const stringTag = "string"
const uInt32Tag = "uInt32"
const uInt64Tag = "uInt64"

const constantOperation = "constant"

func loadStoreFromXML(xmlTags xml.Tag) (Store, error) {
	templateStore := Store{
		Templates: make(map[uint32]Template),
	}

	for _, templateXMLElement := range xmlTags.NestedTags {
		template, err := createTemplate(&templateXMLElement)
		if err != nil {
			return Store{}, err
		}
		templateID, err := strconv.ParseUint(templateXMLElement.Attributes["id"], 10, 32)

		if err != nil {
			return Store{}, fmt.Errorf("Could not parse template ID, make sure it is present and uint: %v", err)
		}
		if _, exists := templateStore.Templates[uint32(templateID)]; exists {
			return Store{}, fmt.Errorf("Template with ID %d, has already been loaded", templateID)
		}

		templateStore.Templates[uint32(templateID)] = template
	}

	return templateStore, nil
}

func createTemplate(templateRoot *xml.Tag) (Template, error) {
	if templateRoot.Type != templateTag {
		return Template{}, fmt.Errorf("expected to find template tag, but found %s", templateRoot.Type)
	}

	template := Template{
		TemplateUnits: make([]Unit, len(templateRoot.NestedTags)),
	}

	for unitNumber, tagInTemplate := range templateRoot.NestedTags {
		templateUnit, err := createTemplateUnit(&tagInTemplate)

		if err != nil {
			return Template{}, err
		}

		template.TemplateUnits[unitNumber] = templateUnit
	}

	return template, nil
}

func createTemplateUnit(tagInTemplate *xml.Tag) (Unit, error) {
	switch tagInTemplate.Type {
	case stringTag:
		field := field.String{}
		fieldDetails, err := createFieldDetails(tagInTemplate)
		field.FieldDetails = fieldDetails
		if err != nil {
			return nil, err
		}
		return field, nil
	case uInt32Tag:
		field := field.UInt32{}
		fieldDetails, err := createFieldDetails(tagInTemplate)
		field.FieldDetails = fieldDetails
		if err != nil {
			return nil, err
		}
		return field, nil
	case uInt64Tag:
		field := field.UInt64{}
		fieldDetails, err := createFieldDetails(tagInTemplate)
		field.FieldDetails = fieldDetails
		if err != nil {
			return nil, err
		}
		return field, nil
	default:
		return nil, fmt.Errorf("Unsupported tag type: %s", tagInTemplate.Type)
	}
}

func createFieldDetails(tagInTemplate *xml.Tag) (field.Field, error) {
	fieldDetails := field.Field{}

	ID, err := getFieldID(tagInTemplate)
	if err != nil {
		return fieldDetails, err
	}
	fieldDetails.ID = ID

	operation, err := getOperation(tagInTemplate)
	if err != nil {
		return fieldDetails, err
	}

	fieldDetails.Operation = operation

	return fieldDetails, nil
}

func getFieldID(tagInTemplate *xml.Tag) (uint64, error) {
	fieldID := tagInTemplate.Attributes["id"]

	if fieldID == "" {
		return 0, fmt.Errorf("Every template field must have an id specified")
	}

	ID, err := strconv.ParseUint(fieldID, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("Unable to parse ID for field: %s", fieldID)
	}

	return ID, nil
}

func getOperation(tagInTemplate *xml.Tag) (operation.Operation, error) {
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
		operation.ConstantValue = constant
		return operation, nil
	default:
		return nil, fmt.Errorf("Unsupported operation type: %s", operationTag)
	}
}
