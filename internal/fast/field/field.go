package field

import (
	"bytes"
	"fmt"
	"math"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
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

// Sequence represents a FAST template <sequence /> type
type Sequence struct {
	FieldDetails   Field
	LengthField    UInt32
	SequenceFields []store.Unit
}

func (field Sequence) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	numberOfElements, err := field.LengthField.Deserialise(inputSource, pMap)
	if err != nil {
		return nil, err
	}

	switch t := numberOfElements.(type) {
	case fix.NullValue:
		return t, nil
	}

	elementCount := numberOfElements.Get().(uint32)
	sequenceValue := fix.NewSequenceValue(elementCount)

	for elementNumber := uint32(0); elementNumber < elementCount; elementNumber++ {
		sequencePmap := presencemap.PresenceMap{}
		if field.RequiresPmap() {
			sequencePmap, err = presencemap.New(inputSource)
			if err != nil {
				return nil, err
			}
		}

		for _, element := range field.SequenceFields {
			value, err := element.Deserialise(inputSource, &sequencePmap)
			if err != nil {
				return nil, err
			}
			sequenceValue.SetValue(elementNumber, element.GetTagId(), value)
		}
	}

	return sequenceValue, nil
}

func (field Sequence) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field Sequence) RequiresPmap() bool {
	for _, element := range field.SequenceFields {
		if element.RequiresPmap() {
			return true
		}
	}

	return field.LengthField.RequiresPmap()
}

// AsciiString represents a FAST template <string charset="ascii"/> type and <string /> type
type AsciiString struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field AsciiString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field AsciiString) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field AsciiString) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// UnicodeString represents a FAST template <string charset="unicode"/> type
type UnicodeString struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UnicodeString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field UnicodeString) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field UnicodeString) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// UInt32 represents a FAST template <uInt32/> type
type UInt32 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UInt32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field UInt32) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field UInt32) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// Int32 represents a FAST template <int32/> type
type Int32 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field Int32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field Int32) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field Int32) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// UInt64 represents a FAST template <uInt64/> type
type UInt64 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field UInt64) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field UInt64) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// Int64 represents a FAST template <int64/> type
type Int64 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field Int64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field Int64) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field Int64) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
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

func (field Decimal) RequiresPmap() bool {
	return field.ExponentField.RequiresPmap() || field.MantissaField.RequiresPmap()
}

// ByteVector represents a FAST template <byteVector/> type
type ByteVector struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field ByteVector) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap) (fix.Value, error) {
	if field.Operation.ShouldReadValue() {
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

	return field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required)
}

func (field ByteVector) GetTagId() uint64 {
	return field.FieldDetails.ID
}

func (field ByteVector) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}
