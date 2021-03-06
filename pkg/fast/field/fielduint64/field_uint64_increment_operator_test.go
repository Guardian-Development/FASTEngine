package fielduint64

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

//<uint64>
//	<increment value="5" />
//</uint64>
func TestCanDeseraliseUInt64IncrementOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dict := dictionary.New()
	expectedMessage := uint64(2)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt64Field", true, testLog), uint64(5))

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

//<uint64>
//	<increment value="5" />
//</uint64>
func TestCanDeseraliseUInt64IncrementOperatorNotEncodedReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint64(5)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt64Field", true, testLog), uint64(5))

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

//<uint64>
//	<increment value="5" />
//</uint64>
func TestCanDeseraliseUInt64IncrementOperatorNotEncodedPreviousValuePresentReturnsPreviousValueIncrementByOne(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint64(11)
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt64Field", true, testLog), uint64(5))

	// Act
	dict.SetValue("UInt64Field", fix.NewRawValue(uint64(10)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint64 presence="optional">
//	<increment />
//</uint64>
func TestCanDeseraliseOptionalUInt64IncrementOperatorNotEncodedNoPreviousValueReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "UInt64Field", false, testLog))

	// Act
	dict.SetValue("UInt64Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<uint64>
//	<increment />
//</uint64>
func TestCanDeseraliseRequiredUInt64IncrementOperatorNotEncodedNoPreviousValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewIncrementOperation(properties.New(1, "UInt64Field", true, testLog))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint64>
//	<increment value="132" />
//</uint64>
func TestRequiresPmapReturnsTrueForRequiredUInt64IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt64Field", true, testLog), uint64(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uint64 presence="optional">
//	<increment value="132" />
//</uint64>
func TestRequiresPmapReturnsTrueForOptionalUInt64IncrementOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewIncrementOperationWithInitialValue(properties.New(1, "UInt64Field", false, testLog), uint64(132))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
