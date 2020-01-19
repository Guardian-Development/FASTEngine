package fast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// ReadUInt32 reads the next FAST encoded value off the inputSource, treating it as a uint32 value. If the next value would overflow a uint32 an err is returned.
// i.e. 00010010 10001000 would become 100100001000
func ReadUInt32(inputSource *bytes.Buffer) (value.UInt32Value, error) {
	var readValue uint32 = 0

	for i := 0; i < 4; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.UInt32Value{}, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := uint32(b & 127)
			readValue = readValue<<7 | removedStopBit
			return value.UInt32Value{Value: readValue}, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the uint we are reading
		readValue = readValue<<7 | uint32(b)
	}

	return value.UInt32Value{}, fmt.Errorf("More than 4 bytes have been read without reading a stop bit, this will overflow a uint32")
}

// ReadOptionalUInt32 reads a uint64 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned. Due to needing to use 0 to.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalUInt32(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadUInt32(inputSource)
	if err != nil {
		return nil, err
	}

	if readValue.Value == uint32(0) {
		return value.NullValue{}, nil
	}

	readValue.Value = readValue.Value - 1

	return readValue, nil
}

// ReadInt32 reads the next FAST encoded value off the inputSource, treating it as an int32 value (2's compliment encoded). If the next value would overflow a int32 an err is returned.
// i.e. 11111111 01001110 would become 11111111001110 -> 11001110 -> -50
func ReadInt32(inputSource *bytes.Buffer) (value.Int32Value, error) {
	var readValue int32 = 0

	b, err := inputSource.ReadByte()
	if err != nil {
		return value.Int32Value{}, err
	}

	// 64 = 01000000, indicating this is negative so we should start with all 1's int32 (-1)
	if isNegative := b & 64; isNegative == 64 {
		readValue = -1
	}

	// reset byte buffer by the one byte we had to read to determine negative/positive number
	err = inputSource.UnreadByte()
	if err != nil {
		return value.Int32Value{}, err
	}

	for i := 0; i < 4; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.Int32Value{}, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := int32(b & 127)
			readValue = readValue<<7 | removedStopBit
			return value.Int32Value{Value: readValue}, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the int we are reading
		readValue = readValue<<7 | int32(b)
	}

	return value.Int32Value{}, fmt.Errorf("More than 4 bytes have been read without reading a stop bit, this will overflow an int32")
}

// ReadOptionalInt32 reads an int32 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1 for positive numbers only.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalInt32(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadInt32(inputSource)
	if err != nil {
		return value.Int32Value{}, err
	}

	if readValue.Value == int32(0) {
		return value.NullValue{}, nil
	}

	if readValue.Value > 0 {
		readValue.Value = readValue.Value - 1
	}

	return readValue, nil
}

// ReadUInt64 reads the next FAST encoded value off the inputSource, treating it as a uint64 value. If the next value would overflow a uint64 an err is returned.
// i.e. 00010010 10001000 would become 100100001000
func ReadUInt64(inputSource *bytes.Buffer) (value.UInt64Value, error) {
	var readValue uint64 = 0

	for i := 0; i < 8; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.UInt64Value{}, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := uint64(b & 127)
			readValue = readValue<<7 | removedStopBit
			return value.UInt64Value{Value: readValue}, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the uint we are reading
		readValue = readValue<<7 | uint64(b)
	}

	return value.UInt64Value{}, fmt.Errorf("More than 8 bytes have been read without reading a stop bit, this will overflow a uint64")
}

// ReadOptionalUInt64 reads a uint64 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalUInt64(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadUInt64(inputSource)
	if err != nil {
		return value.UInt64Value{}, err
	}

	if readValue.Value == uint64(0) {
		return value.NullValue{}, nil
	}

	readValue.Value = readValue.Value - 1

	return readValue, nil
}

// ReadString reads an ASCII encoded string off the buffer. This can be done as ASCII is a subset of UTF-8 which is what GO uses to represent strings.
func ReadString(inputSource *bytes.Buffer) (value.StringValue, error) {
	stringBuilder := strings.Builder{}
	return readString(inputSource, &stringBuilder)
}

// ReadOptionalString reads an ASCII encoded string off the buffer. If the first value is 10000000, this is seen as null. If the first values are
// 00000000 10000000 this is seen as an empty string.
func ReadOptionalString(inputSource *bytes.Buffer) (value.Value, error) {
	possibleNullIndiciator, err := inputSource.ReadByte()
	if err != nil {
		return value.StringValue{}, err
	}

	// 128 = 10000000, this is seen as null in optional string
	if possibleNullIndiciator == 128 {
		return value.NullValue{}, nil
	}

	possibleEmptyStringIndicator, err := inputSource.ReadByte()
	if err != nil {
		return value.StringValue{}, err
	}

	// 0 = 00000000, 128 = 10000000, this is seen as empty string
	if possibleNullIndiciator == 0 && possibleEmptyStringIndicator == 128 {
		return value.StringValue{Value: ""}, nil
	}

	stringBuilder := strings.Builder{}
	appendNotNullChar(possibleNullIndiciator, &stringBuilder)
	appendNotNullChar(possibleEmptyStringIndicator, &stringBuilder)

	return readString(inputSource, &stringBuilder)
}

func readString(inputSource *bytes.Buffer, stringBuilder *strings.Builder) (value.StringValue, error) {
	for {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.StringValue{}, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := byte(b & 127)
			appendNotNullChar(removedStopBit, stringBuilder)

			return value.StringValue{Value: stringBuilder.String()}, nil
		}

		// no stop bit present so 0 in most significant bit, so just add as 7 bit char to string
		stringBuilder.WriteByte(b)
	}
}

func appendNotNullChar(char byte, stringBuilder *strings.Builder) {
	if char != 0 {
		stringBuilder.WriteByte(char)
	}
}

// ReadValue reads the next FAST encoded value off the inputSource, shifting each value by <<1 to remove the stop bit FAST encoding
// i.e. 00010010 10001000 would become [00100100, 00010000]
func ReadValue(inputSource *bytes.Buffer) ([]byte, error) {
	value := make([]byte, 0)

	for {
		b, err := inputSource.ReadByte()
		if err != nil {
			return nil, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			value = append(value, b<<1)
			return value, nil
		}

		value = append(value, b<<1)
	}
}
