package loadproperties

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"math/rand"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// Load id, name, and required presence of field
func Load(tagInTemplate *xml.Tag) (properties.Properties, error) {
	ID, err := getFieldID(tagInTemplate)
	if err != nil {
		return properties.Properties{}, err
	}

	required, err := getRequiredField(tagInTemplate)
	if err != nil {
		return properties.Properties{}, err
	}

	name := tagInTemplate.Attributes["name"]
	if name == "" {
		name = getRandomName(tagInTemplate.Type)
	}

	fieldDetails := properties.New(ID, name, required)
	return fieldDetails, nil
}

func getRandomName(fieldName string) string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	randomPart := string(b)
	return fmt.Sprintf("%s-%s", fieldName, randomPart)
}

func getFieldID(tagInTemplate *xml.Tag) (uint64, error) {
	fieldID := tagInTemplate.Attributes["id"]

	if fieldID == "" {
		return 0, fmt.Errorf("every template field must have an id specified")
	}

	ID, err := strconv.ParseUint(fieldID, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("unable to parse ID for field: %s", fieldID)
	}

	return ID, nil
}

func getRequiredField(tagInTemplate *xml.Tag) (bool, error) {
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

	return false, fmt.Errorf("unsupported presence attribute, must be optional or mandatory but found: %s", fieldPresence)
}
