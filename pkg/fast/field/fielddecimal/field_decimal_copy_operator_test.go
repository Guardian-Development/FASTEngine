package fielddecimal

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
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
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
//	<exponent>
//		<copy value="2"/>
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentCopyOperationNotEncodedTakesInitialValue(t *testing.T) {
	// Arrange pmap = 10000000 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", true), 2),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
//	<exponent>
//		<copy />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentCopyOperationNotEncodedNoInitialValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "no value supplied in message and no initial value with required field" {
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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", true), 2),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true)))

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
//	<exponent />
//	<mantissa>
// 		<copy value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaCopyOperationNotEncodedTakesInitialValue(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldMantissa", true), 2))

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
//	<exponent />
//	<mantissa>
// 		<copy />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaCopyOperationNotEncodedNoInitialValueReturnsError(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "unable to decode mantissa after successful decoding of exponent: no value supplied in message and no initial value with required field" {
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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldMantissa", true), 2))

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
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", false), 2),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperationWithInitialValue(properties.New(1, "DecimalFieldExponent", false), 2),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true)))

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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewCopyOperation(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewCopyOperation(properties.New(1, "DecimalFieldMantissa", true)))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
