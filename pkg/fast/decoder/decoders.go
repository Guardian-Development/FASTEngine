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

// Int32Decoder performs a read/optional read of a FAST encoded int32
type Int32Decoder struct {
}

// ReadValue fast encoded int32
func (Int32Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadInt32(inputSource)
}

// ReadOptionalValue fast encoded optional int32
func (Int32Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalInt32(inputSource)
}

// UInt32Decoder performs a read/optional read of a FAST encoded uint32
type UInt32Decoder struct {
}

// ReadValue fast encoded uint32
func (UInt32Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadUInt32(inputSource)
}

// ReadOptionalValue fast encoded optional uint32
func (UInt32Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalUInt32(inputSource)
}

// Int64Decoder performs a read/optional read of a FAST encoded int64
type Int64Decoder struct {
}

// ReadValue fast encoded int64
func (Int64Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadInt64(inputSource)
}

// ReadOptionalValue fast encoded optional int64
func (Int64Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalInt64(inputSource)
}

// UInt64Decoder performs a read/optional read of a FAST encoded uint64
type UInt64Decoder struct {
}

// ReadValue fast encoded uint64
func (UInt64Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadUInt64(inputSource)
}

// ReadOptionalValue fast encoded optional uint64
func (UInt64Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalUInt64(inputSource)
}

// BitIntDecoder performs a read/optional read of a FAST encoded int64 with allowed overflow of a single byte
type BitIntDecoder struct {
}

// ReadValue fast encoded int64 with allowed overflow
func (BitIntDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadBigInt(inputSource)
}

// ReadOptionalValue fast encoded optional int64 with allowed overflow
func (BitIntDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalBigInt(inputSource)
}

// AsciiStringDecoder performs a read/optional read of a FAST encoded string
type AsciiStringDecoder struct {
}

// ReadValue fast encoded string
func (AsciiStringDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadString(inputSource)
}

// ReadOptionalValue fast encoded optional string
func (AsciiStringDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalString(inputSource)
}

// AsciiStringDeltaDecoder performs a read/optional read of a FAST encoded string delta
type AsciiStringDeltaDecoder struct {
}

// ReadValue fast encoded string delta
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

// ReadOptionalValue fast encoded string delta
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

// ByteVectorDecoder performs a read/optional read of a FAST encoded byte vector
type ByteVectorDecoder struct {
}

// ReadValue fast encoded byte vector
func (ByteVectorDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadByteVector(inputSource)
}

// ReadOptionalValue fast encoded byte vector
func (ByteVectorDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalByteVector(inputSource)
}

// ByteVectorDeltaDecoder performs a read/optional read of a FAST encoded byte vector delta
type ByteVectorDeltaDecoder struct {
}

// ReadValue fast encoded byte vector delta
func (ByteVectorDeltaDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	subtractionLength, err := ReadInt32(inputSource)
	if err != nil {
		return nil, err
	}

	byteVector, err := ReadByteVector(inputSource)
	if err != nil {
		return nil, err
	}
	byteVector.ItemsToRemove = subtractionLength.Value
	return byteVector, nil
}

// ReadOptionalValue fast encoded byte vector delta
func (ByteVectorDeltaDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	subtractionLength, err := ReadOptionalInt32(inputSource)
	if err != nil {
		return nil, err
	}

	// if no subtraction length present, then no delta encoded
	switch t := subtractionLength.(type) {
	case value.NullValue:
		return t, nil
	}

	byteVector, err := ReadByteVector(inputSource)
	if err != nil {
		return nil, err
	}
	byteVector.ItemsToRemove = subtractionLength.(value.Int32Value).Value
	return byteVector, nil
}
