package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<uInt64 />
func TestCanDeseraliseRequiredUInt64(t *testing.T) {
	// Arrange 3 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := uint64(3)
	unitUnderTest := UInt64{
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
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uInt64 presence="optional"/>
func TestCanDeseraliseOptionalUInt64Present(t *testing.T) {
	// Arrange 3 = 10000100
	messageAsBytes := bytes.NewBuffer([]byte{132})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := uint64(3)
	unitUnderTest := UInt64{
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
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uInt64 presence="optional"/>
func TestCanDeseraliseOptionalUInt64Null(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := UInt64{
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

//<uInt64>
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseRequiredUInt64ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := uint64(132)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
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

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	expectedMessage := uint64(132)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uInt64>
//	<constant value="132" />
//</uInt64>
func TestRequiresPmapReturnsFalseForRequiredUInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestRequiresPmapReturnsTrueForOptionalUInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
