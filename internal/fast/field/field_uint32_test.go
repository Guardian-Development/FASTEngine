package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<uInt32 />
func TestCanDeseraliseRequiredUInt32(t *testing.T) {
	// Arrange 3 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(3)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32 presence="optional"/>
func TestCanDeseraliseOptionalUInt32Present(t *testing.T) {
	// Arrange 3 = 10000100
	messageAsBytes := bytes.NewBuffer([]byte{132})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(3)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32 presence="optional"/>
func TestCanDeseraliseOptionalUInt32Null(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

// <uInt32 />
func TestDictionaryIsUpdatedWithAssignedValueWhenUInt32ValueReadFromStream(t *testing.T) {
	// Arrange 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue(uint32(2))}
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("UInt32Field")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <uInt32 presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenUInt32NilValueReadFromStream(t *testing.T) {
	// Arrange 2 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("UInt32Field")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<uInt32>
//	<constant value="132" />
//</uInt32>
func TestCanDeseraliseRequiredUInt32ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(132)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: uint32(132),
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

//<uInt32 presence="optional">
//	<constant value="132" />
//</uInt32>
func TestCanDeseraliseOptionalUInt32ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint32(132),
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

//<uInt32 presence="optional">
//	<constant value="132" />
//</uInt32>
func TestCanDeseraliseOptionalUInt32ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := uint32(132)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint32(132),
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

//<uInt32 />
func TestRequiresPmapReturnsFalseForRequiredUInt32NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32 presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalUInt32NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32>
//	<constant value="132" />
//</uInt32>
func TestRequiresPmapReturnsFalseForRequiredUInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: uint32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<uInt32 presence="optional">
//	<constant value="132" />
//</uInt32>
func TestRequiresPmapReturnsTrueForOptionalUInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: uint32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uInt32>
//	<default value="5" />
//</uInt32>
func TestCanDeseraliseUInt32DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := uint32(2)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: uint32(5),
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

//<uInt32>
//	<default value="5" />
//</uInt32>
func TestCanDeseraliseUInt32DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(5)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: uint32(5),
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

//<uInt32 presence="optional">
//	<default value="5" />
//</uInt32>
func TestCanDeseraliseOptionalUInt32DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := uint32(2)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: uint32(5),
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

//<uInt32 presence="optional">
//	<default value="5" />
//</uInt32>
func TestCanDeseraliseOptionalUInt32DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(5)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: uint32(5),
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

//<uInt32 presence="optional">
//	<default />
//</uInt32>
func TestCanDeseraliseOptionalUInt32DefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32>
//	<default value="132" />
//</uInt32>
func TestRequiresPmapReturnsTrueForRequiredUInt32DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: uint32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uInt32 presence="optional">
//	<default value="132" />
//</uInt32>
func TestRequiresPmapReturnsTrueForOptionalUInt32DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: uint32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uInt32>
//	<copy />
//</uInt32>
func TestCanDeseraliseRequiredUInt32CopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := uint32(2)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: nil,
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

//<uInt32>
//	<copy value="12"/>
//</uInt32>
func TestCanDeseraliseRequiredUInt32CopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := uint32(12)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: uint32(12),
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

//<uInt32>
//	<copy value="15"/>
//</uInt32>
func TestCanDeseraliseRequiredUInt32CopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := uint32(7)
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: uint32(15),
		},
	}

	// Act
	dict.SetValue("UInt32Field", fix.NewRawValue(uint32(7)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<uInt32>
//	<copy />
//</uInt32>
func TestCanDeseraliseRequiredUInt32CopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32 presence="optional">
//	<copy />
//</uInt32>
func TestCanDeseraliseOptionalUInt32CopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
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

//<uInt32 presence="optional">
//	<copy value="12"/>
//</uInt32>
func TestCanDeseraliseOptionalUInt32CopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	dict.SetValue("UInt32Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<uInt32>
//	<copy value="1"/>
//</uInt32>
func TestRequiresPmapReturnsTrueForRequiredUInt32CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: uint32(1),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<uInt32 presence="optional">
//	<copy value="7"/>
//</uInt32>
func TestRequiresPmapReturnsTrueForOptionalUInt32CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UInt32{
		FieldDetails: Field{
			ID:       1,
			Name:     "UInt32Field",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: uint32(1),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
