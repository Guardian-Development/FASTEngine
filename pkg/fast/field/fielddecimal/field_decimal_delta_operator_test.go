package fielddecimal

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"math"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

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

//<decimal>
//	<delta initialValue="1.2"/>
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorWithInitialValueEncodedValueNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010 (10) man = 10001010 (2)
	messageAsBytes := bytes.NewBuffer([]byte{130, 138})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(220)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", true), -1),
		fieldint64.NewDeltaOperationWithInitialValue(properties.New(1, "DecimalFieldMantissa", true), 12))

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

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedPositiveDeltaValueWithPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010 (10) man = 10001010 (2)
	messageAsBytes := bytes.NewBuffer([]byte{130, 138})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(220)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalField", fix.NewRawValue(float64(1.2)))
	dictionary.SetValue("DecimalFieldExponent", fix.NewRawValue(int32(-1)))
	dictionary.SetValue("DecimalFieldMantissa", fix.NewRawValue(int64(12)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedNegativeDeltaWithPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 11111111 (-1) man = 11101100 (-20)
	messageAsBytes := bytes.NewBuffer([]byte{255, 232})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(-0.12)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalField", fix.NewRawValue(float64(1.2)))
	dictionary.SetValue("DecimalFieldExponent", fix.NewRawValue(int32(-1)))
	dictionary.SetValue("DecimalFieldMantissa", fix.NewRawValue(int64(12)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

// TODO: implement max to min tests once decimal uses big int for its final value

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedPositiveExponentDeltaValueOverflowsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000001 (1)
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalFieldExponent", fix.NewRawValue(int32(63)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R1) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedNegativeExponentDeltaValueOverflowsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 11111111 (-1)
	messageAsBytes := bytes.NewBuffer([]byte{255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalFieldExponent", fix.NewRawValue(int32(-63)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R1) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedPositiveMantissaDeltaValueOverflowsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000001 (1) man = 10000001 (1)
	messageAsBytes := bytes.NewBuffer([]byte{129, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalFieldMantissa", fix.NewRawValue(int64(math.MaxInt64)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal>
//	<delta />
//</decimal>
func TestCanDeseraliseRequiredDecimalDeltaOperatorEncodedNegativeMantissaDeltaValueOverflowsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000001 (1) man = 11111111 (-1)
	messageAsBytes := bytes.NewBuffer([]byte{129, 255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalFieldMantissa", fix.NewRawValue(int64(math.MinInt64)))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R4) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal presence="optional">
//	<delta />
//</decimal>
func TestCanDeseraliseOptionalDecimalDeltaOperatorEncodedNullExponentPreviouValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000001 (1) man = 11111111 (-1)
	messageAsBytes := bytes.NewBuffer([]byte{129, 255})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	dictionary.SetValue("DecimalFieldExponent", fix.NullValue{})
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D6) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal presence="optional">
//	<delta/>
//</decimal>
func TestCanDeseraliseOptionalDecimalDeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000000 (0)
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<decimal>
//	<delta/>
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal presence="optional">
//	<delta/>
//</decimal>
func TestRequiresPmapReturnsFalseForOptionalInt32DeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDeltaOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewDeltaOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
