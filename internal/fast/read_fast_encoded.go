package fast

import (
	"bytes"
	"fmt"
)

// ReadUInt32 reads the next FAST encoded value off the inputSource, treating it as a uint32 value. If the next value would overflow a uint32 an err is returned.
// i.e. 00010010 10001000 would become 100100001000
func ReadUInt32(inputSource *bytes.Buffer) (uint32, error) {
	var value uint32 = 0

	for i := 0; i < 4; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return 0, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := uint32(b & 127)
			return value<<7 | removedStopBit, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the uint we are reading
		value = value<<7 | uint32(b)
	}

	return 0, fmt.Errorf("More than 4 bytes have been read without reading a stop bit, this will overflow a uint32")
}

// ReadUInt64 reads the next FAST encoded value off the inputSource, treating it as a uint64 value. If the next value would overflow a uint64 an err is returned.
// i.e. 00010010 10001000 would become 100100001000
func ReadUInt64(inputSource *bytes.Buffer) (uint64, error) {
	var value uint64 = 0

	for i := 0; i < 8; i++ {
		b, err := inputSource.ReadByte()
		if err != nil {
			return 0, err
		}

		// 128 = 10000000, this will equal 128 if we have a stop bit present (most significant bit is 1)
		if result := b & 128; result == 128 {
			removedStopBit := uint64(b & 127)
			return value<<7 | removedStopBit, nil
		}

		// no stop bit present so 0 in most significant bit, add this byte to the uint we are reading
		value = value<<7 | uint64(b)
	}

	return 0, fmt.Errorf("More than 8 bytes have been read without reading a stop bit, this will overflow a uint64")
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

//TODO: implement reading a string
func ReadString(inputSource *bytes.Buffer) (string, error) {
	return "", nil
}
