package converter

import "testing"

func TestCanConvertStringToString(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue string
	}{
		{"hello", "hello"},
		{"-5", "-5"},
		{"TeS@1!!!!", "TeS@1!!!!"},
	}

	for _, testCase := range testCases {
		result, err := ToString(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}
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
		{"1", 1},
		{"-5", -5},
		{"-100", -100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		result, err := ToInt32(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}
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
		{"1", 1},
		{"5", 5},
		{"100", 100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		result, err := ToUInt32(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}
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
		{"1", 1},
		{"-5", -5},
		{"-100", -100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		result, err := ToInt64(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}
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
		{"1", 1},
		{"5", 5},
		{"100", 100},
		{"17", 17},
		{"0", 0},
	}

	for _, testCase := range testCases {
		result, err := ToUInt64(testCase.input)
		if err != nil {
			t.Errorf("Received error when none was expected: %v", err)
			continue
		}
		if result != testCase.expectedValue {
			t.Errorf("Failed to get correct result when converting value ToUInt64, expected : %d, actual %d", testCase.expectedValue, result)
			continue
		}
	}
}
