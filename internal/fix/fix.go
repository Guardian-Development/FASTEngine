package fix

type Value interface {
	Get() interface{}
}

type RawValue struct {
	value interface{}
}

type NullValue struct {
}

func (nullValue NullValue) Get() interface{} {
	return nil
}

func (rawValue RawValue) Get() interface{} {
	return rawValue.value
}

func NewRawValue(value interface{}) Value {
	if value == nil {
		return NullValue{}
	}

	return RawValue{value: value}
}
