package field

import (
	"bytes"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<string charset="unicode"/>
func TestCanDeseraliseRequiredUnicodeString(t *testing.T) {
	// Arrange TEST1 = 10000101 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{133, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
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

//<string charset="unicode" presence="optional"/>
func TestCanDeseraliseOptionalUnicodeStringPresent(t *testing.T) {
	// Arrange TEST1 = 10000110 01010100 01000101 01010011 01010100 00110001
	messageAsBytes := bytes.NewBuffer([]byte{134, 84, 69, 83, 84, 49})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST1"
	unitUnderTest := UnicodeString{
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
func TestCanDeseraliseOptionalUnicodeStringNull(t *testing.T) {
	// Arrange TEST1 = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := UnicodeString{
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
func TestCanDeseraliseRequiredUnicodeStringConstantOperatorNotEncoded(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
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
func TestCanDeseraliseOptionalUnicodeStringConstantOperatorNotEncodedReturnsNilValue(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	unitUnderTest := UnicodeString{
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
func TestCanDeseraliseOptionalUnicodeStringConstantOperatorEncodedReturnsConstantValue(t *testing.T) {
	// Arrange pmap = 11000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	expectedMessage := "TEST2"
	unitUnderTest := UnicodeString{
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

//<string>
//	<constant value="TEST2" />
//</string>
func TestRequiresPmapReturnsFalseForRequiredUnicodeStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
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
func TestRequiresPmapReturnsTrueForOptionalUnicodeStringConstantOperator(t *testing.T) {
	// Arrange
	unitUnderTest := UnicodeString{
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
