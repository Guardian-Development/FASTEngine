package fieldint32

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
	"testing"
)

//<int32>
//	<increment value="5" />
//</int32>
func TestCanDeseraliseInt32IncrementOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int32Field", true), int32(5))

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

//<int32>
//	<increment value="5" />
//</int32>
func TestCanDeseraliseInt32IncrementOperatorNotEncodedReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(5)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int32Field", true), int32(5))

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

//<int32>
//	<increment value="5" />
//</int32>
func TestCanDeseraliseInt32IncrementOperatorNotEncodedPreviousValuePresentReturnsPreviousValueIncrementByOne(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(11)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int32Field", true), int32(5))

	// Act
	dictionary.SetValue("Int32Field", fix.NewRawValue(int32(10)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32 presence="optional">
//	<increment />
//</int32>
func TestCanDeseraliseOptionalInt32IncrementOperatorNotEncodedNoPreviousValueReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "Int32Field", false))

	// Act
	dict.SetValue("Int32Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<int32>
//	<increment />
//</int32>
func TestCanDeseraliseRequiredInt32IncrementOperatorNotEncodedNoPreviousValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "Int32Field", true))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "no value supplied in message and no initial value with required field" {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int32>
//	<increment value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForRequiredInt32IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int32Field", true), int32(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int32 presence="optional">
//	<increment value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForOptionalInt32IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "Int32Field", false), int32(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
