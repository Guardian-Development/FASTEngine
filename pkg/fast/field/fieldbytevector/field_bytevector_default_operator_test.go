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
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseByteVectorDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

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
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseByteVectorDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

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
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000011 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{131, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

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
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

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
//	<default />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDefaultOperation(properties.New(1, "ByteVectorField", false, testLog))

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

//<byteVector>
//	<default value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsTrueForRequiredByteVectorDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<byteVector presence="optional">
//	<default value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsTrueForOptionalByteVectorDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
