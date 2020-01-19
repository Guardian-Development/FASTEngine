package field

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// Field contains information about a TemplateUnit within a FAST Template
type Field struct {
	ID        uint64
	Operation operation.Operation
	Required  bool
}

// String represents a FAST template <string/> type
type String struct {
	FieldDetails Field
}

func (field String) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadString(inputSource)
		} else {
			value, err = fast.ReadOptionalString(inputSource)
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
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
	if field.FieldDetails.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadUInt32(inputSource)
		} else {
			value, err = fast.ReadOptionalUInt32(inputSource)
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// Int32 represents a FAST template <int32/> type
type Int32 struct {
	FieldDetails Field
}

func (field Int32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadInt32(inputSource)
		} else {
			value, err = fast.ReadOptionalInt32(inputSource)
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
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
	if field.FieldDetails.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadUInt64(inputSource)
		} else {
			value, err = fast.ReadOptionalUInt64(inputSource)
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// Int64 represents a FAST template <int64/> type
type Int64 struct {
	FieldDetails Field
}

func (field Int64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadInt64(inputSource)
		} else {
			value, err = fast.ReadOptionalInt64(inputSource)
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}
