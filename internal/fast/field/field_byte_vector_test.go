package field

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<byteVector />
func TestCanDeseraliseRequiredByteVector(t *testing.T) {
	// Arrange 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := []byte{146, 170}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector />
func TestCanDeseraliseRequiredByteVectorLengthZero(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := []byte{}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector presence="optional"/>
func TestCanDeseraliseOptionalByteVectorPresent(t *testing.T) {
	// Arrange 10000010 10010010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := []byte{146}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector presence="optional"/>
func TestCanDeseraliseOptionalByteVectorPresentLengthZero(t *testing.T) {
	// Arrange 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := []byte{}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector presence="optional"/>
func TestCanDeseraliseOptionalByteVectorNull(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<byteVector>
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseRequiredByteVectorConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result.Get())
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<byteVector />
func TestRequiresPmapReturnsFalseForRequiredByteVectorNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<byteVector presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalByteVectorNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<byteVector>
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsFalseForRequiredByteVectorConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsTrueForOptionalByteVectorConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
