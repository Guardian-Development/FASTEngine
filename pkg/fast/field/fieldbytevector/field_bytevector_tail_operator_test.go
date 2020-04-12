package fieldbytevector

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<byteVector>
//	<tail value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorNotEncodedNoPreviousValueReturnsInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<tail value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorNotEncodedNoPreviousValueReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA2, 0xB3, 0xC1}
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	dict.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA2, 0xB3, 0xC1}))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector presence="optional">
//	<tail value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseOptionlByteVectorTailOperatorNotEncodedEmptyPreviousValueReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	dict.SetValue("ByteVectorField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<byteVector>
//	<tail value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorEncodedReturnsInitialValueCombinedWithValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA1, 0x92, 0xAA}
	unitUnderTest := NewTailOperationWithInitialValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<tail />
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := []byte{0x92, 0xAA}
	unitUnderTest := NewTailOperation(properties.New(1, "ByteVectorField", true, testLog))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<tail />
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorEncodedPreviousValueReturnsPreviousValueCombinedWithValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA1, 0x92, 0xAA}
	unitUnderTest := NewTailOperation(properties.New(1, "ByteVectorField", true, testLog))

	// Act
	dict.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA1, 0xB2, 0xCF}))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<tail />
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorEncodedPreviousValueReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := []byte{0x92, 0xAA}
	unitUnderTest := NewTailOperation(properties.New(1, "ByteVectorField", true, testLog))

	// Act
	dict.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xAA}))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<tail />
//</byteVector>
func TestCanDeseraliseByteVectorTailOperatorEncodedPreviousValueEmptyReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedMessage := []byte{0x92, 0xAA}
	unitUnderTest := NewTailOperation(properties.New(1, "ByteVectorField", true, testLog))

	// Act
	dict.SetValue("ByteVectorField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<tail />
//</byteVector>
func TestRequiresPmapReturnsTrueForRequiredByteVectorTailOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewTailOperation(properties.New(1, "ByteVectorField", true, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<byteVector presence="optional">
//	<tail />
//</byteVector>
func TestRequiresPmapReturnsTrueForOptionalByteVectorTailOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewTailOperation(properties.New(1, "ByteVectorField", false, testLog))

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
