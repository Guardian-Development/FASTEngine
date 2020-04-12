package fieldbytevector

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"reflect"
	"testing"
)

//<byteVector>
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseRequiredByteVectorConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := NewConstantOperation(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewConstantOperation(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

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

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := NewConstantOperation(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsFalseForRequiredByteVectorConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsTrueForOptionalByteVectorConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewConstantOperation(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
