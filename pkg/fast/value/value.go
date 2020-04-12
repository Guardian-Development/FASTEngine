package value

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Guardian-Development/fastengine/pkg/fast/errors"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// Value is a fast decoded value that can be converted to a fix representation
type Value interface {
	GetAsFix() fix.Value
	Add(value fix.Value) (fix.Value, error)
}

// NullValue represents a null fast value
type NullValue struct {
}

// GetAsFix returns a fix null representation
func (value NullValue) GetAsFix() fix.Value {
	return fix.NullValue{}
}

// Add returns a fix null representation
func (value NullValue) Add(toAdd fix.Value) (fix.Value, error) {
	return fix.NullValue{}, nil
}

// StringValue represents a string fast value. Items to remove is used when applying a delta to a string
type StringValue struct {
	Value         string
	ItemsToRemove int32
}

// GetAsFix returns a raw string wrapped in a fix type
func (value StringValue) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

// Add applies the read value to the previous toAdd value. ItemsToRemove represents how much of the existing value to overwrite.
// Positive means overwrite from end of existing string, negative means overwrite from beginning of existing string.
func (value StringValue) Add(toAdd fix.Value) (fix.Value, error) {
	existingValue := toAdd.Get().(string)
	// prepend
	if value.ItemsToRemove < 0 {
		if value.ItemsToRemove == -1 {
			return fix.NewRawValue(value.Value + existingValue), nil
		}

		itemsToRemove := (-value.ItemsToRemove) - 1
		if itemsToRemove > int32(len(existingValue)) {
			return nil, fmt.Errorf("%s: removing %d values from string %s", errors.D7, itemsToRemove, existingValue)
		}
		stringWithRemovedChars := existingValue[itemsToRemove:]
		return fix.NewRawValue(value.Value + stringWithRemovedChars), nil
	}

	// append
	itemsToRemove := int32(len(existingValue)) - value.ItemsToRemove
	if itemsToRemove < 0 {
		return nil, fmt.Errorf("%s: removing %d values from string %s", errors.D7, value.ItemsToRemove, existingValue)
	}
	stringWithRemovedChars := existingValue[:int32(len(existingValue))-value.ItemsToRemove]
	return fix.NewRawValue(stringWithRemovedChars + value.Value), nil
}

// ApplyTail overwrites the end of the previous string with the read value
func (value StringValue) ApplyTail(baseValue fix.Value) (fix.Value, error) {
	baseValueAsChars := []rune(baseValue.Get().(string))
	readValueAsChars := []rune(value.Value)
	indexToAppendReadValue := len(baseValueAsChars) - len(readValueAsChars)

	// read more than base value, read value replaces all of base value
	if indexToAppendReadValue <= 0 {
		return value.GetAsFix(), nil
	}

	start := baseValueAsChars[0:indexToAppendReadValue]
	combinedValue := append(start, readValueAsChars...)
	return fix.NewRawValue(string(combinedValue)), nil
}

// ByteVector represents a byte vector fast value. Items to remove is used when applying a delta to a byte vector
type ByteVector struct {
	Value         []byte
	ItemsToRemove int32
}

// GetAsFix returns a raw []byte wrapped in a fix type
func (value ByteVector) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

// Add applies the read value to the previous toAdd value. ItemsToRemove represents how much of the existing value to overwrite.
// Positive means overwrite from end of existing []byte, negative means overwrite from beginning of existing []byte.
func (value ByteVector) Add(toAdd fix.Value) (fix.Value, error) {
	existingValue := toAdd.Get().([]byte)
	// prepend
	if value.ItemsToRemove < 0 {
		if value.ItemsToRemove == -1 {
			return fix.NewRawValue(append(value.Value, existingValue...)), nil
		}

		itemsToRemove := (-value.ItemsToRemove) - 1
		if itemsToRemove > int32(len(existingValue)) {
			return nil, fmt.Errorf("%s: removing %d values from bytevector %#v", errors.D7, itemsToRemove, existingValue)
		}
		vectorWithRemovedBytes := existingValue[itemsToRemove:]
		return fix.NewRawValue(append(value.Value, vectorWithRemovedBytes...)), nil
	}

	// append
	itemsToRemove := int32(len(existingValue)) - value.ItemsToRemove
	if itemsToRemove < 0 {
		return nil, fmt.Errorf("%s: removing %d values from bytevector %#v", errors.D7, value.ItemsToRemove, existingValue)
	}
	vectorWithRemovedBytes := existingValue[:int32(len(existingValue))-value.ItemsToRemove]
	return fix.NewRawValue(append(vectorWithRemovedBytes, value.Value...)), nil
}

// ApplyTail overwrites the end of the previous []byte with the read value
func (value ByteVector) ApplyTail(baseValue fix.Value) (fix.Value, error) {
	baseValueAsBytes := baseValue.Get().([]byte)
	indexToAppendReadValue := len(baseValueAsBytes) - len(value.Value)

	// read more than base value, read value replaces all of base value
	if indexToAppendReadValue <= 0 {
		return value.GetAsFix(), nil
	}

	start := baseValueAsBytes[0:indexToAppendReadValue]
	combinedValue := append(start, value.Value...)
	return fix.NewRawValue(combinedValue), nil
}

