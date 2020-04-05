package decoder

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/value"
	"math/big"
	"reflect"
	"strings"
	"testing"
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

func TestCanReadMaximumUInt32(t *testing.T) {
	// Arrange 4294967295 = (00001111 01111111 01111111 01111111 11111111)
	expectedUintAsBytes := bytes.NewBuffer([]byte{15, 127, 127, 127, 255})
	var expectedUint uint32 = 4294967295

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

func TestUInt32IfStopBitNotFoundWithinBoundsReturnsError(t *testing.T) {
	// Arrange
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 48, 48, 138})

	// Act
	_, err := ReadUInt32(expectedUintAsBytes)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R6) {
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

func TestCanReadMaximumOptionalUInt32(t *testing.T) {
	// Arrange 4294967296 = (00010000 00000000 00000000 00000000 10000000)
	expectedUintAsBytes := bytes.NewBuffer([]byte{16, 0, 0, 0, 128})
	expectedResult := value.UInt32Value{Value: 4294967295}

	// Act
	result, err := ReadOptionalUInt32(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint32 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected uint32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadPositiveSingleByteInt32(t *testing.T) {
	// Arrange 138 = (10001010)
	expectedIntAsBytes := bytes.NewBuffer([]byte{138})
	var expectedInt int32 = 10

	// Act
	result, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int32 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int32, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestCanReadNegativeSingleByteInt32(t *testing.T) {
	// Arrange 246 = (11110110)
	expectedIntAsBytes := bytes.NewBuffer([]byte{246})
	var expectedInt int32 = -10

	// Act
	result, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int32 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int32, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestCanReadPositiveMultipleByteInt32(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedIntAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	var expectedInt int32 = 101455882

	// Act
	result, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int32 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int32, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestCanReadNegativeMultipleByteInt32(t *testing.T) {
	// Arrange -24751 = (01111110 00111110 11010001)
	expectedIntAsBytes := bytes.NewBuffer([]byte{126, 62, 209})
	var expectedInt int32 = -24751

	// Act
	result, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int32 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int32, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestCanReadMaximumInt32(t *testing.T) {
	// Arrange 2147483647 = (00000111 01111111 01111111 01111111 11111111)
	expectedIntAsBytes := bytes.NewBuffer([]byte{7, 127, 127, 127, 255})
	var expectedInt int32 = 2147483647

	// Act
	result, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int32 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int32, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestCanReadMinimumInt32(t *testing.T) {
	// Arrange -2147483648 = (01111000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{120, 0, 0, 0, 128})
	var expectedInt int32 = -2147483648

	// Act
	result, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int32 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int32, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestDoesNotOverflowInt32IfStopBitNotFoundWithinBounds(t *testing.T) {
	// Arrange
	expectedIntAsBytes := bytes.NewBuffer([]byte{126, 62, 62, 62, 62, 209})

	// Act
	_, err := ReadInt32(expectedIntAsBytes)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R6) {
		t.Errorf("Expected error about int32 overflow but got: %#v", err)
	}
}

func TestReadOptionalInt32ReturnsNilIfZeroEncoded(t *testing.T) {
	// Arrange nil = 10000000
	expectedIntAsBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int32 when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
	}
}

func TestReadOptionalPositiveInt32ReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange 130 = 10000010
	expectedIntAsBytes := bytes.NewBuffer([]byte{130})
	expectedResult := value.Int32Value{Value: 1}

	// Act
	result, err := ReadOptionalInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int32 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected int32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestReadOptionalNegativeInt32ReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange 251 = 11111011
	expectedIntAsBytes := bytes.NewBuffer([]byte{251})
	expectedResult := value.Int32Value{Value: -5}

	// Act
	result, err := ReadOptionalInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int32 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected int32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadMaximumOptionalInt32(t *testing.T) {
	// Arrange 2147483647 = (00001000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{8, 0, 0, 0, 128})
	expectedResult := value.Int32Value{Value: 2147483647}

	// Act
	result, err := ReadOptionalInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int32 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected optional int32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadMinimumOptionalInt32(t *testing.T) {
	// Arrange -2147483648 = (01111000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{120, 0, 0, 0, 128})
	expectedResult := value.Int32Value{Value: -2147483648}

	// Act
	result, err := ReadOptionalInt32(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int32 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected optional int32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadSingleByteUInt64(t *testing.T) {
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

func TestCanReadMultipleByteUInt64(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	var expectedUint uint64 = 101455882

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

func TestCanReadMaximumUInt64(t *testing.T) {
	// Arrange 18446744073709551615 = (00000001 01111111 01111111 01111111 01111111 01111111 01111111 01111111 01111111 11111111)
	expectedIntAsBytes := bytes.NewBuffer([]byte{1, 127, 127, 127, 127, 127, 127, 127, 127, 255})
	var expectedInt uint64 = 18446744073709551615

	// Act
	result, err := ReadUInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint64 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected uint64, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestDoesNotOverflowUInt64IfStopBitNotFoundWithinBounds(t *testing.T) {
	// Arrange
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 138})

	// Act
	_, err := ReadUInt64(expectedUintAsBytes)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R6) {
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

func TestCanReadMaximumOptionalUInt64(t *testing.T) {
	// Arrange 18446744073709551615 = (00000010 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{2, 0, 0, 0, 0, 0, 0, 0, 0, 128})
	expectedResult := value.UInt64Value{Value: 18446744073709551615 }

	// Act
	result, err := ReadOptionalUInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading uint64 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected uint64 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadPositiveSingleByteInt64(t *testing.T) {
	// Arrange 138 = (10001010)
	expectedIntAsBytes := bytes.NewBuffer([]byte{138})
	var expectedUint int64 = 10

	// Act
	result, err := ReadInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected int64, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestCanReadNegativeSingleByteInt64(t *testing.T) {
	// Arrange 202 = (11001010)
	expectedIntAsBytes := bytes.NewBuffer([]byte{202})
	var expectedUint int64 = -54

	// Act
	result, err := ReadInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected int64, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestCanReadPositiveMultipleByteInt64(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	var expectedUint int64 = 101455882

	// Act
	result, err := ReadInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected int64, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestCanReadNegativeMultipleByteInt64(t *testing.T) {
	// Arrange -101455882 = (01111111 01001111 01001111 01001111 11110110)
	expectedUintAsBytes := bytes.NewBuffer([]byte{127, 79, 79, 79, 246})
	var expectedUint int64 = -101455882

	// Act
	result, err := ReadInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	if result.Value != expectedUint {
		t.Errorf("Did not read the expected int64, expected: %#v, result: %#v", expectedUint, result)
	}
}

func TestCanReadMaximumInt64(t *testing.T) {
	// Arrange 9223372036854775807 = (00000000 01111111 01111111 01111111 01111111 01111111 01111111 01111111 01111111 11111111)
	expectedIntAsBytes := bytes.NewBuffer([]byte{0, 127, 127, 127, 127, 127, 127, 127, 127, 255})
	var expectedInt int64 = 9223372036854775807

	// Act
	result, err := ReadInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int64, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestCanReadMinimumInt64(t *testing.T) {
	// Arrange -9223372036854775808 = (01111111 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{127, 0, 0, 0, 0, 0, 0, 0, 0, 128})
	var expectedInt int64 = -9223372036854775808

	// Act
	result, err := ReadInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	if result.Value != expectedInt {
		t.Errorf("Did not read the expected int64, expected: %#v, result: %#v", expectedInt, result)
	}
}

func TestDoesNotOverflowInt64IfStopBitNotFoundWithinBounds(t *testing.T) {
	// Arrange
	expectedIntAsBytes := bytes.NewBuffer([]byte{79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 246})

	// Act
	_, err := ReadInt64(expectedIntAsBytes)

	// Assert
	if err == nil || !strings.Contains(err.Error(), errors.R6) {
		t.Errorf("Expected error about int64 overflow but got: %v", err)
	}
}

func TestReadOptionalInt64ReturnsNilIfZeroEncoded(t *testing.T) {
	// Arrange nil = 10000000
	expectedUintAsBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int64 when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
	}
}

func TestReadOptionalPositiveInt64ReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange -101455882 = (01111111 01001111 01001111 01001111 11110110)
	expectedUintAsBytes := bytes.NewBuffer([]byte{127, 79, 79, 79, 246})
	expectedResult := value.Int64Value{Value: -101455882}

	// Act
	result, err := ReadOptionalInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int64 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected int64 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestReadOptionalNegativeInt64ReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	expectedResult := value.Int64Value{Value: 101455881}

	// Act
	result, err := ReadOptionalInt64(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional int64 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected int64 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadMaximumOptionalInt64(t *testing.T) {
	// Arrange 9223372036854775807 = (00000001 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 128})
	expectedResult := value.Int64Value{Value: 9223372036854775807}

	// Act
	result, err := ReadOptionalInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected int64 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadMinimumOptionalInt64(t *testing.T) {
	// Arrange -9223372036854775808 = (01111111 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000000)
	expectedIntAsBytes := bytes.NewBuffer([]byte{127, 0, 0, 0, 0, 0, 0, 0, 0, 128})
	expectedResult := value.Int64Value{Value: -9223372036854775808}

	// Act
	result, err := ReadOptionalInt64(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading int64 when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected int64 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadBigIntWithAllowedOverflowMaxOverflowValue(t *testing.T) {
	// Arrange 18446744073709551615 = 00000001 01111111 01111111 01111111 01111111 01111111 01111111 01111111 01111111 11111111
	expectedIntAsBytes := bytes.NewBuffer([]byte{1, 127, 127, 127, 127, 127, 127, 127, 127, 255})
	bigResult := big.NewInt(0)
	expectedResult, _ := bigResult.SetString("18446744073709551615", 10)

	// Act
	result, err := ReadBigInt(expectedIntAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint32 when none was expected: %s", err)
	}

	areEqual := expectedResult.Cmp(result.Value)
	if areEqual != 0 {
		t.Errorf("Did not read the expected uint32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadBigIntWithAllowedOverflowMinOverflowValue(t *testing.T) {
	// Arrange -18446744073709551615 = 01111110 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000001
	expectedintAsBytes := bytes.NewBuffer([]byte{126, 0, 0, 0, 0, 0, 0, 0, 0, 129})
	bigResult := big.NewInt(0)
	expectedResult, _ := bigResult.SetString("-18446744073709551615", 10)

	// Act
	result, err := ReadBigInt(expectedintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint32 when none was expected: %s", err)
	}

	areEqual := expectedResult.Cmp(result.Value)
	if areEqual != 0 {
		t.Errorf("Did not read the expected uint32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestCanReadBigIntWithAllowedOverflowNoOverflowRequired(t *testing.T) {
	// Arrange 101455882 = (00110000 00110000 00110000 10001010)
	expectedintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 138})
	expectedResult := big.NewInt(101455882)

	// Act
	result, err := ReadBigInt(expectedintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional uint32 when none was expected: %s", err)
	}

	areEqual := expectedResult.Cmp(result.Value)
	if areEqual != 0 {
		t.Errorf("Did not read the expected uint32 value, expected: %#v, result: %#v", expectedResult, result)
	}
}

func TestReadOptionalBigIntReturnsNilIfZeroEncoded(t *testing.T) {
	// Arrange nil = 10000000
	expectedUintAsBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalBigInt(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional bigint when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
	}
}

func TestReadOptionalBigIntReturnsValueMinusOneForNonNilValues(t *testing.T) {
	// Arrange 0 = 10000001
	expectedUintAsBytes := bytes.NewBuffer([]byte{129})
	expectedResult := big.NewInt(0)

	// Act
	result, err := ReadOptionalBigInt(expectedUintAsBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading optional bigint when none was expected: %s", err)
	}

	areEqual := result.(value.BigInt).Value.Cmp(expectedResult)
	if areEqual != 0 {
		t.Errorf("Did not read the expected bigint value, expected: %#v, result: %#v", expectedResult, result)
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
		t.Errorf("Did not read the expected string, expected: %#v, result: %#v", expectedString, result)
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
		t.Errorf("Did not read the expected string, expected: %#v, result: %#v", expectedString, result)
	}
}

func TestReadByteVectorReturnsCorrectNumberOfBytes(t *testing.T) {
	// Arrange 10000010 00000001 00000010
	expectedBytes := bytes.NewBuffer([]byte{130, 1, 2})
	expectedResult := value.ByteVector{Value: []byte{1, 2}}

	// Act
	result, err := ReadByteVector(expectedBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading value when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected byte vector, expected: %v, result: %v", expectedResult, result)
	}
}

func TestReadOptionalByteVectorReturnsCorrectNumberOfBytes(t *testing.T) {
	// Arrange 10000010 00000001
	expectedBytes := bytes.NewBuffer([]byte{130, 1})
	expectedResult := value.ByteVector{Value: []byte{1}}

	// Act
	result, err := ReadOptionalByteVector(expectedBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading value when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedResult, result)
	if !areEqual {
		t.Errorf("Did not read the expected byte vector, expected: %v, result: %v", expectedResult, result)
	}
}

func TestReadOptionalByteVectorReturnsNilIfEncoded(t *testing.T) {
	// Arrange 10000000
	expectedBytes := bytes.NewBuffer([]byte{128})
	expectedNil := value.NullValue{}

	// Act
	result, err := ReadOptionalByteVector(expectedBytes)

	// Assert
	if err != nil {
		t.Errorf("Got an error reading value when none was expected: %s", err)
	}

	if result != expectedNil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result)
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
