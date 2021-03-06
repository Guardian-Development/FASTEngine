package fielduint64

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

//<uInt64>
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseRequiredUInt64ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint64(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt64Field", true, testLog), 132)

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

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt64Field", false, testLog), 132)

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

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := uint64(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt64Field", false, testLog), 132)

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

//<uInt64>
//	<constant value="132" />
//</uInt64>
func TestRequiresPmapReturnsFalseForRequiredUInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt64Field", true, testLog), 132)

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestRequiresPmapReturnsTrueForOptionalUInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt64Field", false, testLog), 132)

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
