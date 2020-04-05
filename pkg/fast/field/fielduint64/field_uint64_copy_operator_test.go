package fielduint64

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

//<uInt64>
//	<copy />
//</uInt64>
func TestCanDeseraliseRequiredUInt64CopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := uint64(2)
	unitUnderTest := NewCopyOperation(properties.New(1, "UInt64Field", true))

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

//<uInt64>
//	<copy value="12"/>
//</uInt64>
func TestCanDeseraliseRequiredUInt64CopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(12)
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "UInt64Field", true), 12)

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

//<uInt64>
//	<copy value="15"/>
//</uInt64>
func TestCanDeseraliseRequiredUInt64CopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint64(7)
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "UInt64Field", true), 15)

	// Act
	dict.SetValue("UInt64Field", fix.NewRawValue(uint64(7)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uInt64>
//	<copy />
//</uInt64>
func TestCanDeseraliseRequiredUInt64CopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "UInt64Field", true))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uInt64 presence="optional">
//	<copy />
//</uInt64>
func TestCanDeseraliseOptionalUInt64CopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "UInt64Field", false))

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

//<uInt64 presence="optional">
//	<copy value="12"/>
//</uInt64>
func TestCanDeseraliseOptionalUInt64CopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "UInt64Field", false), 12)

	// Act
	dict.SetValue("UInt64Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<uInt64>
//	<copy value="1"/>
//</uInt64>
func TestRequiresPmapReturnsTrueForRequiredUInt64CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "UInt64Field", true), 1)

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uInt64 presence="optional">
//	<copy value="7"/>
//</uInt64>
func TestRequiresPmapReturnsTrueForOptionalUInt64CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "UInt64Field", false), 7)

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
