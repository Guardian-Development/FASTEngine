package template

import (
	"fmt"
	"strconv"

	tokenxml "github.com/Guardian-Development/fastengine/internal/xml"
)

func createTemplate(templateRoot *tokenxml.Tag) (Template, error) {
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

func createTemplateUnit(tagInTemplate *tokenxml.Tag) (Unit, error) {
	switch tagInTemplate.Type {
	case stringTag:
		field := FieldString{}
		fieldDetails, err := createFieldDetails(tagInTemplate)
		field.fieldDetails = fieldDetails
		if err != nil {
			return nil, err
		}
		return field, nil
	case uInt32Tag:
		field := FieldUInt32{}
		fieldDetails, err := createFieldDetails(tagInTemplate)
		field.fieldDetails = fieldDetails
		if err != nil {
			return nil, err
		}
		return field, nil
	case uInt64Tag:
		field := FieldUInt64{}
		fieldDetails, err := createFieldDetails(tagInTemplate)
		field.fieldDetails = fieldDetails
		if err != nil {
			return nil, err
		}
		return field, nil
	default:
		return nil, fmt.Errorf("Unsupported tag type: %s", tagInTemplate.Type)
	}
}

func createFieldDetails(tagInTemplate *tokenxml.Tag) (Field, error) {
	fieldDetails := Field{}

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

func getFieldID(tagInTemplate *tokenxml.Tag) (uint64, error) {
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

func getOperation(tagInTemplate *tokenxml.Tag) (Operation, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return OperationNone{}, nil
	}

	operationTag := tagInTemplate.NestedTags[0]

	switch operationTag.Type {
	case constantOperation:
		operation := OperationConstant{}
		constant := operationTag.Attributes["value"]
		if constant == "" {
			return nil, fmt.Errorf("No value specified for constant operation")
		}
		operation.constantValue = constant
		return operation, nil
	default:
		return nil, fmt.Errorf("Unsupported operation type: %s", operationTag)
	}
}
