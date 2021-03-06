package fieldasciistring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

//<string>
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseRequiredAsciiStringConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewConstantOperation(properties.New(1, "AsciiStringField", true, testLog), "TEST2")

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

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewConstantOperation(properties.New(1, "AsciiStringField", false, testLog), "TEST2")

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

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewConstantOperation(properties.New(1, "AsciiStringField", false, testLog), "TEST2")

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

//<string>
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsFalseForRequiredAsciiStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "AsciiStringField", true, testLog), "TEST2")

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "AsciiStringField", false, testLog), "TEST2")

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
