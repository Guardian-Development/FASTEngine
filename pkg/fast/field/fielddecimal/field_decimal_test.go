package fielddecimal

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

var testLog = log.New(os.Stdout, "", log.LstdFlags)

//<decimal>
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimal(t *testing.T) {
	// Arrange exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

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

//<decimal presence="optional">
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentPresent(t *testing.T) {
	// Arrange exp = 10000011 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{131, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

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

//<decimal presence="optional">
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentPresentMantissaNotEncodedCausesError(t *testing.T) {
	// Arrange exp = 10000011 man = nil
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), "failed to read mantissa value after successful read of exponent") {
		t.Errorf("Expected error message informing user of error when decoding mantissa, but got: %v", err)
	}
}

//<decimal presence="optional">
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentNullMantissaNotEncoded(t *testing.T) {
	// Arrange exp = 10000000 man = NOT ENCODED EVEN THOUGH REQUIRED
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<decimal>
//	<exponent />
//	<mantissa />
//</decimal>
func TestDictionaryIsUpdatedWithAssignedValueWhenDecimalExponentValueReadFromStream(t *testing.T) {
	// Arrange exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue(int32(2))}
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("DecimalFieldExponent")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<decimal presence="optional">
//	<exponent />
//	<mantissa />
//</decimal>
func TestDictionaryIsUpdatedWithEmptyValueWhenDecimalExponentValueIsNil(t *testing.T) {
	// Arrange exp = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("DecimalFieldExponent")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<decimal>
//	<exponent />
//	<mantissa />
//</decimal>
func TestDictionaryIsUpdatedWithAssignedValueWhenDecimalMantissaValueReadFromStream(t *testing.T) {
	// Arrange exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue(int64(1))}
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("DecimalFieldMantissa")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<decimal>
//	<exponent />
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal presence="optional">
//	<exponent />
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsFalseForOptionalDecimalNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
