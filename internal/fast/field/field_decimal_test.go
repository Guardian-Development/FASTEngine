package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<decimal>
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimal(t *testing.T) {
	// Arrange exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentPresent(t *testing.T) {
	// Arrange exp = 10000011 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{131, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//	<exponent />
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentPresentMantissaNotEncodedCausesError(t *testing.T) {
	// Arrange exp = 10000011 man = nil
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}
	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || err.Error() != "unable to decode mantissa after successful decoding of exponent" {
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
	dictionary := dictionary.New()
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//	<exponent>
//		<constant value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentConstantOperatorNotEncoded(t *testing.T) {
	// Arrange man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int32(2)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<default value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentDefaultOperatorEncodedReadsFromStream(t *testing.T) {
	// Arrange pmap = 11000000 exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int32(2),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<default value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentDefaultOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int32(2),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
// 		<constant value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaConstantOperatorNotEncoded(t *testing.T) {
	// Arrange exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int64(2)},
		},
	}

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
// 		<default value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaDefaultOperatorEncodedReadsFromStream(t *testing.T) {
	// Arrange pmap = 11000000 exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(2),
			},
		},
	}

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
// 		<default value="10" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaDefaultOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000 exp = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(10),
			},
		},
	}

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
// 		<constant value="2" />
//	</exponent>
//	<mantissa>
// 		<constant value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentAndMantissaConstantOperatorNotEncoded(t *testing.T) {
	// Arrange
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int32(2)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int64(2)},
		},
	}

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
// 		<default value="2" />
//	</exponent>
//	<mantissa>
// 		<default value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentAndMantissaDefaultOperatorEncodedReadsFromStream(t *testing.T) {
	// Arrange pmap = 11100000 exp = 10000010 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{130, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{224}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int32(2),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(2),
			},
		},
	}

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
// 		<default value="2" />
//	</exponent>
//	<mantissa>
// 		<default value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalExponentAndMantissaDefaultOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int32(2),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(2),
			},
		},
	}

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
//		<constant value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentConstantOperatorNotEncodedReturnsNilValueAndDoesNotReadMantissa(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Constant{ConstantValue: int32(2)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<default />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentDefaultOperatorNotEncodedReturnsNilValueAndDoesNotReadMantissa(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Default{
				DefaultValue: nil,
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<default value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentDefaultOperatorEncodedReadsExponentAndMantissa(t *testing.T) {
	// Arrange pmap = 11000000 exp = 10000011 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{131, 129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Default{
				DefaultValue: int32(2),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<default value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentDefaultOperatorNotEncodedReadsDefaultExponentAndMantissaFromStream(t *testing.T) {
	// Arrange pmap = 10000000 man = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(100)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Default{
				DefaultValue: int32(2),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<constant value="2" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentConstantOperatorEncodedReadsMantissa(t *testing.T) {
	// Arrange pmap = 11000000 man = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Constant{ConstantValue: int32(2)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal>
//	<exponent>
//		<default value="12" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsTrueForRequiredDecimalExponentDefault(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int32(12),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal>
//	<exponent>
//		<constant value="12" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalExponentConstant(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int32(12)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal>
//	<exponent/>
//	<mantissa>
//		<default value="12" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForRequiredDecimalMantissaDefault(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(12),
			},
		},
	}

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
//		<constant value="12" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalMantissaConstant(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int64(12)},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal>
//	<exponent>
//		<default value="12" />
//	</exponent>
//	<mantissa>
//		<default value="12" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForRequirdDecimalExponentAndMantissaDefault(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int32(12),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(12),
			},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal>
//	<exponent>
//		<constant value="12" />
//	</exponent>
//	<mantissa>
//		<constant value="12" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsFalseForRequirdDecimalExponentAndMantissaConstant(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: true,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int32(12)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int64(12)},
		},
	}

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
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal presence="optional">
//	<exponent>
//		<default value="7" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalExponentDefault(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Default{
				DefaultValue: int32(7),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal presence="optional">
//	<exponent>
//		<constant value="7" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalExponentConstant(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Constant{ConstantValue: int32(7)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.None{},
		},
	}

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
//		<default value="2" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalMantissaDefault(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(2),
			},
		},
	}

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
//		<constant value="2" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsFalseForOptionalDecimalMantissaConstant(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.None{},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int64(7)},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<decimal presence="optional">
//	<exponent>
//		<default value="2" />
//	</exponent>
//	<mantissa>
//		<default value="2" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalExponentAndRequiredMantissaDefault(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Default{
				DefaultValue: int32(7),
			},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: int64(7),
			},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<decimal presence="optional">
//	<exponent>
//		<constant value="2" />
//	</exponent>
//	<mantissa>
//		<constant value="2" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalExponentAndRequiredMantissaConstant(t *testing.T) {
	// Arrange
	unitUnderTest := Decimal{
		FieldDetails: Field{
			ID:       1,
			Name:     "DecimalField",
			Required: false,
		},
		ExponentField: Int32{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldExponent",
				Required: false,
			},
			Operation: operation.Constant{ConstantValue: int32(7)},
		},
		MantissaField: Int64{
			FieldDetails: Field{
				ID:       1,
				Name:     "DecimalFieldMantissa",
				Required: true,
			},
			Operation: operation.Constant{ConstantValue: int64(7)},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}