package fieldunicodestring

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

// FieldUnicodeString represents a FAST template <string charset="unicode"/> type
type FieldUnicodeString struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise a <string charset="unicode"/> from the input source
func (field FieldUnicodeString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
	previousValue := dictionary.GetValue(field.FieldDetails.Name)
	if field.Operation.ShouldReadValue(pMap) {
		var stringValue value.Value
		var err error

		if field.FieldDetails.Required {
			stringValue, err = field.decode.ReadValue(inputSource)
		} else {
			stringValue, err = field.decode.ReadOptionalValue(inputSource)
		}

		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldUnicodeString][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
			return nil, fmt.Errorf("[FieldUnicodeString][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
		}

		switch t := stringValue.(type) {
		case value.ByteVector:
			stringValue = value.StringValue{Value: string(t.Value), ItemsToRemove: t.ItemsToRemove}
		}

		transformedValue, err := field.Operation.Apply(stringValue, previousValue)
		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldUnicodeString][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, stringValue, previousValue, err)
			return nil, fmt.Errorf("[FieldUnicodeString][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, stringValue, previousValue, err)
		}

		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		field.FieldDetails.Logger.Printf("[FieldUnicodeString][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
		return nil, fmt.Errorf("[FieldUnicodeString][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
	}

	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldUnicodeString) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldUnicodeString) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <string charset="unicode"/> field with the given properties and no operation
func New(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <string charset="unicode"/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <string charset="unicode"/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <string charset="unicode"/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <string charset="unicode"/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <string charset="unicode"/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewTailOperation <string charset="unicode"/> field with the given properties and <tail/> operator
func NewTailOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Tail{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(""),
		},
	}
	return field
}

// NewTailOperationWithInitialValue <string charset="unicode"/> field with the given properties and <tail value="initialValue"/> operator
func NewTailOperationWithInitialValue(properties properties.Properties, initialValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDecoder{},
		Operation: operation.Tail{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(""),
		},
	}
	return field
}

// NewDeltaOperation <string/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDeltaDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(""),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <string/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		decode:       decoder.ByteVectorDeltaDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(""),
		},
	}

	return field
}
