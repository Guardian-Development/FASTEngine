package fieldunicodestring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<string charset="unicode"/>
func TestCanDeseraliseRequiredUnicodeString(t *testing.T) {
	// Arrange TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := New(properties.New(1, "UnicodeStringField", true))

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

//<string charset="unicode" presence="optional"/>
func TestCanDeseraliseOptionalUnicodeStringPresent(t *testing.T) {
	// Arrange TEST1 = 10000110 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{134, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := New(properties.New(1, "UnicodeStringField", false))

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

//<string charset="unicode" presence="optional"/>
func TestCanDeseraliseOptionalUnicodeStringNull(t *testing.T) {
	// Arrange TEST1 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "UnicodeStringField", false))

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

//<string charset="unicode" />
func TestDictionaryIsUpdatedWithAssignedValueWhenUnicodeStringValueReadFromStream(t *testing.T) {
	// Arrange TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue("TEST1")}
	unitUnderTest := New(properties.New(1, "UnicodeStringField", true))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("UnicodeStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<string charset="unicode" presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenUnicodeStringNilValueReadFromStream(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := New(properties.New(1, "UnicodeStringField", false))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("UnicodeStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<string charset="unicode"/>
func TestRequiresPmapReturnsFalseForRequiredUnicodeStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "UnicodeStringField", true))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<string charset="unicode" presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalUnicodeStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "UnicodeStringField", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
