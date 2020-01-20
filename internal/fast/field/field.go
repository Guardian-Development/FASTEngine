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
	internalFix "github.com/Guardian-Development/fastengine/internal/fix"
)

// Field contains information about a TemplateUnit within a FAST Template
type Field struct {
	ID       uint64
	Required bool
}

// String represents a FAST template <string/> type
type String struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field String) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadString(inputSource)
		} else {
			value, err = fast.ReadOptionalString(inputSource)
		}

		transformedValue, err := field.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// UInt32 represents a FAST template <uInt32/> type
type UInt32 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UInt32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadUInt32(inputSource)
		} else {
			value, err = fast.ReadOptionalUInt32(inputSource)
		}

		transformedValue, err := field.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// Int32 represents a FAST template <int32/> type
type Int32 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field Int32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadInt32(inputSource)
		} else {
			value, err = fast.ReadOptionalInt32(inputSource)
		}

		transformedValue, err := field.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// UInt64 represents a FAST template <uInt64/> type
type UInt64 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field UInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadUInt64(inputSource)
		} else {
			value, err = fast.ReadOptionalUInt64(inputSource)
		}

		transformedValue, err := field.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// Int64 represents a FAST template <int64/> type
type Int64 struct {
	FieldDetails Field
	Operation    operation.Operation
}

func (field Int64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.Operation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = fast.ReadInt64(inputSource)
		} else {
			value, err = fast.ReadOptionalInt64(inputSource)
		}

		transformedValue, err := field.Operation.Apply(value)
		fixContext.SetTag(field.FieldDetails.ID, transformedValue)
		return err
	}

	value, err := field.Operation.GetNotEncodedValue()
	fixContext.SetTag(field.FieldDetails.ID, value)
	return err
}

// TODO: if we refactored deseralise to return the value itself, we can use exponent operations to just be int32 fields instead, which would be better
// TODO: once refactores to hold filed rather than operations this all becomes a lot simpler. do this next.
// TODO: may also need to refactor constant operation to store a fix.Value so can represent decimal correctly (fix.Value will need a decimal type)
type Decimal struct {
	FieldDetails      Field
	ExponentOperation operation.Operation
	MantissaOperation operation.Operation
}

func (field Decimal) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, fixContext *fix.Message) error {
	if field.ExponentOperation.ShouldReadValue(pMap, field.FieldDetails.Required) {
		var exponentValue value.Value
		var err error

		if field.FieldDetails.Required {
			exponentValue, err = fast.ReadInt32(inputSource)
		} else {
			exponentValue, err = fast.ReadOptionalInt32(inputSource)
		}
		transformedExponentValue, err := field.ExponentOperation.Apply(exponentValue)

		decimalValue, err := field.buildDecimalFromExponent(inputSource, transformedExponentValue)
		fixContext.SetTag(field.FieldDetails.ID, decimalValue)
		return err
	}

	value, err := field.ExponentOperation.GetNotEncodedValue()
	decimalValue, err := field.buildDecimalFromExponent(inputSource, value)
	fixContext.SetTag(field.FieldDetails.ID, decimalValue)
	return err
}

func (field Decimal) buildDecimalFromExponent(inputSource *bytes.Buffer, transformedExponentValue internalFix.Value) (internalFix.Value, error) {
	switch transformedExponentValue.(type) {
	case internalFix.NullValue:
		return internalFix.NullValue{}, nil
	case internalFix.RawValue:
		exponentValue := transformedExponentValue.Get().(int32)
		mantissaValue, err := fast.ReadInt64(inputSource)
		transformedMantissaValue, err := field.MantissaOperation.Apply(mantissaValue)

		if err != nil {
			return nil, err
		}

		decimalValue := math.Pow(10, float64(exponentValue)) * float64(transformedMantissaValue.Get().(int64))
		return internalFix.NewRawValue(decimalValue), nil
	}

	return internalFix.NullValue{}, fmt.Errorf("Unsupported exponent type for building decimal: %v", transformedExponentValue)
}
