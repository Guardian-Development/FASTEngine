package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<uInt64 />
func TestCanDeseraliseRequiredUInt64(t *testing.T) {
	// Arrange 3 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(3)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: true,
		},
		Operation: operation.None{},
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

//<uInt64 presence="optional"/>
func TestCanDeseraliseOptionalUInt64Present(t *testing.T) {
	// Arrange 3 = 10000100
	messageAsBytes := bytes.NewBuffer([]byte{132})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(3)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.None{},
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

//<uInt64 presence="optional"/>
func TestCanDeseraliseOptionalUInt64Null(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.None{},
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

//<uInt64>
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseRequiredUInt64ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(132)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
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

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
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

//<uInt64 presence="optional">
//	<constant value="132" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := uint64(132)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint64(132),
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

//<uInt64 />
func TestRequiresPmapReturnsFalseForRequiredUInt64NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
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

//<uInt64 presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalUInt64NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
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

//<uInt64>
//	<constant value="132" />
//</uInt64>
func TestRequiresPmapReturnsFalseForRequiredUInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
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
			Name:     "UInt64Field",
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

//<uInt64>
//	<default value="5" />
//</uInt64>
func TestCanDeseraliseUInt64DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := uint64(2)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: uint64(5),
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

//<uInt64>
//	<default value="5" />
//</uInt64>
func TestCanDeseraliseUInt64DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(5)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: uint64(5),
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

//<uInt64 presence="optional">
//	<default value="5" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := uint64(2)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: uint64(5),
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

//<uInt64 presence="optional">
//	<default value="5" />
//</uInt64>
func TestCanDeseraliseOptionalUInt64DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint64(5)
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: uint64(5),
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

//<uInt64 presence="optional">
//	<default />
//</uInt64>
func TestCanDeseraliseOptionalUInt64DefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: nil,
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

//<uInt64>
//	<default value="132" />
//</uInt64>
func TestRequiresPmapReturnsTrueForRequiredUInt64DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: uint64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uInt64 presence="optional">
//	<default value="132" />
//</uInt64>
func TestRequiresPmapReturnsTrueForOptionalUInt64DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt64{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: uint64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
