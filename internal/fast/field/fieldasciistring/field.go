package fieldasciistring

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// FieldAsciiString represents a FAST template <string charset="ascii"/> type and <string /> type
type FieldAsciiString struct {
	FieldDetails properties.Properties
	Operation    operation.Operation
}

// Deserialise a <string/> from the input source
func (field FieldAsciiString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
	previousValue := dictionary.GetValue(field.FieldDetails.Name)
	if field.Operation.ShouldReadValue(pMap) {
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
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <string/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
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
		Operation: operation.Tail{
			InitialValue: fix.NullValue{},
			BaseValue: fix.NewRawValue(""),
		},
	}
	return field
}

// NewTailOperationWithInitialValue <string/> field with the given properties and <tail value="initialValue"/> operator
func NewTailOperationWithInitialValue(properties properties.Properties, initialValue string) FieldAsciiString {
	field := FieldAsciiString{
		FieldDetails: properties,
		Operation: operation.Tail{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue: fix.NewRawValue(""),
		},
	}
	return field
}
