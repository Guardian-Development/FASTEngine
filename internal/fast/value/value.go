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
	// TODO
	return fix.NullValue{}, nil
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
	rawValue := toAdd.Get().(int64)

	// positive value and value you add is greater than the difference between the positive value and the max value, you will positive overflow
	if value.Value > 0 && rawValue > math.MaxInt64-value.Value {
		return nil, fmt.Errorf("%v + %v would overflow int64", value.Value, rawValue)
	}
	// negative value and you're add is greater than the difference between the negative value and the min value, you will negative overflow
	if rawValue < math.MinInt64-value.Value {
		return nil, fmt.Errorf("%v + %v would overflow int64", value.Value, rawValue)
	}

	return fix.NewRawValue(value.Value + rawValue), nil
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

type NullValue struct {
}

func (value NullValue) GetAsFix() fix.Value {
	return fix.NullValue{}
}

func (value NullValue) Add(toAdd fix.Value) (fix.Value, error) {
	return fix.NullValue{}, nil
}
