package fielddecimal

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<decimal>
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentCopyOperationEncodedReadsFromStream(t *testing.T) {
	// Arrange pmap = 11000000 exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true, testLog)),
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

//<decimal>
//	<exponent>
//		<copy value="2"/>
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentCopyOperationNotEncodedTakesInitialValue(t *testing.T) {
	// Arrange pmap = 10000000 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", true, testLog), 2),
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

//<decimal>
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentCopyOperationNotEncodedNoInitialValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal>
//	<exponent>
//		<copy value="2"/>
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentCopyOperationNotEncodedReadsPreviousValueThenMantissaFromStream(t *testing.T) {
	// Arrange pmap = 10000000 man = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", true, testLog), 2),
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

//<decimal>
//	<exponent />
//	<mantissa>
// 		<copy />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaCopyOperationEncodedReadsFromStream(t *testing.T) {
	// Arrange pmap = 11000000 exp = 10000010 man = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130, 130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true, testLog)))

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

//<decimal>
//	<exponent />
//	<mantissa>
// 		<copy value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaCopyOperationNotEncodedTakesInitialValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldMantissa", true, testLog), 2))

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

//<decimal>
//	<exponent />
//	<mantissa>
// 		<copy />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaCopyOperationNotEncodedNoInitialValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<decimal>
//	<exponent />
//	<mantissa>
// 		<copy value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaCopyOperationNotEncodedReadsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(400)
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldMantissa", true, testLog), 2))

	// Act
	dict.SetValue("DecimalFieldMantissa", fix.NewRawValue(int64(4)))
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
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentCopyOperatorEncodedReadsExponentAndMantissa(t *testing.T) {
	// Arrange pmap = 11000000 exp = 10000011 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{131, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false, testLog)),
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
//	<exponent>
//		<copy value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentCopyOperatorNotEncodedReadsInitialValueThenMantissaFromStream(t *testing.T) {
	// Arrange pmap = 10000000 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", false, testLog), 2),
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
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentCopyOperatorNotEncodedNoInitialValueReturnsNilAndDoesNotReadMantissa(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false, testLog)),
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

//<decimal presence="optional">
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentCopyOperatorNotEncodedPreviousValueEmptyReturnsNilAndDoesNotReadMantissa(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	dict.SetValue("DecimalFieldExponent", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<decimal presence="optional">
//	<exponent>
//		<copy value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentCopyOperatorNotEncodedPreviousValuePresentReadsMantissa(t *testing.T) {
	// Arrange pmap = 10000000 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", false, testLog), 2),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	dict.SetValue("DecimalFieldExponent", fix.NewRawValue(int32(2)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<decimal>
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsTrueForRequiredDecimalExponentCopy(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal>
//	<exponent/>
//	<mantissa>
//		<copy />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForRequiredDecimalMantissaCopy(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal>
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa>
//		<copy />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForRequirdDecimalExponentAndMantissaCopy(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true, testLog)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal presence="optional">
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalExponentCopy(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal presence="optional">
//	<exponent/>
//	<mantissa>
//		<copy />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalMantissaCopy(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal presence="optional">
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa>
//		<copy />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalExponentAndRequiredMantissaCopy(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false, testLog),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false, testLog)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true, testLog)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
