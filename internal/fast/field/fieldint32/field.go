package fieldint32

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/internal/fast/decoder"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// FieldInt32 represents a FAST template <int32/> type
type FieldInt32 struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise an <int32/> from the input source
func (field FieldInt32) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
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
func (field FieldInt32) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldInt32) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <int32/> field with the given properties and no operation
func New(properties properties.Properties) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <int32/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue int32) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <int32/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <int32/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue int32) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <int32/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <int32/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue int32) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewIncrementOperation <int32/> field with the given properties and <increment/> operator
func NewIncrementOperation(properties properties.Properties) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewIncrementOperationWithInitialValue <int32/> field with the given properties and <increment value="initialValue"/> operator
func NewIncrementOperationWithInitialValue(properties properties.Properties, initialValue int32) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int32Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewDeltaOperation <int32/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(int32(0)),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <int32/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue int32) FieldInt32 {
	field := FieldInt32{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(int32(0)),
		},
	}

	return field
}
