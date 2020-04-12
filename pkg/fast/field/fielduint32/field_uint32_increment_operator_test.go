package fielduint32

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
	"strings"
	"testing"
)

//<uint32>
//	<increment value="5" />
//</uint32>
func TestCanDeseraliseUInt32IncrementOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := uint32(2)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), uint32(5))

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

//<uint32>
//	<increment value="5" />
//</uint32>
func TestCanDeseraliseUInt32IncrementOperatorNotEncodedReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(5)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), uint32(5))

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

//<uint32>
//	<increment value="5" />
//</uint32>
func TestCanDeseraliseUInt32IncrementOperatorNotEncodedPreviousValuePresentReturnsPreviousValueIncrementByOne(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(11)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), uint32(5))

	// Act
	dictionary.SetValue("UInt32Field", fix.NewRawValue(uint32(10)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint32 presence="optional">
//	<increment />
//</uint32>
func TestCanDeseraliseOptionalUInt32IncrementOperatorNotEncodedNoPreviousValueReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "UInt32Field", false, testLog))

	// Act
	dict.SetValue("UInt32Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<uint32>
//	<increment />
//</uint32>
func TestCanDeseraliseRequiredUInt32IncrementOperatorNotEncodedNoPreviousValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "UInt32Field", true, testLog))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint32>
//	<increment value="132" />
//</uint32>
func TestRequiresPmapReturnsTrueForRequiredUInt32IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), uint32(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uint32 presence="optional">
//	<increment value="132" />
//</uint32>
func TestRequiresPmapReturnsTrueForOptionalUInt32IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt32Field", false, testLog), uint32(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
