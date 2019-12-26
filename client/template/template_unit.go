package template

const stringTag = "string"
const uInt32Tag = "uInt32"
const uInt64Tag = "uInt64"

// Unit represents an element within a FAST Template, with the ability to Serialise/Deserialise a part of a FAST message
type Unit interface {
	Deserialise(inputSource []byte)
}

// Field contains information about a TemplateUnit within a FAST Template
type Field struct {
	ID uint64
}

// FieldString represents a FAST template <string/> type
type FieldString struct {
	fieldDetails Field
}

func (field FieldString) Deserialise(inputSource []byte) {
}

// FieldUInt32 represents a FAST template <uInt32/> type
type FieldUInt32 struct {
	fieldDetails Field
}

func (field FieldUInt32) Deserialise(inputSource []byte) {
}

// FieldUInt64 represents a FAST template <uInt64/> type
type FieldUInt64 struct {
	fieldDetails Field
}

func (field FieldUInt64) Deserialise(inputSource []byte) {
}
