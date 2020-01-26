package converter

import (
	"fmt"
	"strconv"
	"strings"
)

// ToString passes back the value given to it
func ToString(value string) (interface{}, error) {
	return value, nil
}

// ToMantissa converts the value specified to a normalized decimal (mantissa x10 ^ exponent) where the mantissa % 10 != 0.
// The function then returns the mantissa part of the decimal only.
func ToMantissa(value string) (interface{}, error) {
	_, mantissa, err := toDecimal(value)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse mantissa for value: %s", value)
	}
	return mantissa, nil
}

// ToExponent converts the value specified to a normalized decimal (mantissa x10 ^ exponent) where the mantissa % 10 != 0.
// The function then returns the exponent part of the decimal only.
func ToExponent(value string) (interface{}, error) {
	exponent, _, err := toDecimal(value)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse exponent for value: %s", value)
	}
	return exponent, nil
}

func toDecimal(value string) (int32, int64, error) {
	if value == "0" || value == "0.0" || value == "" {
		return int32(0), int64(0), nil
	}

	valueAsArray := strings.Split(value, "")
	exponentValue := 0
	mantissaBuilder := strings.Builder{}
	decimalLocation := 0

	for i, value := range valueAsArray {
		if value != "." {
			mantissaBuilder.WriteString(value)
		} else {
			decimalLocation = i
		}
	}

	mantissaValue, err := strconv.ParseInt(mantissaBuilder.String(), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	if decimalLocation == 0 {
		newMantissa := mantissaValue
		for newMantissa%10 == 0 {
			exponentValue = exponentValue + 1
			newMantissa = newMantissa / 10
		}
		mantissaValue = newMantissa
	} else {
		decimalLocation = decimalLocation + 1
		arrayLength := len(valueAsArray)
		exponentValue = -(arrayLength % decimalLocation)
	}

	return int32(exponentValue), int64(mantissaValue), nil
}

// ToInt32 converts the string to an int32 type, returning an error if the conversion fails
func ToInt32(value string) (interface{}, error) {
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}
	return int32(val), nil
}

// ToUInt32 converts the string to an uint32 type, returning an error if the conversion fails
func ToUInt32(value string) (interface{}, error) {
	val, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return nil, err
	}
	return uint32(val), nil
}

// ToInt64 converts the string to an int64 type, returning an error if the conversion fails
func ToInt64(value string) (interface{}, error) {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return int64(val), nil
}

// ToUInt64 converts the string to an uint64 type, returning an error if the conversion fails
func ToUInt64(value string) (interface{}, error) {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return uint64(val), nil
}