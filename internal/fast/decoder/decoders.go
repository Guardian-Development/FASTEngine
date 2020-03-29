package decoder

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// Decoder is used to couple the reading of required and optional values of the same type
type Decoder interface {
	ReadValue(inputSource *bytes.Buffer) (value.Value, error)
	ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error)
}

//TODO: tests in this package should use these structs not the underlying methods themselves

type Int32Decoder struct {
}

func (Int32Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadInt32(inputSource)
}

func (Int32Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalInt32(inputSource)
}

type UInt32Decoder struct {
}

func (UInt32Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadUInt32(inputSource)
}

func (UInt32Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalUInt32(inputSource)
}

type Int64Decoder struct {
}

func (Int64Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadInt64(inputSource)
}

func (Int64Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalInt64(inputSource)
}

type UInt64Decoder struct {
}

func (UInt64Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadUInt64(inputSource)
}

func (UInt64Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalUInt64(inputSource)
}

type BitIntDecoder struct {
}

func (BitIntDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadBigInt(inputSource)
}

func (BitIntDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalBigInt(inputSource)
}
