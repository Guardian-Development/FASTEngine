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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", true), 2),
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
// 		<constant value="2" />
//	</mantissa>
//</decimal>
func TestCanDeseraliseRequiredDecimalMantissaConstantOperatorNotEncoded(t *testing.T) {
	// Arrange exp = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := float64(200)
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewConstantOperation(properties.New(1, "DecimalFieldMantissa", true), 2))

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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", true), 2),
		fieldint64.NewConstantOperation(properties.New(1, "DecimalFieldMantissa", true), 2))

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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", false), 2),
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", false), 2),
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
//		<constant value="12" />
//	</exponent>
//	<mantissa />
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalExponentConstant(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", true), 12),
		fieldint64.New(properties.New(1, "DecimalFieldMantissa", true)))

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
//		<constant value="12" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsFalseForRequiredDecimalMantissaConstant(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", true)),
		fieldint64.NewConstantOperation(properties.New(1, "DecimalFieldMantissa", true), 12))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
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
	unitUnderTest := New(properties.New(1, "DecimalField", true),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", true), 12),
		fieldint64.NewConstantOperation(properties.New(1, "DecimalFieldMantissa", true), 12))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", false), 7),
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
//		<constant value="2" />
//	</mantissa>
//</decimal>
func TestRequiresPmapReturnsFalseForOptionalDecimalMantissaConstant(t *testing.T) {
	// Arrange
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.New(properties.New(1, "DecimalFieldExponent", false)),
		fieldint64.NewConstantOperation(properties.New(1, "DecimalFieldMantissa", true), 2))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
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
	unitUnderTest := New(properties.New(1, "DecimalField", false),
		fieldint32.NewConstantOperation(properties.New(1, "DecimalFieldExponent", false), 2),
		fieldint64.NewConstantOperation(properties.New(1, "DecimalFieldMantissa", true), 2))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
