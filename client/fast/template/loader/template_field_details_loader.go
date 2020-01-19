package loader

import (
	"fmt"
	"strconv"

	"github.com/Guardian-Development/fastengine/internal/fast/field"
	tokenxml "github.com/Guardian-Development/fastengine/internal/xml"
)

func createFieldDetails(tagInTemplate *tokenxml.Tag) (field.Field, error) {
	fieldDetails := field.Field{}

	ID, err := getFieldID(tagInTemplate)
	if err != nil {
		return fieldDetails, err
	}
	fieldDetails.ID = ID

	required, err := getRequiredField(tagInTemplate)
	if err != nil {
		return fieldDetails, err
	}
	fieldDetails.Required = required

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

func getRequiredField(tagInTemplate *tokenxml.Tag) (bool, error) {
	fieldPresence := tagInTemplate.Attributes["presence"]

	if fieldPresence == "" {
		return true, nil
	}

	if fieldPresence == "optional" {
		return false, nil
	}

	if fieldPresence == "mandatory" {
		return true, nil
	}

	return false, fmt.Errorf("Unsupported presence attribute, must be optional or mandatory but found: %s", fieldPresence)
}
