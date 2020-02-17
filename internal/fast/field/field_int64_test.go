package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<int64 />
func TestCanDeseraliseRequiredPositiveInt64(t *testing.T) {
	// Arrange 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 />
func TestCanDeseraliseRequiredNegativeInt64(t *testing.T) {
	// Arrange -2 = 11111110
	messageAsBytes := bytes.NewBuffer([]byte{254})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(-2)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 presence="optional"/>
func TestCanDeseraliseOptionalPositiveInt64Present(t *testing.T) {
	// Arrange 3 = 10000100
	messageAsBytes := bytes.NewBuffer([]byte{132})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(3)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 presence="optional"/>
func TestCanDeseraliseOptionalNegativeInt64Present(t *testing.T) {
	// Arrange 3 = 11111101
	messageAsBytes := bytes.NewBuffer([]byte{253})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(-3)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 presence="optional"/>
func TestCanDeseraliseOptionalInt64Null(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

// <int64 />
func TestDictionaryIsUpdatedWithAssignedValueWhenInt64ValueReadFromStream(t *testing.T) {
	// Arrange 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue(int64(2))}
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("Int64Field")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <int64 presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenInt64NilValueReadFromStream(t *testing.T) {
	// Arrange 2 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("Int64Field")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<int64>
//	<constant value="132" />
//</int64>
func TestCanDeseraliseRequiredInt64ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(132)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: int64(132),
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

//<int64 presence="optional">
//	<constant value="132" />
//</int64>
func TestCanDeseraliseOptionalInt64ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int64(132),
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

//<int64 presence="optional">
//	<constant value="132" />
//</int64>
func TestCanDeseraliseOptionalInt64ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := int64(132)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int64(132),
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

//<int64 />
func TestRequiresPmapReturnsFalseForRequiredInt64NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalInt64NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64>
//	<constant value="132" />
//</int64>
func TestRequiresPmapReturnsFalseForRequiredInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: int64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<int64 presence="optional">
//	<constant value="132" />
//</int64>
func TestRequiresPmapReturnsTrueForOptionalInt64ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int64>
//	<default value="6" />
//</int64>
func TestCanDeseraliseInt64DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int64(6),
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

//<int64>
//	<default value="6" />
//</int64>
func TestCanDeseraliseInt64DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(6)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int64(6),
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

//<int64 presence="optional">
//	<default value="132" />
//</int64>
func TestCanDeseraliseOptionalInt64DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int64(6),
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

//<int64 presence="optional">
//	<default value="132" />
//</int64>
func TestCanDeseraliseOptionalInt64DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(6)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int64(6),
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

//<int64 presence="optional">
//	<default />
//</int64>
func TestCanDeseraliseOptionalInt64DefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64>
//	<default value="132" />
//</int64>
func TestRequiresPmapReturnsTrueForRequiredInt64DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int64 presence="optional">
//	<default value="132" />
//</int64>
func TestRequiresPmapReturnsTrueForOptionalInt64DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int64(132),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int64>
//	<copy />
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := int64(2)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64>
//	<copy value="12"/>
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := int64(12)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: int64(12),
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

//<int64>
//	<copy value="15"/>
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := int64(7)
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: int64(15),
		},
	}

	// Act
	dict.SetValue("Int64Field", fix.NewRawValue(int64(7)))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<int64>
//	<copy />
//</int64>
func TestCanDeseraliseRequiredInt64CopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 presence="optional">
//	<copy />
//</int64>
func TestCanDeseraliseOptionalInt64CopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
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

//<int64 presence="optional">
//	<copy value="12"/>
//</int64>
func TestCanDeseraliseOptionalInt64CopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	dict.SetValue("Int64Field", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<int64>
//	<copy value="1"/>
//</int64>
func TestRequiresPmapReturnsTrueForRequiredInt64CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: int64(1),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<int64 presence="optional">
//	<copy value="7"/>
//</int64>
func TestRequiresPmapReturnsTrueForOptionalInt64CopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int64{
		FieldDetails: Field{
			ID:       1,
			Name:     "Int64Field",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: int64(1),
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
