package fieldint64

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

// FieldInt64 represents a FAST template <int64/> type
type FieldInt64 struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise an <int64/> from the input source
func (field FieldInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
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
			field.FieldDetails.Logger.Printf("[FieldInt64][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
			return nil, fmt.Errorf("[FieldInt64][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
		}

		transformedValue, err := field.Operation.Apply(readValue, previousValue)
		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldInt64][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, readValue, previousValue, err)
			return nil, fmt.Errorf("[FieldInt64][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, readValue, previousValue, err)
		}

		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		field.FieldDetails.Logger.Printf("[FieldInt64][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
		return nil, fmt.Errorf("[FieldInt64][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
	}

	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldInt64) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldInt64) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <int64/> field with the given properties and no operation
func New(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <int64/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <int64/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <int64/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <int64/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <int64/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewIncrementOperation <int64/> field with the given properties and <increment/> operator
func NewIncrementOperation(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewIncrementOperationWithInitialValue <int64/> field with the given properties and <increment value="initialValue"/> operator
func NewIncrementOperationWithInitialValue(properties properties.Properties, initialValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.Int64Decoder{},
		Operation: operation.Increment{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewDeltaOperation <int64/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.BitIntDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(int64(0)),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <int64/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		decode:       decoder.BitIntDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(int64(0)),
		},
	}

	return field
}
