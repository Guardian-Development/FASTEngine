package fieldint32

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

//<int32>
//	<constant value="132" />
//</int32>
func TestCanDeseraliseRequiredInt32ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "Int32Field", true, testLog), int32(132))

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

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestCanDeseraliseOptionalInt32ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewConstantOperation(properties.New(1, "Int32Field", false, testLog), int32(132))

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

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestCanDeseraliseOptionalInt32ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := int32(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "Int32Field", false, testLog), int32(132))

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

//<int32>
//	<constant value="132" />
//</int32>
func TestRequiresPmapReturnsFalseForRequiredInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "Int32Field", true, testLog), int32(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForOptionalInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "Int32Field", false, testLog), int32(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
