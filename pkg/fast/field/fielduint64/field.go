package fielduint64

import (
	"bytes"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/decoder"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/operation"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fast/value"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// FieldUInt64 represents a FAST template <uInt64/> type
type FieldUInt64 struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise an <uint64/> from the input source
func (field FieldUInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
	previousValue := dictionary.GetValue(field.FieldDetails.Name)
	if field.Operation.ShouldReadValue(pMap) {
		var readValue value.Value
		var err error

		if field.FieldDetails.Required {
			readValue, err = field.decode.ReadValue(inputSource)
		} else {
			readValue, err = field.decode.ReadOptionalValue(inputSource)
		}

		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldUInt64][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
			return nil, fmt.Errorf("[FieldUInt64][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
		}

		transformedValue, err := field.Operation.Apply(readValue, previousValue)
		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldUInt64][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, readValue, previousValue, err)
			return nil, fmt.Errorf("[FieldUInt64][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, readValue, previousValue, err)
		}

		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		field.FieldDetails.Logger.Printf("[FieldUInt64][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
		return nil, fmt.Errorf("[FieldUInt64][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
	}

	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldUInt64) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldUInt64) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <uint64/> field with the given properties and no operation
func New(properties properties.Properties) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <uint64/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue uint64) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <uint64/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <uint64/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue uint64) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <uint64/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <uint64/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue uint64) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewIncrementOperation <uint64/> field with the given properties and <increment/> operator
func NewIncrementOperation(properties properties.Properties) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewIncrementOperationWithInitialValue <uint64/> field with the given properties and <increment value="initialValue"/> operator
func NewIncrementOperationWithInitialValue(properties properties.Properties, initialValue uint64) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.UInt64Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewDeltaOperation <uint64/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.BitIntDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(uint64(0)),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <uint64/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue uint64) FieldUInt64 {
	field := FieldUInt64{
		FieldDetails: properties,
		decode:       decoder.BitIntDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(uint64(0)),
		},
	}

	return field
}
