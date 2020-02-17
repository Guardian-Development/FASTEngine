package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<string />
func TestCanDeseraliseRequiredAsciiString(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string presence="optional"/>
func TestCanDeseraliseOptionalAsciiStringPresent(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string presence="optional"/>
func TestCanDeseraliseOptionalAsciiStringNull(t *testing.T) {
	// Arrange TEST1 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

// <string />
func TestDictionaryIsUpdatedWithAssignedValueWhenAsciiStringValueReadFromStream(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue("TEST1")}
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("AsciiStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

// <string presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenAsciiStringNilValueReadFromStream(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedValue := dictionary.EmptyValue{}
	dict := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("AsciiStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<string>
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseRequiredAsciiStringConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
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

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
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

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
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

//<string />
func TestRequiresPmapReturnsFalseForRequiredAsciiStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalAsciiStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string>
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsFalseForRequiredAsciiStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string>
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseAsciiStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string>
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseAsciiStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string presence="optional">
//	<default />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string>
//	<default value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForRequiredAsciiStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string presence="optional">
//	<default />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string>
//	<copy />
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string>
//	<copy value="TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: "TEST1",
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

//<string>
//	<copy value="TEST1"/>
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "PREVIOUS_VALUE"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: "TEST1",
		},
	}

	// Act
	dict.SetValue("AsciiStringField", fix.NewRawValue("PREVIOUS_VALUE"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string>
//	<copy />
//</string>
func TestCanDeseraliseAsciiStringCopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string presence="optional">
//	<copy />
//</string>
func TestCanDeseraliseOptionalAsciiStringCopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
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

//<string presence="optional">
//	<copy />
//</string>
func TestCanDeseraliseOptionalAsciiStringCopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	dict.SetValue("AsciiStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<string>
//	<copy />
//</string>
func TestRequiresPmapReturnsTrueForRequiredAsciiStringCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: "TEST2",
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string presence="optional">
//	<copy />
//</string>
func TestRequiresPmapReturnsTrueForOptionalAsciiStringCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Name:     "AsciiStringField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: "TEST2",
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
