package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<int32 />
func TestCanDeseraliseRequiredPositiveInt32(t *testing.T) {
	// Arrange 2 = 10000010
	messageAsBytes := bytes.NewBuffer([]byte{130})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(2)
	unitUnderTest := Int32{
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

//<int32 />
func TestCanDeseraliseRequiredNegativeInt32(t *testing.T) {
	// Arrange -2 = 11111110
	messageAsBytes := bytes.NewBuffer([]byte{254})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(-2)
	unitUnderTest := Int32{
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

//<int32 presence="optional"/>
func TestCanDeseraliseOptionalPositiveInt32Present(t *testing.T) {
	// Arrange 3 = 10000100
	messageAsBytes := bytes.NewBuffer([]byte{132})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(3)
	unitUnderTest := Int32{
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

//<int32 presence="optional"/>
func TestCanDeseraliseOptionalNegativeInt32Present(t *testing.T) {
	// Arrange 3 = 11111101
	messageAsBytes := bytes.NewBuffer([]byte{253})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(-3)
	unitUnderTest := Int32{
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

//<int32 presence="optional"/>
func TestCanDeseraliseOptionalInt32Null(t *testing.T) {
	// Arrange 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := Int32{
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

//<int32>
//	<constant value="132" />
//</int32>
func TestCanDeseraliseRequiredInt32ConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(132)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
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

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestCanDeseraliseOptionalInt32ConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
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

//<int32 presence="optional">
//	<constant value="132" />
//</int32>
func TestCanDeseraliseOptionalInt32ConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	expectedMessage := int32(132)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: int32(132),
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

//<int32 />
func TestRequiresPmapReturnsFalseForRequiredInt32NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
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

//<int32 presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalInt32NoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
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

//<int32>
//	<constant value="132" />
//</int32>
func TestRequiresPmapReturnsFalseForRequiredInt32ConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
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
	expectedMessage := int32(2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32>
//	<default value="5" />
//</int32>
func TestCanDeseraliseInt32DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(5)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32 presence="optional">
//	<default value="5" />
//</int32>
func TestCanDeseraliseOptionalInt32DefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 2 = 10000011
	messageAsBytes := bytes.NewBuffer([]byte{131})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{197}))
	expectedMessage := int32(2)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32 presence="optional">
//	<default value="5" />
//</int32>
func TestCanDeseraliseOptionalInt32DefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := int32(5)
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: int32(5),
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

//<int32 presence="optional">
//	<default />
//</int32>
func TestCanDeseraliseOptionalInt32DefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: nil,
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

//<int32>
//	<default value="132" />
//</int32>
func TestRequiresPmapReturnsTrueForRequiredInt32DefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Int32{
		FieldDetails: Field{
			ID:       1,
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
