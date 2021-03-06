package fieldasciistring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"log"
	"os"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

var testLog = log.New(os.Stdout, "", log.LstdFlags)

//<string />
func TestCanDeseraliseRequiredAsciiString(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := New(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string presence="optional"/>
func TestCanDeseraliseOptionalAsciiStringPresent(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := New(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string presence="optional"/>
func TestCanDeseraliseOptionalAsciiStringNull(t *testing.T) {
	// Arrange TEST1 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

// <string />
func TestDictionaryIsUpdatedWithAssignedValueWhenAsciiStringValueReadFromStream(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue("TEST1")}
	unitUnderTest := New(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("AsciiStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <string presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenAsciiStringNilValueReadFromStream(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedValue := dictionary.EmptyValue{}
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("AsciiStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<string />
func TestRequiresPmapReturnsFalseForRequiredAsciiStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<string presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalAsciiStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
