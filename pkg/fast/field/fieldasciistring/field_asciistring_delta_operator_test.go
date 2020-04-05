package fieldasciistring

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<string>
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorAppendToBaseValue(t *testing.T) {
	// Arrange length = 1000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{128, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

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
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorPrependToBaseValue(t *testing.T) {
	// Arrange length = 1111111 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{255, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

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
//	<delta value="THE TEST IS: "/>
//</string>
func TestRequiredAsciiStringDeltaOperatorAppendToInitialValue(t *testing.T) {
	// Arrange length = 1000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{128, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "THE TEST IS: TEST1"
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "AsciiStringField", true), "THE TEST IS: ")

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
//	<delta value=": TEST COMPLETE"/>
//</string>
func TestRequiredAsciiStringDeltaOperatorPrependToInitialValue(t *testing.T) {
	// Arrange length = 1111111 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{255, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1: TEST COMPLETE"
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "AsciiStringField", true), ": TEST COMPLETE")

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
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorAppendWithOverwiteToPreviousValue(t *testing.T) {
	// Arrange length = 10000100 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{132, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "THE TEST IS: TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("THE TEST IS: OVER"))
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
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorPrependWithOverwriteToPreviousValue(t *testing.T) {
	// Arrange length = 11111010 (-6) TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{250, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1: TEST COMPLETE"
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("TEST2: TEST COMPLETE"))
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
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorAppendWithOverwiteTooLargeReturnsError(t *testing.T) {
	// Arrange length = 10000100 (4) TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{132, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("FAI"))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D7) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<string>
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorPrependWithOverwiteTooLargeReturnsError(t *testing.T) {
	// Arrange length = 11111011 (-5) TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{251, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("FAI"))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D7) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<string presence="optional">
//	<delta/>
//</string>
func TestOptionalAsciiStringDeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange length = 10000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", false))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NewRawValue("TEST2: TEST COMPLETE"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<string presence="optional">
//	<delta/>
//</string>
func TestOptionalAsciiStringDeltaOperatorEncodedPreviousNullValueReturnsError(t *testing.T) {
	// Arrange length = 10000001 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{129, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", false))

	// Act
	dictionary.SetValue("AsciiStringField", fix.NullValue{})
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D6) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<string>
//	<delta/>
//</string>
func TestRequiresPmapReturnsFalseForRequiredAsciiStringDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", true))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<string presence="optional">
//	<delta/>
//</string>
func TestRequiresPmapReturnsFalseForOptionalAsciiStringDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "AsciiStringField", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
