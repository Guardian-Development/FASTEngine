package value

type Value interface {
}

type StringValue struct {
	Value string
}

type UInt32Value struct {
	Value uint32
}

type Int32Value struct {
	Value int32
}

type UInt64Value struct {
	Value uint64
}

type Int64Value struct {
	Value int64
}

type NullValue struct {
}
