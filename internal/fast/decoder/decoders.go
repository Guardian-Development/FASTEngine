package decoder

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

//TODO: tests in this package should use these structs not the underlying methods themselves

type Int64Decoder struct {
}

func (Int64Decoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadInt64(inputSource)
}

func (Int64Decoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalInt64(inputSource)
}

type BitIntDecoder struct {
}

func (BitIntDecoder) ReadValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadBigInt(inputSource)
}

func (BitIntDecoder) ReadOptionalValue(inputSource *bytes.Buffer) (value.Value, error) {
	return ReadOptionalBigInt(inputSource)
}
