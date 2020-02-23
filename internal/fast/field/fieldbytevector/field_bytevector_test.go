package fieldbytevector

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<byteVector />
func TestCanDeseraliseRequiredByteVector(t *testing.T) {
	// Arrange 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := New(properties.New(1, "ByteVectorField", true))

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

//<byteVector />
func TestCanDeseraliseRequiredByteVectorLengthZero(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{}
	unitUnderTest := New(properties.New(1, "ByteVectorField", true))

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

//<byteVector presence="optional"/>
func TestCanDeseraliseOptionalByteVectorPresent(t *testing.T) {
	// Arrange 10000010 10010010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146}
	unitUnderTest := New(properties.New(1, "ByteVectorField", false))

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

//<byteVector presence="optional"/>
func TestCanDeseraliseOptionalByteVectorPresentLengthZero(t *testing.T) {
	// Arrange 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{}
	unitUnderTest := New(properties.New(1, "ByteVectorField", false))

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

//<byteVector presence="optional"/>
func TestCanDeseraliseOptionalByteVectorNull(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "ByteVectorField", false))

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

// <byteVector />
func TestDictionaryIsUpdatedWithAssignedValueWhenByteVectorValueReadFromStream(t *testing.T) {
	// Arrange 10000001 10010010
	messageAsBytes := bytes.NewBuffer([]byte{129, 146})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue([]byte{146})}
	unitUnderTest := New(properties.New(1, "ByteVectorField", true))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("ByteVectorField")
	areEqual := reflect.DeepEqual(expectedValue, result)
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <byteVector presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenByteVectorNilValueReadFromStream(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := New(properties.New(1, "ByteVectorField", false))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("ByteVectorField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<byteVector />
func TestRequiresPmapReturnsFalseForRequiredByteVectorNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "ByteVectorField", true))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<byteVector presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalByteVectorNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "ByteVectorField", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
