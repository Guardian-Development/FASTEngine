package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<int32 />
func TestCanDeseraliseRequiredPositiveInt32(t *testing.T) {
	// Arrange 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 />
func TestCanDeseraliseRequiredNegativeInt32(t *testing.T) {
	// Arrange -2 = 11111110
	messageAsBytes := bytes.NewBuffer([]byte{254})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(-2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 presence="optional"/>
func TestCanDeseraliseOptionalPositiveInt32Present(t *testing.T) {
	// Arrange 3 = 10000100
	messageAsBytes := bytes.NewBuffer([]byte{132})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(3)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 presence="optional"/>
func TestCanDeseraliseOptionalNegativeInt32Present(t *testing.T) {
	// Arrange 3 = 11111101
	messageAsBytes := bytes.NewBuffer([]byte{253})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(-3)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 presence="optional"/>
func TestCanDeseraliseOptionalInt32Null(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

// <int32 />
func TestDictionaryIsUpdatedWithAssignedValueWhenInt32ValueReadFromStream(t *testing.T) {
	// Arrange 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue(int32(2))}
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("Int32Field")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <int32 presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenInt32NilValueReadFromStream(t *testing.T) {
	// Arrange 2 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("Int32Field")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<int32>
//	<constant value="132" />
//</int32>
func TestCanDeseraliseRequiredInt32ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(132)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
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

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestCanDeseraliseOptionalInt32ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
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

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestCanDeseraliseOptionalInt32ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := int32(132)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
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

//<int32 />
func TestRequiresPmapReturnsFalseForRequiredInt32NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalInt32NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32>
//	<constant value="132" />
//</int32>
func TestRequiresPmapReturnsFalseForRequiredInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForOptionalInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int32>
//	<default value="5" />
//</int32>
func TestCanDeseraliseInt32DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32>
//	<default value="5" />
//</int32>
func TestCanDeseraliseInt32DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(5)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32 presence="optional">
//	<default value="5" />
//</int32>
func TestCanDeseraliseOptionalInt32DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32 presence="optional">
//	<default value="5" />
//</int32>
func TestCanDeseraliseOptionalInt32DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(5)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32 presence="optional">
//	<default />
//</int32>
func TestCanDeseraliseOptionalInt32DefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32>
//	<default value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForRequiredInt32DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int32 presence="optional">
//	<default value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForOptionalInt32DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int32(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int32>
//	<copy />
//</int32>
func TestCanDeseraliseRequiredInt32CopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := int32(2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32>
//	<copy value="12"/>
//</int32>
func TestCanDeseraliseRequiredInt32CopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int32(12)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: int32(12),
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

//<int32>
//	<copy value="15"/>
//</int32>
func TestCanDeseraliseRequiredInt32CopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int32(7)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: int32(15),
		},
	}

	// Act
	dict.SetValue("Int32Field", fix.NewRawValue(int32(7)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int32>
//	<copy />
//</int32>
func TestCanDeseraliseRequiredInt32CopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 presence="optional">
//	<copy />
//</int32>
func TestCanDeseraliseOptionalInt32CopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
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

//<int32 presence="optional">
//	<copy value="12"/>
//</int32>
func TestCanDeseraliseOptionalInt32CopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	dict.SetValue("Int32Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<int32>
//	<copy value="1"/>
//</int32>
func TestRequiresPmapReturnsTrueForRequiredInt32CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: int32(1),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int32 presence="optional">
//	<copy value="7"/>
//</int32>
func TestRequiresPmapReturnsTrueForOptionalInt32CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int32Field",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: int32(1),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
