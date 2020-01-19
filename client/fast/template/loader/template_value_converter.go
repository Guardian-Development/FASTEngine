package loader

import "strconv"

type valueConverter func(string) (interface{}, error)

func toString(value string) (interface{}, error) {
	return value, nil
}

func toInt32(value string) (interface{}, error) {
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}
	return int32(val), err
}

func toUInt32(value string) (interface{}, error) {
	val, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return nil, err
	}
	return uint32(val), err
}

func toInt64(value string) (interface{}, error) {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return int64(val), err
}

func toUInt64(value string) (interface{}, error) {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return uint64(val), err
}
