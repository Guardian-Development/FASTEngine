package fieldint64

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"testing"
)

//<int64>
//	<increment value="5" />
//</int64>
func TestCanDeseraliseInt64IncrementOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(5))

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
//	<increment value="5" />
//</int64>
func TestCanDeseraliseInt64IncrementOperatorNotEncodedReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(5)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(5))

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
//	<increment value="5" />
//</int64>
func TestCanDeseraliseInt64IncrementOperatorNotEncodedPreviousValuePresentReturnsPreviousValueIncrementByOne(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(11)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(5))

	// Act
	dictionary.SetValue("Int64Field", fix.NewRawValue(int64(10)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int64 presence="optional">
//	<increment />
//</int64>
func TestCanDeseraliseOptionalInt64IncrementOperatorNotEncodedNoPreviousValueReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "Int64Field", false))

	// Act
	dict.SetValue("Int64Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<int64>
//	<increment />
//</int64>
func TestCanDeseraliseRequiredInt64IncrementOperatorNotEncodedNoPreviousValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "Int64Field", true))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "no value supplied in message and no initial value with required field" {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int64>
//	<increment value="132" />
//</int64>
func TestRequiresPmapReturnsTrueForRequiredInt64IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int64Field", true), int64(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int64 presence="optional">
//	<increment value="132" />
//</int64>
func TestRequiresPmapReturnsTrueForOptionalInt64IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int64Field", false), int64(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
