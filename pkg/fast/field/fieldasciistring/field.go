package fieldasciistring

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

// FieldAsciiString represents a FAST template <string charset="ascii"/> type and <string /> type
type FieldAsciiString struct {
	FieldDetails properties.Properties
	Operation    operation.Operation

	decode decoder.Decoder
}

// Deserialise a <string/> from the input source
func (field FieldAsciiString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
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
			field.FieldDetails.Logger.Printf("[FieldAsciiString][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
			return nil, fmt.Errorf("[FieldAsciiString][%#v][%#v] failed to decode value from byte buffer, reason: %s", field.FieldDetails, field.Operation, err)
		}

		transformedValue, err := field.Operation.Apply(readValue, previousValue)
		if err != nil {
			field.FieldDetails.Logger.Printf("[FieldAsciiString][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, readValue, previousValue, err)
			return nil, fmt.Errorf("[FieldAsciiString][%#v][%#v] failed to apply operation with readValue %#v, previousValue: %#v, reason: %s", field.FieldDetails, field.Operation, readValue, previousValue, err)
		}

		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		field.FieldDetails.Logger.Printf("[FieldAsciiString][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
		return nil, fmt.Errorf("[FieldAsciiString][%#v][%#v] failed to get value for field when not encoded in message, reason: %s", field.FieldDetails, field.Operation, err)
	}

	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldAsciiString) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldAsciiString) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <string/> field with the given properties and no operation
func New(properties properties.Properties) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <string/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <string/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <string/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <string/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <string/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewTailOperation <string/> field with the given properties and <tail/> operator
func NewTailOperation(properties properties.Properties) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Tail{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(""),
		},
	}
	return field
}

// NewTailOperationWithInitialValue <string/> field with the given properties and <tail value="initialValue"/> operator
func NewTailOperationWithInitialValue(properties properties.Properties, initialValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDecoder{},
		Operation: operation.Tail{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(""),
		},
	}
	return field
}

// NewDeltaOperation <string/> field with the given properties and <delta/> operator
func NewDeltaOperation(properties properties.Properties) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDeltaDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NullValue{},
			BaseValue:    fix.NewRawValue(""),
		},
	}

	return field
}

// NewDeltaOperationWithInitialValue <string/> field with the given properties and <delta value="initialValue"/> operator
func NewDeltaOperationWithInitialValue(properties properties.Properties, initialValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		decode:       decoder.AsciiStringDeltaDecoder{},
		Operation: operation.Delta{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue:    fix.NewRawValue(""),
		},
	}

	return field
}
