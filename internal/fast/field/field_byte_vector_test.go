package field

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<byteVector />
func TestCanDeseraliseRequiredByteVector(t *testing.T) {
	// Arrange 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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
	dictionary := dictionary.New()
	expectedMessage := []byte{}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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
	dictionary := dictionary.New()
	expectedMessage := []byte{146}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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
	dictionary := dictionary.New()
	expectedMessage := []byte{}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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
	dictionary := dictionary.New()
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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

// <byteVector />
func TestDictionaryIsUpdatedWithAssignedValueWhenByteVectorValueReadFromStream(t *testing.T) {
	// Arrange 10000001 10010010
	messageAsBytes := bytes.NewBuffer([]byte{129, 146})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue([]byte{146})}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("ByteVectorField")
	areEqual := reflect.DeepEqual(expectedValue, result)
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <byteVector presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenByteVectorNilValueReadFromStream(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("ByteVectorField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<byteVector>
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseRequiredByteVectorConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
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

//<byteVector presence="optional">
//	<constant value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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

//<byteVector />
func TestRequiresPmapReturnsFalseForRequiredByteVectorNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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
			Name:     "ByteVectorField",
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
			Name:     "ByteVectorField",
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
			Name:     "ByteVectorField",
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

//<byteVector>
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseByteVectorDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseByteVectorDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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

//<byteVector presence="optional">
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 value = 10000011 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{131, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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

//<byteVector presence="optional">
//	<default value="A1 B2 CF" />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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

//<byteVector presence="optional">
//	<default />
//</byteVector>
func TestCanDeseraliseOptionalByteVectorDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
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

//<byteVector>
//	<default value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsTrueForRequiredByteVectorDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<byteVector presence="optional">
//	<default value="A1 B2 CF" />
//</byteVector>
func TestRequiresPmapReturnsTrueForOptionalByteVectorDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<byteVector>
//	<copy />
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000, value = 11000010 10010010 10101010
	messageAsBytes := bytes.NewBuffer([]byte{130, 146, 170})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{198}))
	dictionary := dictionary.New()
	expectedMessage := []byte{146, 170}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Copy{},
	}

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
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := []byte{0xA1, 0xB2, 0xCF}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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
//	<copy value="A1 B2 CF"/>
//</byteVector>
func TestCanDeseraliseRequiredByteVectorCopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := []byte{0xA2, 0xB3}
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	_, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	if err == nil || err.Error() != "no value supplied in message and no initial value with required field" {
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
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

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
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

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
	unitUnderTest := ByteVector{
		FieldDetails: Field{
			ID:       1,
			Name:     "ByteVectorField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: []byte{0xA1, 0xB2, 0xCF},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
