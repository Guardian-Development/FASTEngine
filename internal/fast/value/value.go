package value

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Guardian-Development/fastengine/client/fix"
)

type Value interface {
	GetAsFix() fix.Value
	Add(value fix.Value) (fix.Value, error)
}

type NullValue struct {
}

func (value NullValue) GetAsFix() fix.Value {
	return fix.NullValue{}
}

func (value NullValue) Add(toAdd fix.Value) (fix.Value, error) {
	return fix.NullValue{}, nil
}

type StringValue struct {
	Value string
}

func (value StringValue) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

func (value StringValue) Add(toAdd fix.Value) (fix.Value, error) {
	// TODO
	return fix.NullValue{}, nil
}

type ByteVector struct {
	Value []byte
}

func (value ByteVector) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

func (value ByteVector) Add(toAdd fix.Value) (fix.Value, error) {
	// TODO
	return fix.NullValue{}, nil
}

type UInt32Value struct {
	Value uint32
}

func (value UInt32Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

func (value UInt32Value) Add(toAdd fix.Value) (fix.Value, error) {
	// TODO
	return fix.NullValue{}, nil
}

type Int32Value struct {
	Value int32
}

func (value Int32Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

func (value Int32Value) Add(toAdd fix.Value) (fix.Value, error) {
	rawValue := toAdd.Get().(int32)
	return addValueWithinInt32Constraints(int64(value.Value), int64(rawValue))
}

type UInt64Value struct {
	Value uint64
}

func (value UInt64Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

func (value UInt64Value) Add(toAdd fix.Value) (fix.Value, error) {
	// TODO
	return fix.NullValue{}, nil
}

type Int64Value struct {
	Value int64
}

func (value Int64Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

func (value Int64Value) Add(toAdd fix.Value) (fix.Value, error) {
	switch t := toAdd.Get().(type) {
	case int32:
		return addValueWithinInt32Constraints(value.Value, int64(t))
	case int64:
		return addValueWithinInt64Constraints(value.Value, t)
	}

	return fix.NullValue{}, fmt.Errorf("unsupported type to add int64 to: %#v", toAdd.Get())
}

type BigInt struct {
	Value *big.Int
}

func (value BigInt) GetAsFix() fix.Value {
	return fix.NewRawValue(*value.Value)
}

func (value BigInt) Add(toAdd fix.Value) (fix.Value, error) {
	rawValue := toAdd.Get().(int64)
	copy := big.NewInt(0).Set(value.Value)
	valueAfterAddition := copy.Add(copy, big.NewInt(rawValue))

	// if the addition does not stay within the bounds of an int64, we have an overflow and report an error
	if !valueAfterAddition.IsInt64() {
		return nil, fmt.Errorf("%v + %v would overflow int64", rawValue, value.Value.Int64())
	}

	return fix.NewRawValue(valueAfterAddition.Int64()), nil
}

func addValueWithinInt32Constraints(readValue int64, value int64) (fix.Value, error) {
	// positive value and value you add is greater than the difference between the positive value and the max value, you will positive overflow
	if readValue > 0 && value > math.MaxInt32-readValue {
		return nil, fmt.Errorf("%v + %v would overflow int32", readValue, value)
	}
	// negative value and you're add is greater than the difference between the negative value and the min value, you will negative overflow
	if value < math.MinInt32-readValue {
		return nil, fmt.Errorf("%v + %v would overflow int32", readValue, value)
	}

	return fix.NewRawValue(int32(readValue + value)), nil
}

func addValueWithinInt64Constraints(readValue int64, value int64) (fix.Value, error) {
	// positive value and value you add is greater than the difference between the positive value and the max value, you will positive overflow
	if readValue > 0 && value > math.MaxInt64-readValue {
		return nil, fmt.Errorf("%v + %v would overflow int64", readValue, value)
	}
	// negative value and you're add is greater than the difference between the negative value and the min value, you will negative overflow
	if value < math.MinInt64-readValue {
		return nil, fmt.Errorf("%v + %v would overflow int64", readValue, value)
	}

	return fix.NewRawValue(int64(readValue + value)), nil
}
