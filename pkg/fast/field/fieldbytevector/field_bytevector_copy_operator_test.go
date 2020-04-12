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
//	<copy />
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000, value = 11000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{198}))
	dict := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := NewCopyOperation(properties.New(1, "ByteVectorField", true, testLog))

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
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

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
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA2, 0xB3}
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	dict.SetValue("ByteVectorField", fix.NewRawValue([]byte{0xA2, 0xB3}))
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
//	<copy />
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "ByteVectorField", true, testLog))

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.D5) {
		t.Errorf("Expected error about nil value when a required field: %#v", err)
	}
}

//<byteVector presence="optional">
//	<copy />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorCopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperation(properties.New(1, "ByteVectorField", false, testLog))

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<byteVector presence="optional">
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseOptionalByteVectorCopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

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
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestRequiresPmapReturnsTrueForRequiredByteVectorCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "ByteVectorField", true, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<byteVector presence="optional">
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestRequiresPmapReturnsTrueForOptionalByteVectorCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := NewCopyOperationWithInitialValue(properties.New(1, "ByteVectorField", false, testLog), []byte{0xA1, 0xB2, 0xCF})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
