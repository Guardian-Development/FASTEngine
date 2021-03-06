package fieldunicodestring

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

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorAppendToBaseValue(t *testing.T) {
	// Arrange length = 1000000 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{128, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

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

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorPrependToBaseValue(t *testing.T) {
	// Arrange length = 1111111 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{255, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

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

//<string charset="unicode">
//	<delta value="THE TEST IS: "/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorAppendToInitialValue(t *testing.T) {
	// Arrange length = 1000000 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{128, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "THE TEST IS: TEST1"
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UnicodeStringField", true, testLog), "THE TEST IS: ")

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

//<string charset="unicode">
//	<delta value=": TEST COMPLETE"/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorPrependToInitialValue(t *testing.T) {
	// Arrange length = 1111111 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{255, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1: TEST COMPLETE"
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UnicodeStringField", true, testLog), ": TEST COMPLETE")

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

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorAppendWithOverwiteToPreviousValue(t *testing.T) {
	// Arrange length = 10000100 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{132, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "THE TEST IS: TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

	// Act
	dict.SetValue("UnicodeStringField", fix.NewRawValue("THE TEST IS: OVER"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorPrependWithOverwriteToPreviousValue(t *testing.T) {
	// Arrange length = 11111010 (-6) TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{250, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1: TEST COMPLETE"
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

	// Act
	dict.SetValue("UnicodeStringField", fix.NewRawValue("TEST2: TEST COMPLETE"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiredAsciiStringDeltaOperatorAppendWithOverwiteTooLargeReturnsError(t *testing.T) {
	// Arrange length = 10000100 (4) TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{132, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

	// Act
	dict.SetValue("UnicodeStringField", fix.NewRawValue("FAI"))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D7) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiredUnicodeStringDeltaOperatorPrependWithOverwiteTooLargeReturnsError(t *testing.T) {
	// Arrange length = 11111011 (-5) TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{251, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

	// Act
	dict.SetValue("UnicodeStringField", fix.NewRawValue("FAI"))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D7) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<string charset="unicode" presence="optional">
//	<delta/>
//</string>
func TestOptionalUnicodeStringDeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", false, testLog))

	// Act
	dict.SetValue("UnicodeStringField", fix.NewRawValue("TEST2: TEST COMPLETE"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<string charset="unicode" presence="optional">
//	<delta/>
//</string>
func TestOptionalUnicodeStringDeltaOperatorEncodedPreviousNullValueUsesBaseValue(t *testing.T) {
	// Arrange length = 10000001 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{129, 133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", false, testLog))

	// Act
	dict.SetValue("UnicodeStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string charset="unicode">
//	<delta/>
//</string>
func TestRequiresPmapReturnsFalseForRequiredUnicodeStringDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<string charset="unicode" presence="optional">
//	<delta/>
//</string>
func TestRequiresPmapReturnsFalseForOptionalUnicodeStringDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "UnicodeStringField", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
