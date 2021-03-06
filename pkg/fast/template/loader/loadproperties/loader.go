package loadproperties

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// Load id, name, and required presence of field
func Load(tagInTemplate *xml.Tag, logger *log.Logger) (properties.Properties, error) {
	ID, err := getFieldID(tagInTemplate)
	if err != nil {
		logger.Printf("error loading id for tag from xml: %v", tagInTemplate.Attributes)
		return properties.Properties{}, err
	}

	required, err := getRequiredField(tagInTemplate)
	if err != nil {
		logger.Printf("error getting required option for tag from xml: %v", tagInTemplate.Attributes)
		return properties.Properties{}, err
	}

	name := tagInTemplate.Attributes["name"]
	if name == "" {
		name = getRandomName(tagInTemplate.Type)
	}

	fieldDetails := properties.New(ID, name, required, logger)
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
		return 0, nil
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
