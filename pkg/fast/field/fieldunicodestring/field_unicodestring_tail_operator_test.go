package fieldunicodestring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<string charset="unicode">
//	<tail value="TEST1"/>
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorNotEncodedNoPreviousValueReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "UnicodeStringField", true), "TEST1")

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
//	<tail value="TEST1"/>
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorNotEncodedNoPreviousValueReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "UnicodeStringField", true), "TEST1")

	// Act
	dictionary.SetValue("UnicodeStringField", fix.NewRawValue("TEST2"))
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
//	<tail value="TEST1"/>
//</string>
func TestCanDeseraliseOptionlUnicodeStringTailOperatorNotEncodedEmptyPreviousValueReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "UnicodeStringField", false), "TEST1")

	// Act
	dictionary.SetValue("UnicodeStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<string charset="unicode">
//	<tail value="TEST: TEST1"/>
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorEncodedReturnsInitialValueCombinedWithValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 10000101 01010100 01000101 01010011 01010100 00110010
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 50})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST: TEST2"
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "UnicodeStringField", true), "TEST: TEST1")

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
//	<tail />
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 10000101 01010100 01000101 01010011 01010100 00110010
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 50})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "UnicodeStringField", true))

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
//	<tail />
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorEncodedPreviousValueReturnsPreviousValueCombinedWithValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 10000101 01010100 01000101 01010011 01010100 00110010
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 50})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST: TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "UnicodeStringField", true))

	// Act
	dictionary.SetValue("UnicodeStringField", fix.NewRawValue("TEST: TEST1"))
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
//	<tail />
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorEncodedPreviousValueReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 10000101 01010100 01000101 01010011 01010100 00110010
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 50})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "UnicodeStringField", true))

	// Act
	dictionary.SetValue("UnicodeStringField", fix.NewRawValue("1"))
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
//	<tail />
//</string>
func TestCanDeseraliseUnicodeStringTailOperatorEncodedPreviousValueEmptyReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 10000101 01010100 01000101 01010011 01010100 00110010
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 50})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "UnicodeStringField", true))

	// Act
	dictionary.SetValue("UnicodeStringField", fix.NullValue{})
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
//	<tail />
//</string>
func TestRequiresPmapReturnsTrueForRequiredUnicodeStringTailOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewTailOperation(properties.New(1, "UnicodeStringField", true))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string charset="unicode" presence="optional">
//	<tail />
//</string>
func TestRequiresPmapReturnsTrueForOptionalUnicodeStringTailOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewTailOperation(properties.New(1, "UnicodeStringField", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
