package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<string charset="unicode"/>
func TestCanDeseraliseRequiredUnicodeString(t *testing.T) {
	// Arrange TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional"/>
func TestCanDeseraliseOptionalUnicodeStringPresent(t *testing.T) {
	// Arrange TEST1 = 10000110 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{134, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional"/>
func TestCanDeseraliseOptionalUnicodeStringNull(t *testing.T) {
	// Arrange TEST1 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" />
func TestDictionaryIsUpdatedWithAssignedValueWhenUnicodeStringValueReadFromStream(t *testing.T) {
	// Arrange TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue("TEST1")}
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: true,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("UnicodeStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<string charset="unicode" presence="optional"/>
func TestDictionaryIsUpdatedWithEmptyValueWhenUnicodeStringNilValueReadFromStream(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: false,
		},
		Operation: operation.None{},
	}

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("UnicodeStringField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<string charset="unicode">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseRequiredUnicodeStringConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalUnicodeStringConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalUnicodeStringConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode"/>
func TestRequiresPmapReturnsFalseForRequiredUnicodeStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalUnicodeStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode">
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsFalseForRequiredUnicodeStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForOptionalUnicodeStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseUnicodeStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseUnicodeStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalUnicodeStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 10000110 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{134, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalUnicodeStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<default />
//</string>
func TestCanDeseraliseOptionalUnicodeStringDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode">
//	<default value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForRequiredUnicodeStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<default />
//</string>
func TestRequiresPmapReturnsTrueForOptionalUnicodeStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode">
//	<copy />
//</string>
func TestCanDeseraliseUnicodeStringCopyOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode">
//	<copy value="TEST2"/>
//</string>
func TestCanDeseraliseUnicodeStringCopyOperatorNotEncodedReturnsInitialValueIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: "TEST2",
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

//<string charset="unicode">
//	<copy value="TEST2"/>
//</string>
func TestCanDeseraliseUnicodeStringCopyOperatorNotEncodedReturnsPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "PREVIOUS_VALUE"
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: "TEST2",
		},
	}

	// Act
	dict.SetValue("UnicodeStringField", fix.NewRawValue("PREVIOUS_VALUE"))
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != expectedMessage {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<string charset="unicode">
//	<copy />
//</string>
func TestCanDeseraliseUnicodeStringCopyOperatorNotEncodedReturnsErrorIfNoPreviousValueOrInitialValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<copy />
//</string>
func TestCanDeseraliseOptionalUnicodeStringCopyOperatorNotEncodedReturnsNilIfNoPreviousValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
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

//<string charset="unicode" presence="optional">
//	<copy />
//</string>
func TestCanDeseraliseOptionalUnicodeStringCopyOperatorNotEncodedReturnsNilIfPreviousValueEmpty(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	dict.SetValue("UnicodeStringField", fix.NullValue{})
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Expected nil but deserialised actual value was: %v", result.Get())
	}
}

//<string charset="unicode">
//	<copy />
//</string>
func TestRequiresPmapReturnsTrueForRequiredUnicodeStringCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: true,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<string charset="unicode" presence="optional">
//	<copy />
//</string>
func TestRequiresPmapReturnsTrueForOptionalUnicodeStringCopyOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
		FieldDetails: Field{
			ID:       1,
			Name:     "UnicodeStringField",
			Required: false,
		},
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
