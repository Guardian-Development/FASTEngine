package engine

import (
	"bytes"
	"os"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fast/template/loader"
)

// TODO: delta
// TODO: fix bugs around reading max integers at the decoding bytes level (can read more than 4 bytes for int32 and same for int64), add tests for min and max values
// TODO: look at the public interface, just have the engine, everything else private
// TODO: rename some of the directories to make them more correct
// TODO: look at test names and make them better
// TODO: engine should be only public element of this library really, with a nice interface to create it from a template file only.
// TODO: logging!!
// TODO: add a series of messages, ranging in complexity, that cover all types we want to test at integration level (integration test directory or something?)
// 			look at codecoverage and how we can use them to generate coverage for the whole project
// TODO: series of messages should focus on state (copy, increment etc) to show the engine works when parsing a feed
// TODO: making everything immutable, use constructor init methods, cleanup what should be public/private
// TODO: evaluate errors properly and make them useful, use custom errors at each level (follow best practices) USE ERROR CODES FROM SPEC!
// TODO: move to an interface / factory based method of generating FIX message itself, so abstraction can be made if needed
// TODO: readme and documentation
// TODO: pretty print FIX message using pipe character to make readable

func TestTemplateIdNotFoundInTemplateStoreErrorReturned(t *testing.T) {
	// Arrange
	message := bytes.NewBuffer([]byte{192, 1, 150, 130, 210, 129, 210, 130, 131})
	file, _ := os.Open("../../../test/test_heartbeat_template.xml")
	templateStore, _ := loader.Load(file)
	fastEngine := New(templateStore)

	// Act
	_, err := fastEngine.Deserialise(message)

	// Assert
	if err == nil || err.Error() != "no template found in store to deserialise message with ID: 150" {
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
	templateStore, _ := loader.Load(file)
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
	templateStore, _ := loader.Load(file)
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
	templateStore, _ := loader.Load(file)
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
