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
//	<copy />
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewCopyOperation(properties.New(1, "AsciiStringField", true, testLog))

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

//<string>
//	<copy value="TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "AsciiStringField", true, testLog), "TEST1")

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

//<string>
//	<copy value="TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "PREVIOUS_VALUE"
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "AsciiStringField", true, testLog), "TEST1")

	// Act
	dict.SetValue("AsciiStringField", fix.NewRawValue("PREVIOUS_VALUE"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string>
//	<copy />
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<string presence="optional">
//	<copy />
//</string>
func TestCanDeseraliseOptionalAsciiStringCopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<string presence="optional">
//	<copy />
//</string>
func TestCanDeseraliseOptionalAsciiStringCopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	dict.SetValue("AsciiStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<string>
//	<copy />
//</string>
func TestRequiresPmapReturnsTrueForRequiredAsciiStringCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperation(properties.New(1, "AsciiStringField", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string presence="optional">
//	<copy />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperation(properties.New(1, "AsciiStringField", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