// UInt32Value represents a uint32 fast value
type UInt32Value struct {
	Value uint32
}

// GetAsFix returns a raw uint32 wrapped in a fix type
func (value UInt32Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

// Add the previous value to the read value, assuring we stay within the constraints of a uint32
func (value UInt32Value) Add(toAdd fix.Value) (fix.Value, error) {
	rawValue := toAdd.Get().(uint32)
	return addValueWithinUInt32Constraints(int64(value.Value), int64(rawValue))
}

// Int32Value represents a int32 fast value
type Int32Value struct {
	Value int32
}

// GetAsFix returns a raw int32 wrapped in a fix type
func (value Int32Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

// Add the previous value to the read value, assuring we stay within the constraints of a int32
func (value Int32Value) Add(toAdd fix.Value) (fix.Value, error) {
	rawValue := toAdd.Get().(int32)
	return addValueWithinInt32Constraints(int64(value.Value), int64(rawValue))
}

// UInt64Value represents a uint64 fast value
type UInt64Value struct {
	Value uint64
}

// GetAsFix returns a raw uint64 wrapped in a fix type
func (value UInt64Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

// Add should never be called on this type, as we should be treating uint64 values of big ints when decoding them from the stream to allow for byte overflow
func (value UInt64Value) Add(toAdd fix.Value) (fix.Value, error) {
	// this is never used, as we use BigInt to allow addition where we temporarily overflow int64
	return nil, fmt.Errorf("unsupported operation, big int should be used to add uint64 operations")
}

// Int64Value represents a int64 fast value
type Int64Value struct {
	Value int64
}

// GetAsFix returns a raw int64 wrapped in a fix type
func (value Int64Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

// Add is used when we are deserialising an int32 (this is its overflow type, much like int64 overflow type is big int).
func (value Int64Value) Add(toAdd fix.Value) (fix.Value, error) {
	// this is used to add to 32 bit integer types, allowing for temporarily overflow of 32 bits
	switch t := toAdd.Get().(type) {
	case int32:
		return addValueWithinInt32Constraints(value.Value, int64(t))
	case uint32:
		return addValueWithinUInt32Constraints(value.Value, int64(t))
	}

	return fix.NullValue{}, fmt.Errorf("unsupported type to add int64 to: %#v", toAdd.Get())
}

// BigInt represents a uint64 and int64 when we are allowing for byte overflows
type BigInt struct {
	Value *big.Int
}

// GetAsFix returns a raw big.Int wrapped in a fix type
func (value BigInt) GetAsFix() fix.Value {
	return fix.NewRawValue(*value.Value)
}

// Add the previous value to the read value, assuring we stay within the constraints of either an int64 or uint64 depending on the previous value
func (value BigInt) Add(toAdd fix.Value) (fix.Value, error) {
	valueCopy := big.NewInt(0).Set(value.Value)

	switch t := toAdd.Get().(type) {
	case int64:
		valueAfterAddition := valueCopy.Add(valueCopy, big.NewInt(t))

		// if the addition does not stay within the bounds of an int64, we have an overflow and report an error
		if !valueAfterAddition.IsInt64() {
			return nil, fmt.Errorf("%s, %v + %v would overflow int64", errors.R4, t, value.Value.Int64())
		}
		return fix.NewRawValue(valueAfterAddition.Int64()), nil
	case uint64:
		valueAfterAddition := valueCopy.Add(valueCopy, big.NewInt(0).SetUint64(t))

		// if the addition does not stay within the bounds of an uint64, we have an overflow and report an error
		if !valueAfterAddition.IsUint64() {
			return nil, fmt.Errorf("%s, %v + %v would overflow uint64", errors.R4, t, value.Value.Uint64())
		}
		return fix.NewRawValue(valueAfterAddition.Uint64()), nil
	}

	return fix.NullValue{}, fmt.Errorf("unsupported type to add big int to: %#v", toAdd.Get())
}

func addValueWithinUInt32Constraints(readValue int64, value int64) (fix.Value, error) {
	// positive value and value you add is greater than the difference between the positive value and the max value, you will positive overflow
	if readValue > 0 && uint64(value) > uint64(math.MaxUint32)-uint64(readValue) {
		return nil, fmt.Errorf("%s, %v + %v would overflow uint32", errors.R4, readValue, value)
	}
	// if subtracting the value would take us below 0, you will negative overflow
	if 0 > value+readValue {
		return nil, fmt.Errorf("%s, %v + %v would overflow uint32", errors.R4, readValue, value)
	}

	return fix.NewRawValue(uint32(readValue + value)), nil
}

func addValueWithinInt32Constraints(readValue int64, value int64) (fix.Value, error) {
	// positive value and value you add is greater than the difference between the positive value and the max value, you will positive overflow
	if readValue > 0 && value > math.MaxInt32-readValue {
		return nil, fmt.Errorf("%s, %v + %v would overflow int32", errors.R4, readValue, value)
	}
	// negative value and you're add is greater than the difference between the negative value and the min value, you will negative overflow
	if value < math.MinInt32-readValue {
		return nil, fmt.Errorf("%s, %v + %v would overflow int32", errors.R4, readValue, value)
	}

	return fix.NewRawValue(int32(readValue + value)), nil
}
