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
	Required  bool
}

// String represents a FAST template <string/> type
type String struct {
	FieldDetails Field
}

func (field String) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.FieldDetails.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var stringValue *string

		if field.FieldDetails.Required {
			value, err := fast.ReadString(inputSource)
			if err != nil {
				return err
			}
			stringValue = &value
		} else {
			value, err := fast.ReadOptionalString(inputSource)
			if err != nil {
				return err
			}
			stringValue = value
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(stringValue)
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
		var uint32Value *uint32

		if field.FieldDetails.Required {
			value, err := fast.ReadUInt32(inputSource)
			if err != nil {
				return err
			}
			uint32Value = &value
		} else {
			value, err := fast.ReadOptionalUInt32(inputSource)
			if err != nil {
				return err
			}
			uint32Value = value
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(uint32Value)
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
		var uint64Value *uint64

		if field.FieldDetails.Required {
			value, err := fast.ReadUInt64(inputSource)
			if err != nil {
				return err
			}
			uint64Value = &value
		} else {
			value, err := fast.ReadOptionalUInt64(inputSource)
			if err != nil {
				return err
			}
			uint64Value = value
		}

		transformedValue, err := field.FieldDetails.Operation.Apply(uint64Value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.FieldDetails.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}
