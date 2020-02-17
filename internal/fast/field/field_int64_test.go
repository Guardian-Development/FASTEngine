package field

import (
	"bytes"
	"testing"

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
