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

type AsciiStringDecoder struct {
}

func (AsciiStringDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadString(inputSource)
}

func (AsciiStringDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalString(inputSource)
}

type AsciiStringDeltaDecoder struct {
}

func (AsciiStringDeltaDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	subtractionLength, err := ReadInt32(inputSource)
	if err != nil {
		return nil, err
	}

	asciiValue, err := ReadString(inputSource)
	if err != nil {
		return nil, err
	}
	asciiValue.ItemsToRemove = subtractionLength.Value
	return asciiValue, nil
}

func (AsciiStringDeltaDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	subtractionLength, err := ReadOptionalInt32(inputSource)
	if err != nil {
		return nil, err
	}

	// if no subtraction length present, then no delta encoded
	switch t := subtractionLength.(type) {
	case value.NullValue:
		return t, nil
	}

	asciiValue, err := ReadString(inputSource)
	if err != nil {
		return nil, err
	}
	asciiValue.ItemsToRemove = subtractionLength.(value.Int32Value).Value
	return asciiValue, nil
}
