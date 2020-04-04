package fielduint32

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

//<uInt32>
//	<constant value="132" />
//</uInt32>
func TestCanDeseraliseRequiredUInt32ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt32Field", true), 132)

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

//<uInt32 presence="optional">
//	<constant value="132" />
//</uInt32>
func TestCanDeseraliseOptionalUInt32ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt32Field", false), 132)

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

//<uInt32 presence="optional">
//	<constant value="132" />
//</uInt32>
func TestCanDeseraliseOptionalUInt32ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := uint32(132)
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt32Field", false), 132)

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

//<uInt32>
//	<constant value="132" />
//</uInt32>
func TestRequiresPmapReturnsFalseForRequiredUInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt32Field", true), 132)

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<uInt32 presence="optional">
//	<constant value="132" />
//</uInt32>
func TestRequiresPmapReturnsTrueForOptionalUInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "UInt32Field", false), 132)

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
