package fielduint32

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/decoder"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/operation"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fast/value"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// FieldUInt32 represents a FAST template <uInt32/> type
type FieldUInt32 struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise an <uint32/> from the input source
func (field FieldUInt32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
	previousValue := dictionary.GetValue(field.FieldDetails.Name)
	if field.Operation.ShouldReadValue(pMap) {
		var value value.Value
		var err error

		if field.FieldDetails.Required {
			value, err = field.decode.ReadValue(inputSource)
		} else {
			value, err = field.decode.ReadOptionalValue(inputSource)
		}

		if err != nil {
			return nil, err
		}

		transformedValue, err := field.Operation.Apply(value, previousValue)
		if err != nil {
			return nil, err
		}
		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		return nil, err
	}
	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldUInt32) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldUInt32) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <uint32/> field with the given properties and no operation
func New(properties properties.Properties) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <uint32/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue uint32) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <uint32/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <uint32/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue uint32) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <uint32/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <uint32/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue uint32) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewIncrementOperation <uint32/> field with the given properties and <increment/> operator
func NewIncrementOperation(properties properties.Properties) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewIncrementOperationWithInitialValue <uint32/> field with the given properties and <increment value="initialValue"/> operator
func NewIncrementOperationWithInitialValue(properties properties.Properties, initialValue uint32) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.UInt32Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewDeltaOperation <uint32/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(uint32(0)),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <uint32/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue uint32) FieldUInt32 {
	field := FieldUInt32{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(uint32(0)),
		},
	}

	return field
}
