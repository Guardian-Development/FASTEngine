package decoder

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/value"
)

// ReadUInt32 reads the next FAST encoded value off the inputSource, treating it as a uint32 value. If the next value would overflow a uint32 an err is returned.
// i.e. 00010010 10001000 would become 100100001000
func ReadUInt32(inputSource *bytes.Buffer) (value.UInt32Value, error) {
	var readValue uint32 = 0

	for i := 0; i < 5; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.UInt32Value{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
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

	return value.UInt32Value{}, fmt.Errorf("%s, uint32", errors.R6)
}

// ReadOptionalUInt32 reads a uint32 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalUInt32(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadUInt64(inputSource) // allow for overflow
	if err != nil {
		return value.NullValue{}, fmt.Errorf("unable to read value before assesing nullability, reason: %s", err)
	}

	if readValue.Value == uint64(0) {
		return value.NullValue{}, nil
	}

	return value.UInt32Value{Value: uint32(readValue.Value - 1)}, nil
}

// ReadInt32 reads the next FAST encoded value off the inputSource, treating it as an int32 value (2's compliment encoded). If the next value would overflow an int32 an err is returned.
// i.e. 11111111 01001110 would become 11111111001110 -> 11001110 -> -50
func ReadInt32(inputSource *bytes.Buffer) (value.Int32Value, error) {
	var readValue int32 = 0

	b, err := inputSource.ReadByte()
	if err != nil {
		return value.Int32Value{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
	}

	// 64 = 01000000, indicating this is negative so we should start with all 1's int32 (-1)
	if isNegative := b & 64; isNegative == 64 {
		readValue = -1
	}

	// reset byte buffer by the one byte we had to read to determine negative/positive number
	err = inputSource.UnreadByte()
	if err != nil {
		return value.Int32Value{}, fmt.Errorf("unable to rewind byte buffer, reason: %s", err)
	}

	for i := 0; i < 5; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.Int32Value{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
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

	return value.Int32Value{}, fmt.Errorf("%s, int32", errors.R6)
}

// ReadOptionalInt32 reads an int32 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1 for positive numbers only.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalInt32(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadInt64(inputSource) // allow for overflow
	if err != nil {
		return value.Int32Value{}, fmt.Errorf("unable to read value before assesing nullability, reason: %s", err)
	}

	if readValue.Value == int64(0) {
		return value.NullValue{}, nil
	}

	if readValue.Value > 0 {
		readValue.Value = readValue.Value - 1
	}

	return value.Int32Value{Value: int32(readValue.Value)}, nil
}

// ReadUInt64 reads the next FAST encoded value off the inputSource, treating it as a uint64 value. If the next value would overflow a uint64 an err is returned.
// i.e. 00010010 10001000 would become 100100001000
func ReadUInt64(inputSource *bytes.Buffer) (value.UInt64Value, error) {
	var readValue uint64 = 0

	for i := 0; i < 10; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.UInt64Value{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
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

	return value.UInt64Value{}, fmt.Errorf("%s, uint64", errors.R6)
}

// ReadOptionalUInt64 reads a uint64 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalUInt64(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadBigUInt(inputSource)
	if err != nil {
		return value.UInt64Value{}, fmt.Errorf("unable to read value before assesing nullability, reason: %s", err)
	}

	equalToZero := readValue.Value.Cmp(big.NewInt(0))
	if equalToZero == 0 {
		return value.NullValue{}, nil
	}

	readValue.Value = readValue.Value.Sub(readValue.Value, big.NewInt(1))

	if readValue.Value.IsUint64() {
		return value.UInt64Value{Value: readValue.Value.Uint64()}, nil
	}

	return value.NullValue{}, fmt.Errorf("%s, uint64", errors.R6)
}

// ReadInt64 reads the next FAST encoded value off the inputSource, treating it as an int64 value (2's compliment encoded). If the next value would overflow an int64 an err is returned.
// i.e. 11111111 01001110 would become 11111111001110 -> 11001110 -> -50
func ReadInt64(inputSource *bytes.Buffer) (value.Int64Value, error) {
	var readValue int64 = 0

	b, err := inputSource.ReadByte()
	if err != nil {
		return value.Int64Value{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
	}

	// 64 = 01000000, indicating this is negative so we should start with all 1's int64 (-1)
	if isNegative := b & 64; isNegative == 64 {
		readValue = -1
	}

	// reset byte buffer by the one byte we had to read to determine negative/positive number
	err = inputSource.UnreadByte()
	if err != nil {
		return value.Int64Value{}, fmt.Errorf("unable to rewind byte buffer, reason: %s", err)
	}

	for i := 0; i < 10; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.Int64Value{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := int64(b & 127)
			readValue = readValue<<7 | removedStopBit
			return value.Int64Value{Value: readValue}, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the int we are reading
		readValue = readValue<<7 | int64(b)
	}

	return value.Int64Value{}, fmt.Errorf("%s, int64", errors.R6)
}

// ReadOptionalInt64 reads an int64 off the buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1 for positive numbers only.
// i.e. 10000000 would become nil, 10000001 would become 0
func ReadOptionalInt64(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadBigInt(inputSource) // allow for overflow
	if err != nil {
		return value.Int64Value{}, fmt.Errorf("unable to read value before assesing nullability, reason: %s", err)
	}

	equalToZero := readValue.Value.Cmp(big.NewInt(0))
	if equalToZero == 0 {
		return value.NullValue{}, nil
	}

	if equalToZero > 0 {
		readValue.Value = readValue.Value.Sub(readValue.Value, big.NewInt(1))
	}

	if readValue.Value.IsInt64() {
		return value.Int64Value{Value: readValue.Value.Int64()}, nil
	}

	return value.NullValue{}, fmt.Errorf("%s, int64", errors.R6)
}

// ReadBigUInt reads the next FAST encoded value off the inputSource, treating it as an uint64 value. However, this value may overflow an uint64 by 1 byte (for delta encoding)
// and therefore we can return a value.BigInt if this happens. The least significant byte is in the overflow value. If the next value would till overflow this structure an err is returned.
func ReadBigUInt(inputSource *bytes.Buffer) (value.BigInt, error) {
	var readValue uint64 = 0

	i := 1
	for ; i < 10; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.BigInt{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := uint64(b & 127)
			readValue = readValue<<7 | removedStopBit
			return value.BigInt{Value: big.NewInt(0).SetUint64(readValue)}, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the int we are reading
		readValue = readValue<<7 | uint64(b)
	}

	// we try to read one more byte (for the overflow), if this does not read a stop bit we return an error
	if i == 10 {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.BigInt{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := int64(b & 127)
			resultBig := big.NewInt(0).SetUint64(readValue)
			resultBig = resultBig.Lsh(resultBig, 7)                         // readValue << 7
			resultBig = resultBig.Or(resultBig, big.NewInt(removedStopBit)) // result | removedStopBit
			return value.BigInt{Value: resultBig}, nil
		}
	}

	return value.BigInt{}, fmt.Errorf("%s, uint64", errors.R6)
}

// ReadBigInt reads the next FAST encoded value off the inputSource, treating it as an int64 value. However, this value may overflow an int64 by 1 byte (for delta encoding)
// and therefore we can return a value.BigInt if this happens. The least significant byte is in the overflow value. If the next value would till overflow this structure an err is returned.
func ReadBigInt(inputSource *bytes.Buffer) (value.BigInt, error) {
	var readValue int64 = 0

	b, err := inputSource.ReadByte()
	if err != nil {
		return value.BigInt{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
	}

	// 64 = 01000000, indicating this is negative so we should start with all 1's int64 (-1)
	if isNegative := b & 64; isNegative == 64 {
		readValue = -1
	}

	// reset byte buffer by the one byte we had to read to determine negative/positive number
	err = inputSource.UnreadByte()
	if err != nil {
		return value.BigInt{}, fmt.Errorf("unable to rewind byte buffer, reason: %s", err)
	}

	i := 1
	for ; i < 10; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.BigInt{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := int64(b & 127)
			readValue = readValue<<7 | removedStopBit
			return value.BigInt{Value: big.NewInt(readValue)}, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the int we are reading
		readValue = readValue<<7 | int64(b)
	}

	// we try to read one more byte (for the overflow), if this does not read a stop bit we return an error
	if i == 10 {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.BigInt{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := int64(b & 127)
			resultBig := big.NewInt(readValue)
			resultBig = resultBig.Lsh(resultBig, 7)                         // readValue << 7
			resultBig = resultBig.Or(resultBig, big.NewInt(removedStopBit)) // result | removedStopBit
			return value.BigInt{Value: resultBig}, nil
		}
	}

	return value.BigInt{}, fmt.Errorf("%s, int64", errors.R6)
}

// ReadOptionalBigInt reads a value.BigInt off the input buffer. If the value returned is 0, this is marked as nil, and nil is returned.
// Due to needing to use 0 as a nil value for optionals, the value returned by this is: value - 1 for positive numbers only.
func ReadOptionalBigInt(inputSource *bytes.Buffer) (value.Value, error) {
	readValue, err := ReadBigInt(inputSource)
	if err != nil {
		return value.BigInt{}, fmt.Errorf("unable to read value before assesing nullability, reason: %s", err)
	}

	equalToZero := readValue.Value.Cmp(big.NewInt(0))
	if equalToZero == 0 {
		return value.NullValue{}, nil
	}

	if equalToZero > 0 {
		readValue.Value = readValue.Value.Sub(readValue.Value, big.NewInt(1))
	}

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
		return value.StringValue{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
	}

	// 128 = 10000000, this is seen as null in optional string
	if possibleNullIndiciator == 128 {
		return value.NullValue{}, nil
	}

	stringBuilder := strings.Builder{}

	// if this is the end of the string (its a 1 byte string) return result
	if result := possibleNullIndiciator & 128; result == 128 {
		removedStopBit := byte(possibleNullIndiciator & 127)
		appendNotNullChar(removedStopBit, &stringBuilder)

		return value.StringValue{Value: stringBuilder.String()}, nil
	}

	possibleEmptyStringIndicator, err := inputSource.ReadByte()
	if err != nil {
		return value.StringValue{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
	}

	// if this is the end of the string (its a 2 byte string) return result
	if result := possibleEmptyStringIndicator & 128; result == 128 {
		appendNotNullChar(byte(possibleNullIndiciator&127), &stringBuilder)
		appendNotNullChar(byte(possibleEmptyStringIndicator&127), &stringBuilder)

		return value.StringValue{Value: stringBuilder.String()}, nil
	}

	// 0 = 00000000, 128 = 10000000, this is seen as empty string
	if possibleNullIndiciator == 0 && possibleEmptyStringIndicator == 128 {
		return value.StringValue{Value: ""}, nil
	}

	// not an empty or null string, append bytes read and read rest of string as normal
	appendNotNullChar(possibleNullIndiciator, &stringBuilder)
	appendNotNullChar(possibleEmptyStringIndicator, &stringBuilder)

	return readString(inputSource, &stringBuilder)
}

func readString(inputSource *bytes.Buffer, stringBuilder *strings.Builder) (value.StringValue, error) {
	for {
		b, err := inputSource.ReadByte()
		if err != nil {
			return value.StringValue{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
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

// ReadByteVector reads a uint32 length off the buffer which represents the length of the vector to then read. The vector read is not stop bit encoded.
// i.e. 10000010 00000001 00000010 would become (length 2) -> [1, 2]
func ReadByteVector(inputSource *bytes.Buffer) (value.ByteVector, error) {
	length, err := ReadUInt32(inputSource)
	if err != nil {
		return value.ByteVector{}, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
	}

	byteVector := make([]byte, length.Value)
	number, err := inputSource.Read(byteVector)
	if err != nil {
		return value.ByteVector{}, fmt.Errorf("unable to read multiple bytes [%d] off byte buffer, reason: %s", length.Value, err)
	}
	if number != int(length.Value) {
		return value.ByteVector{}, fmt.Errorf("did not read full length of byte vector, expected to read: %d, but actually read %d", length.Value, number)
	}

	return value.ByteVector{Value: byteVector}, nil
}

// ReadOptionalByteVector treats the uint32 length preamble as an optional uint32, reading 0 as a null marker. The byte vector itself is read as long as the
// length is not null.
// i.e. 10000010 00000001 would become (length 1) -> [1]
// i.e. 10000000 would become 0, and be marked as null
func ReadOptionalByteVector(inputSource *bytes.Buffer) (value.Value, error) {
	length, err := ReadOptionalUInt32(inputSource)
	if err != nil {
		return nil, fmt.Errorf("unable to read value before assesing nullability, reason: %s", err)
	}

	switch t := length.(type) {
	case value.NullValue:
		return t, nil
	case value.UInt32Value:
		byteVector := make([]byte, t.Value)
		number, err := inputSource.Read(byteVector)
		if err != nil {
			return value.ByteVector{}, fmt.Errorf("unable to read multiple bytes [%d] off byte buffer, reason: %s", t.Value, err)
		}
		if number != int(t.Value) {
			return value.ByteVector{}, fmt.Errorf("did not read full length of byte vector, expected to read: %d, but actually read %d", t.Value, number)
		}
		return value.ByteVector{Value: byteVector}, nil
	default:
		return value.ByteVector{}, fmt.Errorf("unsupported type returned from reading optional uint32 as length of byte vector")
	}
}

// ReadValue reads the values off the byte buffer until a stop but is detected. Stop bits are not removed from the bytes returned.
func ReadValue(inputSource *bytes.Buffer) ([]byte, error) {
	readValue := make([]byte, 0)

	for {
		b, err := inputSource.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unable to read byte off byte buffer, reason: %s", err)
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			readValue = append(readValue, b)
			return readValue, nil
		}

		readValue = append(readValue, b)
	}
}
