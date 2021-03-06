package fielduint32

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"math"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<uint32>
//	<delta />
//</uint32>
func TestCanDeseraliseRequiredUInt64DeltaOperatorEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(2)
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", true, testLog))

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

//<uint32>
//	<delta value="7"/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorWithInitialValueEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(9)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), 7)

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

//<uint32>
//	<delta value="7"/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorEncodedPositiveDeltaValuePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 9 = 10001001
	messageAsBytes := bytes.NewBuffer([]byte{137})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(11)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), 7)

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(2)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint32>
//	<delta value="7"/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorEncodedNegativeDeltaValuePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 -29 = 11100011
	messageAsBytes := bytes.NewBuffer([]byte{227})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(500)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), 7)

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(529)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint32>
//	<delta value="7"/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorEncodedMaxPositiveUInt32ToZero(t *testing.T) {
	// Arrange pmap = 10000000 -4294967295 = 01111111 01110000 00000000 00000000 00000000 10000001
	messageAsBytes := bytes.NewBuffer([]byte{127, 112, 0, 0, 0, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(0)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UInt32Field", true, testLog), 7)

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(math.MaxUint32)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint32>
//	<delta value="7"/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorEncodedZeroToMaxPositiveUInt32(t *testing.T) {
	// Arrange pmap = 10000000 4294967295 = 00000000 00001111 01111111 01111111 01111111 11111111
	messageAsBytes := bytes.NewBuffer([]byte{0, 15, 127, 127, 127, 255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(math.MaxUint32)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "UIn32Field", true, testLog), 7)

	// Act
	dict.SetValue("UIn32Field", fix.NewRawValue(uint32(0)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint32>
//	<delta/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorEncodedPositiveDeltaValueOverflowsUInt32Error(t *testing.T) {
	// Arrange pmap = 10000000 1 = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", true, testLog))

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(math.MaxUint32)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint32>
//	<delta/>
//</uint32>
func TestCanDeseraliseRequiredUInt32DeltaOperatorEncodedNegativeDeltaValueOverflowsUInt32Error(t *testing.T) {
	// Arrange pmap = 10000000 -1 = 11111111
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", true, testLog))

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(0)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<uint32 presence="optional">
//	<delta />
//</uint32>
func TestCanDeseraliseOptionalUInt32DeltaOperatorEncodedPreviousNullValueReturnsBaseValue(t *testing.T) {
	// Arrange pmap = 10000000 1 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(1)
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", false, testLog))

	// Act
	dict.SetValue("UInt32Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uint32 presence="optional">
//	<delta/>
//</uint32>
func TestCanDeseraliseOptionalUInt32DeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange pmap = 10000000 null = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", false, testLog))

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(27)))
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
//	<delta/>
//</uint32>
func TestRequiresPmapReturnsFalseForRequiredUInt32DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<uint32 presence="optional">
//	<delta/>
//</uint32>
func TestRequiresPmapReturnsFalseForOptionalUInt32DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "UInt32Field", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
