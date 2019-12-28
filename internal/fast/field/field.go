package field

import "github.com/Guardian-Development/fastengine/internal/fast/operation"

// Field contains information about a TemplateUnit within a FAST Template
type Field struct {
	ID        uint64
	Operation operation.Operation
}

// String represents a FAST template <string/> type
type String struct {
	FieldDetails Field
}

func (field String) Deserialise(inputSource []byte) {
}

// UInt32 represents a FAST template <uInt32/> type
type UInt32 struct {
	FieldDetails Field
}

func (field UInt32) Deserialise(inputSource []byte) {
}

// UInt64 represents a FAST template <uInt64/> type
type UInt64 struct {
	FieldDetails Field
}

func (field UInt64) Deserialise(inputSource []byte) {
}
