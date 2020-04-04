package fieldasciistring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

//<string>
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseAsciiStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "AsciiStringField", true), "TEST2")

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string>
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseAsciiStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "AsciiStringField", true), "TEST2")

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "AsciiStringField", false), "TEST2")

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "AsciiStringField", false), "TEST2")

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string presence="optional">
//	<default />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDefaultOperation(properties.New(1, "AsciiStringField", false))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<string>
//	<default value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForRequiredAsciiStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "AsciiStringField", true), "TEST2")

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string presence="optional">
//	<default />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDefaultOperation(properties.New(1, "AsciiStringField", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
