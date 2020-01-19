package engine

import (
	"bytes"
	"os"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fast/template"
)

// TODO: implement decimal, byte vector, unicode string
// TODO: add a series of messages, ranging in complexity, that cover all types we want to test at integration level
// TODO: look at how the different operations work
// TODO: handle repeating groups
// TODO: create context

func TestTemplateIdNotFoundInTemplateStoreErrorReturned(t *testing.T) {
	// Arrange
	message := bytes.NewBuffer([]byte{192, 1, 150, 130, 210, 129, 210, 130, 131})
	file, _ := os.Open("../../../test/test_heartbeat_template.xml")
	templateStore, _ := template.New(file)
	fastEngine := New(templateStore)

	// Act
	_, err := fastEngine.Deserialise(message)

	// Assert
	if err == nil || err.Error() != "No template found in store to deserialise message with ID: 150" {
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
	file, _ := os.Open("../../../test/test_heartbeat_template.xml")
	templateStore, _ := template.New(file)
	fastEngine := New(templateStore)

	// Act
	fixMessage, _ := fastEngine.Deserialise(message)

	// Assert
	tag1128, _ := fixMessage.GetTag(1128)
	if tag1128 != "9" {
		t.Errorf("Expected: 9, but got: %s", tag1128)
	}
	tag35, _ := fixMessage.GetTag(35)
	if tag35 != "0" {
		t.Errorf("Expected: 0, but got: %s", tag35)
	}
	tag34, _ := fixMessage.GetTag(34)
	if tag34 != uint32(10) {
		t.Errorf("Expected: 10, but got: %s", tag34)
	}
	tag52, _ := fixMessage.GetTag(52)
	if tag52 != uint64(11) {
		t.Errorf("Expected: 11, but got: %s", tag52)
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
	file, _ := os.Open("../../../test/test_optional_value_template.xml")
	templateStore, _ := template.New(file)
	fastEngine := New(templateStore)

	// Act
	fixMessage, _ := fastEngine.Deserialise(message)

	// Assert
	tag1128, _ := fixMessage.GetTag(1128)
	if tag1128 != "9" {
		t.Errorf("Expected: 9, but got: %s", tag1128)
	}
	tag35, _ := fixMessage.GetTag(35)
	if tag35 != "0" {
		t.Errorf("Expected: 0, but got: %s", tag35)
	}
	tag34, _ := fixMessage.GetTag(34)
	if tag34 != nil {
		t.Errorf("Expected: nil, but got: %s", tag34)
	}
	tag52, _ := fixMessage.GetTag(52)
	if tag52 != uint64(10) {
		t.Errorf("Expected: 10, but got: %s", tag52)
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
	file, _ := os.Open("../../../test/test_optional_value_template.xml")
	templateStore, _ := template.New(file)
	fastEngine := New(templateStore)

	// Act
	fixMessage, _ := fastEngine.Deserialise(message)

	// Assert
	tag1128, _ := fixMessage.GetTag(1128)
	if tag1128 != "9" {
		t.Errorf("Expected: 9, but got: %s", tag1128)
	}
	tag35, _ := fixMessage.GetTag(35)
	if tag35 != "0" {
		t.Errorf("Expected: 0, but got: %s", tag35)
	}
	tag34, _ := fixMessage.GetTag(34)
	if tag34 != uint32(0) {
		t.Errorf("Expected: 0, but got: %s", tag34)
	}
	tag52, _ := fixMessage.GetTag(52)
	if tag52 != uint64(10) {
		t.Errorf("Expected: 10, but got: %s", tag52)
	}
}

// func printByteArrayAsBits(array *[]byte) {
// 	for _, n := range *array {
// 		fmt.Printf("% 08b", n)
// 	}
// }
