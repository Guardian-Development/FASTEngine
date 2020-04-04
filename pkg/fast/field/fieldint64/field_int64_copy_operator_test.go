package fieldint64

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<int64>
//	<copy />
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := NewCopyOperation(properties.New(1, "Int64Field", true))

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

//<int64>
//	<copy value="12"/>
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(12)
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(12))

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

//<int64>
//	<copy value="15"/>
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(7)
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(15))

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(7)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int64>
//	<copy />
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "Int64Field", true))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || err.Error() != "no value supplied in message and no initial value with required field" {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int64 presence="optional">
//	<copy />
//</int64>
func TestCanDeseraliseOptionalInt64CopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "Int64Field", false))

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

//<int64 presence="optional">
//	<copy value="12"/>
//</int64>
func TestCanDeseraliseOptionalInt64CopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "Int64Field", false))

	// Act
	dict.SetValue("Int64Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<int64>
//	<copy value="1"/>
//</int64>
func TestRequiresPmapReturnsTrueForRequiredInt64CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(1))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int64 presence="optional">
//	<copy value="7"/>
//</int64>
func TestRequiresPmapReturnsTrueForOptionalInt64CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "Int64Field", false), int64(1))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
