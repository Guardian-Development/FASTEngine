package fieldint64

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

//<int64>
//	<delta />
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", true, testLog))

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

//<int64>
//	<delta value="7"/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorWithInitialValueEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(9)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int64Field", true, testLog), 7)

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

//<int64>
//	<delta value="7"/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedPositiveDeltaValueWithNegativePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 9 = 10001001
	messageAsBytes := bytes.NewBuffer([]byte{137})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int64Field", true, testLog), 7)

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(-7)))
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
//	<delta value="7"/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedNegativeDeltaValueWithPositivePreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 -29 = 11100011
	messageAsBytes := bytes.NewBuffer([]byte{227})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(-2)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int64Field", true, testLog), 7)

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(27)))
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
//	<delta value="7"/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedMaxPositiveInt64ToMaxNegativeInt64(t *testing.T) {
	// Arrange pmap = 10000000 -18446744073709551615 = 01111110 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000001
	messageAsBytes := bytes.NewBuffer([]byte{126, 0, 0, 0, 0, 0, 0, 0, 0, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(math.MinInt64)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int64Field", true, testLog), 7)

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(math.MaxInt64)))
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
//	<delta value="7"/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedMaxNegativeInt64ToMaxPositiveInt64(t *testing.T) {
	// Arrange pmap = 10000000 18446744073709551615 = 00000001 01111111 01111111 01111111 01111111 01111111 01111111 01111111 01111111 11111111
	messageAsBytes := bytes.NewBuffer([]byte{1, 127, 127, 127, 127, 127, 127, 127, 127, 255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(math.MaxInt64)
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "Int64Field", true, testLog), 7)

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(math.MinInt64)))
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
//	<delta/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedPositiveDeltaValueOverflowsInt64Error(t *testing.T) {
	// Arrange pmap = 10000000 1 = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", true, testLog))

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(math.MaxInt64)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int64>
//	<delta/>
//</int64>
func TestCanDeseraliseRequiredInt64DeltaOperatorEncodedNegativeDeltaValueOverflowsInt64Error(t *testing.T) {
	// Arrange pmap = 10000000 -1 = 11111111
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", true, testLog))

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(math.MinInt64)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<int64 presence="optional">
//	<delta />
//</int64>
func TestCanDeseraliseOptionalInt64DeltaOperatorEncodedNullPreviousValueReturnsBaseValue(t *testing.T) {
	// Arrange pmap = 10000000 -1 = 11111111
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(-1)
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", false, testLog))

	// Act
	dict.SetValue("Int64Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int64 presence="optional">
//	<delta/>
//</int64>
func TestCanDeseraliseOptionalInt64DeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange pmap = 10000000 null = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", false, testLog))

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(27)))
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
//	<delta/>
//</int64>
func TestRequiresPmapReturnsFalseForRequiredInt64DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<int64 presence="optional">
//	<delta/>
//</int64>
func TestRequiresPmapReturnsFalseForOptionalInt64DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "Int64Field", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
