package field

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/client/fast/fix"
	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

// Field contains information about a TemplateUnit within a FAST Template
type Field struct {
	ID        uint64
	Operation operation.Operation
}

// String represents a FAST template <string/> type
type String struct {
	FieldDetails Field
}

func (field String) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap) {
		value, err := fast.ReadString(inputSource)
		// TODO: apply operation on value
		fixContext.SetTag(field.FieldDetails.ID, value)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// UInt32 represents a FAST template <uInt32/> type
type UInt32 struct {
	FieldDetails Field
}

func (field UInt32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap) {
		value, err := fast.ReadUInt32(inputSource)
		// TODO: apply operation on value
		fixContext.SetTag(field.FieldDetails.ID, value)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// UInt64 represents a FAST template <uInt64/> type
type UInt64 struct {
	FieldDetails Field
}

func (field UInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap) {
		value, err := fast.ReadUInt64(inputSource)
		// TODO: apply operation on value
		fixContext.SetTag(field.FieldDetails.ID, value)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}
