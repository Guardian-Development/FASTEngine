package fast

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

func TestCanReadSingleByteUint32(t *testing.T) {
	// Arrange 138 = (10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{138})
	var expectedUint uint32 = 10

	// Act
	result, err := ReadUInt32(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint32 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected uint32, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestCanReadMultipleByteUint32(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	var expectedUint uint32 = 101455882

	// Act
	result, err := ReadUInt32(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint32 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected uint32, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestDoesNotOverflowUint32IfStopBitNotFoundWithinBounds(t *testing.T) {
	// Arrange
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 48, 138})

	// Act
	_, err := ReadUInt32(expectedUintAsBytes)

	// Assert
	if err == nil || err.Error() != "More than 4 bytes have been read without reading a stop bit, this will overflow a uint32" {
		t.Errorf("Expected error about uint32 overflow but got: %#v", err)
	}
}

func TestReadOptionalUInt32ReturnsNilIfZeroEncoded(t *testing.T) {
	// Arrange nil = 10000000
	expectedUintAsBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalUInt32(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint32 when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
	}
}

func TestReadOptionalUInt32ReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange 0 = 10000001
	expectedUintAsBytes := bytes.NewBuffer([]byte{129})
	expectedResult := value.UInt32Value{Value: 0}

	// Act
	result, err := ReadOptionalUInt32(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint32 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected uint32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadSingleByteUint64(t *testing.T) {
	// Arrange 138 = (10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{138})
	var expectedUint uint64 = 10

	// Act
	result, err := ReadUInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint64 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected uint64, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestCanReadMultipleByteUint64(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	var expectedUint uint64 = 101455882

	// Act
	result, err := ReadUInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint32 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected uint32, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestDoesNotOverflowUint64IfStopBitNotFoundWithinBounds(t *testing.T) {
	// Arrange
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 48, 48, 48, 48, 48, 138})

	// Act
	_, err := ReadUInt64(expectedUintAsBytes)

	// Assert
	if err == nil || err.Error() != "More than 8 bytes have been read without reading a stop bit, this will overflow a uint64" {
		t.Errorf("Expected error about uint64 overflow but got: %v", err)
	}
}

func TestReadOptionalUInt64ReturnsNilIfZeroEncoded(t *testing.T) {
	// Arrange nil = 10000000
	expectedUintAsBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalUInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint64 when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
	}
}

func TestReadOptionalUInt64ReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange 0 = 10000001
	expectedUintAsBytes := bytes.NewBuffer([]byte{129})
	expectedResult := value.UInt64Value{Value: 0}

	// Act
	result, err := ReadOptionalUInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint64 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected uint64 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadSingleByteString(t *testing.T) {
	// Arrange 'A' = (11000001)
	expectedStringAsBytes := bytes.NewBuffer([]byte{193})
	var expectedString string = "A"

	// Act
	result, err := ReadString(expectedStringAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading string when none was expected: %s", err)
	}

	if result.Value != expectedString {
		t.Errorf("Did not read the expected string, expected: %s, result: %v", expectedString, result)
	}
}

func TestCanReadEmptyString(t *testing.T) {
	// Arrange '' = (10000000)
	expectedStringAsBytes := bytes.NewBuffer([]byte{128})
	expectedString := ""

	// Act
	result, err := ReadString(expectedStringAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading string when none was expected: %s", err)
	}

	if result.Value != expectedString {
		t.Errorf("Did not read the expected string, expected: %s, result: %v", expectedString, result)
	}
}

func TestCanReadMultipleByteString(t *testing.T) {
	// Arrange 'AbC12~' = (01000001, 01100010, 01000011, 00110001, 00110010, 11111110)
	expectedStringAsBytes := bytes.NewBuffer([]byte{65, 98, 67, 49, 50, 254})
	var expectedString string = "AbC12~"

	// Act
	result, err := ReadString(expectedStringAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading string when none was expected: %s", err)
	}

	if result.Value != expectedString {
		t.Errorf("Did not read the expected string, expected: %s, result: %#v", expectedString, result)
	}
}

func TestReadOptionalStringReturnsNilIfEncoded(t *testing.T) {
	// Arrange 'AbC12~' = (10000000)
	expectedStringAsBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalString(expectedStringAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading string when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
	}
}

func TestReadOptionalStringReturnsEmptyStringIfEncoded(t *testing.T) {
	// Arrange '' = (00000000, 10000000)
	expectedStringAsBytes := bytes.NewBuffer([]byte{0, 128})
	expectedString := value.StringValue{Value: ""}

	// Act
	result, err := ReadOptionalString(expectedStringAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading string when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedString, result)
	if !areEqual {
		t.Errorf("Did not read the expected string, expected: %s, result: %#v", expectedString, result)
	}
}

func TestReadOptionalStringReturnsEncodedString(t *testing.T) {
	// Arrange 'AbC12~' = (01000001, 01100010, 01000011, 00110001, 00110010, 11111110)
	expectedStringAsBytes := bytes.NewBuffer([]byte{65, 98, 67, 49, 50, 254})
	expectedString := value.StringValue{Value: "AbC12~"}

	// Act
	result, err := ReadOptionalString(expectedStringAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading string when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedString, result)
	if !areEqual {
		t.Errorf("Did not read the expected string, expected: %s, result: %#v", expectedString, result)
	}
}

func TestReadValueReturnsRawBytesWithStopBitRemoved(t *testing.T) {
	// Arrrange 00010010 10001000 -> [00100100, 00010000]
	expectedBytes := bytes.NewBuffer([]byte{18, 136})
	expectedResult := []byte{36, 16}

	// Act
	result, err := ReadValue(expectedBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading value when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected byte array, expected: %v, result: %v", expectedResult, result)
	}
}
