package fieldbytevector

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

// FieldByteVector represents a FAST template <byteVector/> type
type FieldByteVector struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise a <byteVector/> from the input source
func (field FieldByteVector) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
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
			field.FieldDetails.Logger.Printf("[FieldByteVector][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
			return nil, fmt.Errorf("[FieldByteVector][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
		}

		transformedValue, err := field.Operation.Apply(value, previousValue)
		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldByteVector][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, value, previousValue, err)
			return nil, fmt.Errorf("[FieldByteVector][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, value, previousValue, err)
		}

		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		field.FieldDetails.Logger.Printf("[FieldByteVector][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
		return nil, fmt.Errorf("[FieldByteVector][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
	}

	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldByteVector) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldByteVector) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <byteVector/> field with the given properties and no operation
func New(properties properties.Properties) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <byteVector/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue []byte) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <byteVector/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <byteVector/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue []byte) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <byteVector/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <byteVector/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue []byte) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewTailOperation <byteVector/> field with the given properties and <tail/> operator
func NewTailOperation(properties properties.Properties) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Tail{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue([]byte{}),
		},
	}
	return field
}

// NewTailOperationWithInitialValue <byteVector/> field with the given properties and <tail value="initialValue"/> operator
func NewTailOperationWithInitialValue(properties properties.Properties, initialValue []byte) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Tail{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue([]byte{}),
		},
	}
	return field
}

// NewDeltaOperation <byteVector/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDeltaDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue([]byte{}),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <byteVector/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue []byte) FieldByteVector {
	field := FieldByteVector{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDeltaDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue([]byte{}),
		},
	}

	return field
}
