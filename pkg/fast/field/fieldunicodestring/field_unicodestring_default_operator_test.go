package fieldunicodestring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

//<string charset="unicode">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseUnicodeStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "UnicodeStringField", true, testLog), "TEST2")

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

//<string charset="unicode">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseUnicodeStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "UnicodeStringField", true, testLog), "TEST2")

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

//<string charset="unicode" presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalUnicodeStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 10000110 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{134, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "UnicodeStringField", false, testLog), "TEST2")

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

//<string charset="unicode" presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalUnicodeStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "UnicodeStringField", false, testLog), "TEST2")

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

//<string charset="unicode" presence="optional">
//	<default />
//</string>
func TestCanDeseraliseOptionalUnicodeStringDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDefaultOperation(properties.New(1, "UnicodeStringField", false, testLog))

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

//<string charset="unicode">
//	<default value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForRequiredUnicodeStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "UnicodeStringField", true, testLog), "TEST2")

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string charset="unicode" presence="optional">
//	<default />
//</string>
func TestRequiresPmapReturnsTrueForOptionalUnicodeStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDefaultOperationWithValue(properties.New(1, "UnicodeStringField", false, testLog), "TEST2")

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
