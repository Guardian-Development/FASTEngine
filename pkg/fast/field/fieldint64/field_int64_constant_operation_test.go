package fieldint64

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<int64>
//	<constant value="132" />
//</int64>
func TestCanDeseraliseRequiredInt64ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "Int64Field", true), int64(132))

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

//<int64 presence="optional">
//	<constant value="132" />
//</int64>
func TestCanDeseraliseOptionalInt64ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewConstantOperation(properties.New(1, "Int64Field", false), int64(132))

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

//<int64 presence="optional">
//	<constant value="132" />
//</int64>
func TestCanDeseraliseOptionalInt64ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := int64(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "Int64Field", false), int64(132))

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

//<int64>
//	<constant value="132" />
//</int64>
func TestRequiresPmapReturnsFalseForRequiredInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "Int64Field", true), int64(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<int64 presence="optional">
//	<constant value="132" />
//</int64>
func TestRequiresPmapReturnsTrueForOptionalInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "Int64Field", false), int64(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
