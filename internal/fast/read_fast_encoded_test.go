package fast

import (
	"bytes"
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

	if result != expectedUint {
		t.Errorf("Did not read the expected uint32, expected: %d, result: %d", expectedUint, result)
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

	if result != expectedUint {
		t.Errorf("Did not read the expected uint32, expected: %d, result: %d", expectedUint, result)
	}
}

func TestDoesNotOverflowUint32IfStopBitNotFoundWithinBounds(t *testing.T) {
	// Arrange
	expectedUintAsBytes := bytes.NewBuffer([]byte{48, 48, 48, 48, 138})

	// Act
	_, err := ReadUInt32(expectedUintAsBytes)

	// Assert
	if err == nil || err.Error() != "More than 4 bytes have been read without reading a stop bit, this will overflow a uint32" {
		t.Errorf("Expected error about uint32 overflow but got: %v", err)
	}
}
