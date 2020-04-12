package engine

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
)

// TODO: readme and documentation (cleanup all warnings)
// TODO: code cleanup and doc strings

func TestTemplateIdNotFoundInTemplateStoreErrorReturned(t *testing.T) {
	// Arrange
	message := bytes.NewBuffer([]byte{192, 1, 150, 130, 210, 129, 210, 130, 131})
	fastEngine, _ := NewFromTemplateFile("../../test/test_heartbeat_template.xml", log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile))

	// Act
	_, err := fastEngine.Deserialise(message)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D9) {
		t.Errorf("Expected error message informing user template ID is not found in store for message, but got: %v", err)
	}
}

func TestCanDeserialiseHeartbeatMessageBasedOnTemplateInTemplateStore(t *testing.T) {
	// Arrange
	/*
		Message format:
		11000000           pmap
		00000001 10010000  template 144
		10001010           34 = 10
		10001011           52 = 11
	*/
	message := bytes.NewBuffer([]byte{192, 1, 144, 138, 139})
	fastEngine, _ := NewFromTemplateFile("../../test/test_heartbeat_template.xml", log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile))

	// Act
	fixMessage, _ := fastEngine.Deserialise(message)

	// Assert
	fixMessageAsString := fixMessage.String()
	if fixMessageAsString != "1128=9|35=0|34=10|52=11|" {
		t.Errorf("Expected message and actual message were not equal, actual: %s", fixMessageAsString)
	}
}

func TestCanDeserialiseMessageWithOptionalValueNotPresent(t *testing.T) {
	// Arrange
	/*
		Message format:
		11000000           pmap
		00000001 10010000  template 144
		10000000           34 = Nil
		10001010           52 = 10
	*/
	message := bytes.NewBuffer([]byte{192, 1, 144, 128, 138})
	fastEngine, _ := NewFromTemplateFile("../../test/test_optional_value_template.xml", log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile))

	// Act
	fixMessage, _ := fastEngine.Deserialise(message)

	// Assert
	fixMessageAsString := fixMessage.String()
	if fixMessageAsString != "1128=9|35=0|34=nil|52=10|" {
		t.Errorf("Expected message and actual message were not equal, actual: %s", fixMessageAsString)
	}
}

func TestCanDeserialiseMessageWithOptionalValuePresent(t *testing.T) {
	// Arrange
	/*
		Message format:
		11000000           pmap
		00000001 10010000  template 144
		10000001           34 = 0
		10001010           52 = 10
	*/
	message := bytes.NewBuffer([]byte{192, 1, 144, 129, 138})
	fastEngine, _ := NewFromTemplateFile("../../test/test_optional_value_template.xml", log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile))

	// Act
	fixMessage, _ := fastEngine.Deserialise(message)

	// Assert
	fixMessageAsString := fixMessage.String()
	if fixMessageAsString != "1128=9|35=0|34=0|52=10|" {
		t.Errorf("Expected message and actual message were not equal, actual: %s", fixMessageAsString)
	}
}

// func printByteArrayAsBits(array *[]byte) {
// 	for _, n := range *array {
// 		fmt.Printf("% 08b", n)
// 	}
// }
