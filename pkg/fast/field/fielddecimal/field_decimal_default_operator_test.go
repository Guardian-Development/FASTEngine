package fielddecimal

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"
)

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", true), 2),
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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", true), 2),
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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 2))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 10))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", true), 2),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 2))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", true), 2),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 2))

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
//		<default />
//	</exponent>
//	<mantissa />
//</decimal>
func TestCanDeseraliseOptionalDecimalExponentDefaultOperatorNotEncodedReturnsNilValueAndDoesNotReadMantissa(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDefaultOperation(properties.New(1, "DecimalFieldExponent", false)),
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", false), 2),
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", false), 2),
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
//		<default value="12" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsTrueForRequiredDecimalExponentDefault(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", true), 12),
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
//		<default value="12" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForRequiredDecimalMantissaDefault(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 12))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", true), 12),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 12))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", false), 7),
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
//		<default value="2" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsTrueForOptionalDecimalMantissaDefault(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 2))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldExponent", false), 2),
		fieldint64.NewDefaultOperationWithValue(properties.New(1, "DecimalFieldMantissa", true), 2))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
