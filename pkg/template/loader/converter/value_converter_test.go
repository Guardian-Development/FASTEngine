package converter

import (
	"reflect"
	"testing"
)

func TestCanConvertStringToString(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue string
	}{
		// Arrange
		{"hello", "hello"},
		{"-5", "-5"},
		{"TeS@1!!!!", "TeS@1!!!!"},
	}

	for _, testCase := range testCases {
		// Act
		result, err := ToString(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		if result != testCase.expectedValue {
			t.Errorf("Failed to get correct result when converting value ToString, expected : %s, actual %s", testCase.expectedValue, result)
			continue
		}
	}
}

func TestCanConvertStringToExponentAndMantissaParts(t *testing.T) {
	testCases := []struct {
		input                 string
		expectedMantissaValue int64
		expectedExponentValue int32
	}{
		// Arrange
		{"10", 1, 1},
		{"-1.5", -15, -1},
		{"7.6", 76, -1},
		{"0.2", 2, -1},
		{"100", 1, 2},
		{"152", 152, 0},
		{"1", 1, 0},
		{"1002", 1002, 0},
		{"0", 0, 0},
	}

	for _, testCase := range testCases {
		// Act
		mantissaResult, err := ToMantissa(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}
		exponentResult, err := ToExponent(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		if mantissaResult != testCase.expectedMantissaValue {
			t.Errorf(
				"Failed to get correct mantissa when converting value ToDecimal, expectedMantissa : %d, actualMantissa %d, expectedExponent: %d, actualExponent: %d",
				testCase.expectedMantissaValue, mantissaResult, testCase.expectedExponentValue, exponentResult)
			continue
		}
		if exponentResult != testCase.expectedExponentValue {
			t.Errorf(
				"Failed to get correct exponent when converting value ToDecimal, expectedMantissa : %d, actualMantissa %d, expectedExponent: %d, actualExponent: %d",
				testCase.expectedMantissaValue, mantissaResult, testCase.expectedExponentValue, exponentResult)
			continue
		}
	}
}

func TestCanConvertStringToInt32(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue int32
	}{
		// Arrange
		{"1", 1},
		{"-5", -5},
		{"-100", -100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		// Act
		result, err := ToInt32(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		if result != testCase.expectedValue {
			t.Errorf("Failed to get correct result when converting value ToInt32, expected : %d, actual %d", testCase.expectedValue, result)
			continue
		}
	}
}

func TestCanConvertStringToUInt32(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue uint32
	}{
		// Arrange
		{"1", 1},
		{"5", 5},
		{"100", 100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		// Act
		result, err := ToUInt32(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		if result != testCase.expectedValue {
			t.Errorf("Failed to get correct result when converting value ToUInt32, expected : %d, actual %d", testCase.expectedValue, result)
			continue
		}
	}
}

func TestCanConvertStringToInt64(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue int64
	}{
		// Arrange
		{"1", 1},
		{"-5", -5},
		{"-100", -100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		// Act
		result, err := ToInt64(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		if result != testCase.expectedValue {
			t.Errorf("Failed to get correct result when converting value ToInt64, expected : %d, actual %d", testCase.expectedValue, result)
			continue
		}
	}
}

func TestCanConvertStringToUInt64(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue uint64
	}{
		// Arrange
		{"1", 1},
		{"5", 5},
		{"100", 100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		// Act
		result, err := ToUInt64(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		if result != testCase.expectedValue {
			t.Errorf("Failed to get correct result when converting value ToUInt64, expected : %d, actual %d", testCase.expectedValue, result)
			continue
		}
	}
}

func TestCanConvertStringToByteArray(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue []byte
	}{
		// Arrange
		{"A1 B2 CF DE", []byte{0xA1, 0xB2, 0xCF, 0xDE}},
		{"01 11 DD ee", []byte{0x01, 0x11, 0xDD, 0xEE}},
		{"", []byte{}},
		{"FF 00", []byte{0xFF, 0x00}},
		{"00", []byte{0x00}},
	}

	for _, testCase := range testCases {
		// Act
		result, err := ToByteVector(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}

		// Assert
		areEqual := reflect.DeepEqual(result, testCase.expectedValue)
		if !areEqual {
			t.Errorf("Failed to get correct result when converting value ToByteVector, expected : %v, actual %v", testCase.expectedValue, result)
			continue
		}
	}
}

func TestConvertStringToByteArrayReturnsErrorIfOddNumberOfCharacters(t *testing.T) {
	// Arrange
	unevenInput := "54F"

	// Act
	_, err := ToByteVector(unevenInput)

	// Assert
	if err == nil || err.Error() != "you must specify a byte vector as an even amount of hex digits" {
		t.Errorf("Expected error message informing user of correct input type for byteArray value, but got: %v", err)
	}
}
