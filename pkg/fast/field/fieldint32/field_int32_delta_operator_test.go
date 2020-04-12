package fieldint32

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

//<int32>
//	<delta />
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", true, testLog))

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

//<int32>
//	<delta value="7"/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorWithInitialValueEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(9)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int32Field", true, testLog), 7)

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

//<int32>
//	<delta value="7"/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedPositiveDeltaValueWithNegativePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 9 = 10001001
	messageAsBytes := bytes.NewBuffer([]byte{137})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int32Field", true, testLog), 7)

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(-7)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32>
//	<delta value="7"/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedNegativeDeltaValueWithPositivePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 -29 = 11100011
	messageAsBytes := bytes.NewBuffer([]byte{227})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(-2)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int32Field", true, testLog), 7)

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(27)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32>
//	<delta value="7"/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedMaxPositiveInt32ToMaxNegativeInt32(t *testing.T) {
	// Arrange pmap = 10000000 -4294967295 = 01111111 01110000 00000000 00000000 00000000 10000001
	messageAsBytes := bytes.NewBuffer([]byte{127, 112, 0, 0, 0, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(math.MinInt32)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int32Field", true, testLog), 7)

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(math.MaxInt32)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32>
//	<delta value="7"/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedMaxNegativeInt32ToMaxPositiveInt64(t *testing.T) {
	// Arrange pmap = 10000000 4294967295 = 00000000 00001111 01111111 01111111 01111111 01111111
	messageAsBytes := bytes.NewBuffer([]byte{0, 15, 127, 127, 127, 255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(math.MaxInt32)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int32Field", true, testLog), 7)

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(math.MinInt32)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32>
//	<delta/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedPositiveDeltaValueOverflowsInt32Error(t *testing.T) {
	// Arrange pmap = 10000000 1 = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", true, testLog))

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(math.MaxInt32)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int32>
//	<delta/>
//</int32>
func TestCanDeseraliseRequiredInt32DeltaOperatorEncodedNegativeDeltaValueOverflowsInt32Error(t *testing.T) {
	// Arrange pmap = 10000000 -1 = 11111111
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", true, testLog))

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(math.MinInt32)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int32 presence="optional">
//	<delta />
//</int32>
func TestCanDeseraliseOptionalInt32DeltaOperatorEncodedNullPreviouValueReturnsBaseValue(t *testing.T) {
	// Arrange pmap = 10000000 -1 = 11111111
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(-1)
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", false, testLog))

	// Act
	dict.SetValue("Int32Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32 presence="optional">
//	<delta/>
//</int32>
func TestCanDeseraliseOptionalInt32DeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange pmap = 10000000 null = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", false, testLog))

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(27)))
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
//	<delta/>
//</int32>
func TestRequiresPmapReturnsFalseForRequiredInt32DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<int32 presence="optional">
//	<delta/>
//</int32>
func TestRequiresPmapReturnsFalseForOptionalInt32DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int32Field", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
