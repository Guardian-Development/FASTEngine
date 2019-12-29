package fast

import (
	"bytes"
	"fmt"
)

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
