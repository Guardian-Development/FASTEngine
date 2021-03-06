package structure

import "github.com/Guardian-Development/fastengine/internal/xml"

const TemplatesTag = "templates"
const TemplateTag = "template"
const StringTag = "string"
const UInt32Tag = "uInt32"
const Int32Tag = "int32"
const UInt64Tag = "uInt64"
const Int64Tag = "int64"
const ByteVectorTag = "byteVector"
const SequenceTag = "sequence"
const LengthTag = "length"
const DecimalTag = "decimal"
const UnicodeStringLabel = "unicode"

const ConstantOperation = "constant"
const DefaultOperation = "default"
const CopyOperation = "copy"
const IncrementOperation = "increment"
const TailOperation = "tail"
const DeltaOperation = "delta"

const ValueAttribute = "value"

// HasValue returns whether the value attribute is set on the xml tags
func HasValue(tagInTemplate *xml.Tag) bool {
	return tagInTemplate.Attributes[ValueAttribute] != ""
}

// IsNullString returns whether the value is equal to ""
func IsNullString(value string) bool {
	return value == ""
}
