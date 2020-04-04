package fielduint64

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"math"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<uint64>
//	<delta />
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(2)
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", true))

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

//<uint64>
//	<delta value="7"/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorWithInitialValueEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(9)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt64Field", true), 7)

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

//<uint64>
//	<delta value="7"/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedPositiveDeltaValuePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 9 = 10001001
	messageAsBytes := bytes.NewBuffer([]byte{137})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(11)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt64Field", true), 7)

	// Act
	dictionary.SetValue("UInt64Field", fix.NewRawValue(uint64(2)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint64>
//	<delta value="7"/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedNegativeDeltaValuePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 -29 = 11100011
	messageAsBytes := bytes.NewBuffer([]byte{227})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(500)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt64Field", true), 7)

	// Act
	dictionary.SetValue("UInt64Field", fix.NewRawValue(uint64(529)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint64>
//	<delta value="7"/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedMaxPositiveUInt64ToZero(t *testing.T) {
	// Arrange pmap = 10000000 -18446744073709551615 = 01111110 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000001
	messageAsBytes := bytes.NewBuffer([]byte{126, 0, 0, 0, 0, 0, 0, 0, 0, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(0)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt64Field", true), 7)

	// Act
	dictionary.SetValue("UInt64Field", fix.NewRawValue(uint64(math.MaxUint64)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint64>
//	<delta value="7"/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedZeroToMaxPositiveUInt64(t *testing.T) {
	// Arrange pmap = 10000000 18446744073709551615 = 00000001 01111111 01111111 01111111 01111111 01111111 01111111 01111111 01111111 11111111
	messageAsBytes := bytes.NewBuffer([]byte{1, 127, 127, 127, 127, 127, 127, 127, 127, 255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(math.MaxUint64)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UIn64Field", true), 7)

	// Act
	dictionary.SetValue("UIn64Field", fix.NewRawValue(uint64(0)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint64>
//	<delta/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedPositiveDeltaValueOverflowsUInt64Error(t *testing.T) {
	// Arrange pmap = 10000000 1 = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", true))

	// Act
	dictionary.SetValue("UInt64Field", fix.NewRawValue(uint64(math.MaxUint64)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "18446744073709551615 + 1 would overflow uint64" {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint64>
//	<delta/>
//</uint64>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedNegativeDeltaValueOverflowsUInt64Error(t *testing.T) {
	// Arrange pmap = 10000000 -1 = 11111111
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", true))

	// Act
	dictionary.SetValue("UInt64Field", fix.NewRawValue(uint64(0)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "0 + 1 would overflow uint64" {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint64 presence="optional">
//	<delta />
//</uint64>
func TestCanDeseraliseOptionalUInt64DeltaOperatorEncodedPreviousNullValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000 1 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", false))

	// Act
	dictionary.SetValue("UInt64Field", fix.NullValue{})
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "you cannot apply a delta to a null previous value" {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint64 presence="optional">
//	<delta/>
//</uint64>
func TestCanDeseraliseOptionalUInt64DeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange pmap = 10000000 null = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", false))

	// Act
	dictionary.SetValue("UInt64Field", fix.NewRawValue(uint64(27)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<uint64>
//	<delta/>
//</uint64>
func TestRequiresPmapReturnsFalseForRequiredUInt64DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", true))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<uint64 presence="optional">
//	<delta/>
//</uint64>
func TestRequiresPmapReturnsFalseForOptionalUInt64DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt64Field", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
