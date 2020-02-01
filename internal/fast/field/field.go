package field

import (
	"bytes"
	"fmt"
	"math"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// Field contains information about a TemplateUnit within a FAST Template
type Field struct {
	ID       uint64
	Required bool
}

// String represents a FAST template <string charset="ascii"/> type
type String struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field String) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadString(inputSource)
		} else {
			value, err = fast.ReadOptionalString(inputSource)
		}

		if err != nil {
			return nil, err
		}

		return field.Operation.Apply(value)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field String) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// UnicodeString represents a FAST template <string charset="unicode"/> type
type UnicodeString struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UnicodeString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var stringValue value.Value
		var err error

		if field.FieldDetails.Required {
			stringValue, err = fast.ReadByteVector(inputSource)
		} else {
			stringValue, err = fast.ReadOptionalByteVector(inputSource)
		}

		if err != nil {
			return nil, err
		}

		switch t := stringValue.(type) {
		case value.ByteVector:
			stringValue = value.StringValue{Value: string(t.Value)}
		}

		return field.Operation.Apply(stringValue)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field UnicodeString) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// UInt32 represents a FAST template <uInt32/> type
type UInt32 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UInt32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadUInt32(inputSource)
		} else {
			value, err = fast.ReadOptionalUInt32(inputSource)
		}

		if err != nil {
			return nil, err
		}

		return field.Operation.Apply(value)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field UInt32) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// Int32 represents a FAST template <int32/> type
type Int32 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field Int32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadInt32(inputSource)
		} else {
			value, err = fast.ReadOptionalInt32(inputSource)
		}

		if err != nil {
			return nil, err
		}

		return field.Operation.Apply(value)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field Int32) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// UInt64 represents a FAST template <uInt64/> type
type UInt64 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadUInt64(inputSource)
		} else {
			value, err = fast.ReadOptionalUInt64(inputSource)
		}

		if err != nil {
			return nil, err
		}

		return field.Operation.Apply(value)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field UInt64) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// Int64 represents a FAST template <int64/> type
type Int64 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field Int64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadInt64(inputSource)
		} else {
			value, err = fast.ReadOptionalInt64(inputSource)
		}

		if err != nil {
			return nil, err
		}

		return field.Operation.Apply(value)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field Int64) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// Decimal represents a FAST template <decimal/> type
type Decimal struct {
	FieldDetails  Field
	ExponentField Int32
	MantissaField Int64
}

func (field Decimal) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	exponentValue, err := field.ExponentField.Deserialise(inputSource, pMap)
	if err != nil {
		return nil, err
	}
	switch exponentValue.(type) {
	case fix.NullValue:
		return fix.NullValue{}, nil
	case fix.RawValue:
		mantissaValue, err := field.MantissaField.Deserialise(inputSource, pMap)
		if err != nil {
			return nil, err
		}
		decimalValue := math.Pow(10, float64(exponentValue.Get().(int32))) * float64(mantissaValue.Get().(int64))
		return fix.NewRawValue(decimalValue), nil
	}

	return nil, fmt.Errorf("Exponent value of decimal was not expected type: %v", exponentValue)
}

func (field Decimal) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// ByteVector represents a FAST template <byteVector/> type
type ByteVector struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field ByteVector) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadByteVector(inputSource)
		} else {
			value, err = fast.ReadOptionalByteVector(inputSource)
		}

		if err != nil {
			return nil, err
		}

		return field.Operation.Apply(value)
	}

	return field.Operation.GetNotEncodedValue()
}

func (field ByteVector) GetTagId() uint64 {
	return field.FieldDetails.ID
}
