package loadint64

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/client/fast/template/structure"
	"github.com/Guardian-Development/fastengine/internal/converter"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/xml"
)

type Int64Converter func(string) (int64, error)

// Load an <int64 /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fieldint64.FieldInt64, error) {
	return LoadWithConverter(tagInTemplate, fieldDetails, converter.ToInt64)
}

// Load an <int64 /> tag with supported operation
func LoadWithConverter(tagInTemplate *xml.Tag, fieldDetails properties.Properties, int64Converter Int64Converter) (fieldint64.FieldInt64, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fieldint64.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fieldint64.NewDefaultOperation(fieldDetails), nil
		}

		operationValue, err := int64Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint64.FieldInt64{}, err
		}

		return fieldint64.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fieldint64.FieldInt64{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue, err := int64Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint64.FieldInt64{}, err
		}

		return fieldint64.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fieldint64.NewCopyOperation(fieldDetails), nil
		}

		operationValue, err := int64Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint64.FieldInt64{}, err
		}

		return fieldint64.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.IncrementOperation:
		if !hasOperationValue {
			return fieldint64.NewIncrementOperation(fieldDetails), nil
		}

		operationValue, err := int64Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint64.FieldInt64{}, err
		}

		return fieldint64.NewIncrementOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fieldint64.FieldInt64{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}
