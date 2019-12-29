package engine

import (
	"bytes"
	"os"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fast/template"
)

func TestTemplateIdNotFoundInTemplateStoreErrorReturned(t *testing.T) {
	// Arrange
	/*
		Message format:
		11000000           pmap
		00000001 10010110  template 150
		10000010           384 = 2
		11010010           372 = 82 (R ascii)
		10000001           385 = 1
		11010010           372 = 82 (R ascii)
		10000010           385 = 2
		10000011           96 = 3
	*/
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
	// Act
	// Assert
}

// func printByteArrayAsBits(array *[]byte) {
// 	for _, n := range *array {
// 		fmt.Printf("% 08b", n)
// 	}
// }
