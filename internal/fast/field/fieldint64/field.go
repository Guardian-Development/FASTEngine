package fieldint64

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

// FieldInt64 represents a FAST template <int64/> type
type FieldInt64 struct {
	FieldDetails properties.Properties
	Operation    operation.Operation
}

// Deserialise an <int64/> from the input source
func (field FieldInt64) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
	if field.Operation.ShouldReadValue(pMap) {
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

		transformedValue, err := field.Operation.Apply(value)
		if err != nil {
			return nil, err
		}
		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	previousValue := dictionary.GetValue(field.FieldDetails.Name)
	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		return nil, err
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
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <int64/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		Operation: operation.Constant{
			ConstantValue: constantValue,
		},
	}

	return field
}

// NewDefaultOperation <int64/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		Operation: operation.Default{
			DefaultValue: nil,
		},
	}

	return field
}

// NewDefaultOperationWithValue <int64/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		Operation: operation.Default{
			DefaultValue: defaultValue,
		},
	}

	return field
}

// NewCopyOperation <int64/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		Operation: operation.Copy{
			InitialValue: nil,
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <int64/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue int64) FieldInt64 {
	field := FieldInt64{
		FieldDetails: properties,
		Operation: operation.Copy{
			InitialValue: initialValue,
		},
	}

	return field
}
