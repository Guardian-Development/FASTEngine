package fieldbytevector

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"reflect"
	"strings"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorAppendToBaseValue(t *testing.T) {
	// Arrange length = 1000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{128, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorPrependToBaseValue(t *testing.T) {
	// Arrange length = 1111111 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{255, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<delta value="A1 B2 CF"/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorAppendToInitialValue(t *testing.T) {
	// Arrange length = 1000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{128, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF, 146, 170}
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "ByteVectorField", true), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<delta value="A1 B2 CF""/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorPrependToInitialValue(t *testing.T) {
	// Arrange length = 11111111 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{255, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170, 0xA1, 0xB2, 0xCF}
	unitUnderTest := NewDeltaOperationWithInitialValue(properties.New(1, "ByteVectorField", true), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorAppendWithOverwiteToPreviousValue(t *testing.T) {
	// Arrange length = 10000100 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{132, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF, 146, 170}
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	dictionary.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA1, 0xB2, 0xCF, 111, 222, 223, 224}))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorPrependWithOverwriteToPreviousValue(t *testing.T) {
	// Arrange length = 11111010 (-6) value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{250, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170, 223, 224}
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	dictionary.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA1, 0xB2, 0xCF, 111, 222, 223, 224}))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorAppendWithOverwiteTooLargeReturnsError(t *testing.T) {
	// Arrange length = 10000100 (4) value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{132, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	dictionary.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA1, 0xB2, 0xCF}))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D7) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiredByteVectorDeltaOperatorPrependWithOverwiteTooLargeReturnsError(t *testing.T) {
	// Arrange length = 11111011 (-5) value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{251, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	dictionary.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA1, 0xB2, 0xCF}))
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D7) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<byteVector presence="optional">
//	<delta/>
//</byteVector>
func TestOptionalByteVectorDeltaOperatorNotEncodedReturnsNull(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", false))

	// Act
	dictionary.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA1, 0xB2, 0xCF}))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<byteVector presence="optional">
//	<delta/>
//</byteVector>
func TestOptionalByteVectorDeltaOperatorEncodedPreviousNullValueReturnsError(t *testing.T) {
	// Arrange length = 10000001 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{129, 130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", false))

	// Act
	dictionary.SetValue("ByteVectorField", fix.NullValue{})
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D6) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<byteVector>
//	<delta/>
//</byteVector>
func TestRequiresPmapReturnsFalseForRequiredByteVectorDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", true))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<byteVector presence="optional">
//	<delta/>
//</byteVector>
func TestRequiresPmapReturnsFalseForOptionalByteVectorDeltaOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewDeltaOperation(properties.New(1, "ByteVectorField", false))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}
