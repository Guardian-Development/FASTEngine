package fieldasciistring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<string>
//	<tail value="TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringTailOperatorNotEncodedNoPreviousValueReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "AsciiStringField", true, testLog), "TEST1")

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

//<string>
//	<tail value="TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringTailOperatorNotEncodedNoPreviousValueReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "AsciiStringField", true, testLog), "TEST1")

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("TEST2"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string presence="optional">
//	<tail value="TEST1"/>
//</string>
func TestCanDeseraliseOptionlAsciiStringTailOperatorNotEncodedEmptyPreviousValueReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "AsciiStringField", false, testLog), "TEST1")

	// Act
	dictionary.SetValue("AsciiStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<string>
//	<tail value="TEST: TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringTailOperatorEncodedReturnsInitialValueCombinedWithValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST: TEST2"
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "AsciiStringField", true, testLog), "TEST: TEST1")

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

//<string>
//	<tail />
//</string>
func TestCanDeseraliseAsciiStringTailOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "AsciiStringField", true, testLog))

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

//<string>
//	<tail />
//</string>
func TestCanDeseraliseAsciiStringTailOperatorEncodedPreviousValueReturnsPreviousValueCombinedWithValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST: TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("TEST: TEST1"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string>
//	<tail />
//</string>
func TestCanDeseraliseAsciiStringTailOperatorEncodedPreviousValueReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("1"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string>
//	<tail />
//</string>
func TestCanDeseraliseAsciiStringTailOperatorEncodedPreviousValueEmptyReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST2 = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := NewTailOperation(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string>
//	<tail />
//</string>
func TestRequiresPmapReturnsTrueForRequiredAsciiStringTailOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewTailOperation(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string presence="optional">
//	<tail />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringTailOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewTailOperation(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
