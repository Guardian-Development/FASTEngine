package value

import (
	"math/big"

	"github.com/Guardian-Development/fastengine/client/fix"
)

type Value interface {
	GetAsFix() fix.Value
}

type StringValue struct {
	Value string
}

func (value StringValue) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

type UInt32Value struct {
	Value uint32
}

func (value UInt32Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

type Int32Value struct {
	Value int32
}

func (value Int32Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

type UInt64Value struct {
	Value uint64
}

func (value UInt64Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

type Int64Value struct {
	Value int64
}

func (value Int64Value) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

type BigInt struct {
	Value *big.Int
}

func (value BigInt) GetAsFix() fix.Value {
	return fix.NewRawValue(*value.Value)
}

type ByteVector struct {
	Value []byte
}

func (value ByteVector) GetAsFix() fix.Value {
	return fix.NewRawValue(value.Value)
}

type NullValue struct {
}

func (value NullValue) GetAsFix() fix.Value {
	return fix.NullValue{}
}
