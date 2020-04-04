package presencemap

import (
	"bytes"
	"testing"
)

func TestCanReadSingleBytePMap(t *testing.T) {
	// Arrange 170 = (10101010)
	message := bytes.NewBuffer([]byte{170})
	pMap, _ := New(message)

	cases := []struct {
		expectedSet bool
	}{
		// 0101010 <- pMap
		{false},
		{true},
		{false},
		{true},
		{false},
		{true},
		{false},
	}

	for byteNumber, table := range cases {
		// Act
		isSet := pMap.GetIsSetAndIncrement()

		// Assert
		if isSet != table.expectedSet {
			t.Errorf("Expected pMap to not be set but was set byteNumber: %d", byteNumber)
		}
	}
}

func TestCanReadTwoBytePMap(t *testing.T) {
	// Arrange 42 = (00101010) 170 = (10101010)
	message := bytes.NewBuffer([]byte{42, 170})
	pMap, _ := New(message)

	cases := []struct {
		expectedSet bool
	}{
		// 0101010 0101010 <- pMap
		{false},
		{true},
		{false},
		{true},
		{false},
		{true},
		{false},
		{false},
		{true},
		{false},
		{true},
		{false},
		{true},
		{false},
	}

	for byteNumber, table := range cases {
		// Act
		isSet := pMap.GetIsSetAndIncrement()

		// Assert
		if isSet != table.expectedSet {
			t.Errorf("Expected pMap to not be set but was set byteNumber: %d", byteNumber)
		}
	}
}

func TestCanReadThreeBytePMap(t *testing.T) {
	// Arrange 42 = (00101010) 42 = (00101010) 235 = (11101011)
	message := bytes.NewBuffer([]byte{42, 42, 235})
	pMap, _ := New(message)

	cases := []struct {
		expectedSet bool
	}{
		// 0101010 0101010 1101011 <- pMap
		{false},
		{true},
		{false},
		{true},
		{false},
		{true},
		{false},
		{false},
		{true},
		{false},
		{true},
		{false},
		{true},
		{false},
		{true},
		{true},
		{false},
		{true},
		{false},
		{true},
		{true},
	}

	for byteNumber, table := range cases {
		// Act
		isSet := pMap.GetIsSetAndIncrement()

		// Assert
		if isSet != table.expectedSet {
			t.Errorf("Expected pMap to not be set but was set byteNumber: %d", byteNumber)
		}
	}
}
