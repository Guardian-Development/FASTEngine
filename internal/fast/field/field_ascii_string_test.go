package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<string />
func TestCanDeseraliseRequiredAsciiString(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
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

//<string presence="optional"/>
func TestCanDeseraliseOptionalAsciiStringPresent(t *testing.T) {
	// Arrange TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
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

//<string presence="optional"/>
func TestCanDeseraliseOptionalAsciiStringNull(t *testing.T) {
	// Arrange TEST1 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := AsciiString{
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

//<string>
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseRequiredAsciiStringConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
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

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
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

//<string presence="optional">
//	<constant value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Constant{
			ConstantValue: "TEST2",
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

//<string />
func TestRequiresPmapReturnsFalseForRequiredAsciiStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
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

//<string presence="optional"/>
func TestRequiresPmapReturnsFalseForOptionalAsciiStringNoOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
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

//<string>
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsFalseForRequiredAsciiStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
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
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string>
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseAsciiStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorEncodedReturnsValueFromStream(t *testing.T) {
	// Arrange pmap = 11000000 TEST1 = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	expectedMessage := "TEST1"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string presence="optional">
//	<default value="TEST2" />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorNotEncodedReturnsDefaultValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST2"
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		Operation: operation.Default{
			DefaultValue: "TEST2",
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

//<string presence="optional">
//	<default />
//</string>
func TestCanDeseraliseOptionalAsciiStringDefaultOperatorNotEncodedReturnsDefaultNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := AsciiString{
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

//<string>
//	<default value="TEST2" />
//</string>
func TestRequiresPmapReturnsTrueForRequiredAsciiStringDefaultOperator(t *testing.T) {
	// Arrange
	unitUnderTest := AsciiString{
		FieldDetails: Field{
			ID:       1,
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
