package loadint32

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/client/fast/template/structure"
	"github.com/Guardian-Development/fastengine/internal/converter"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/xml"
)

type Int32Converter func(string) (int32, error)

// Load an <int32 /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fieldint32.FieldInt32, error) {
	return LoadWithConverter(tagInTemplate, fieldDetails, converter.ToInt32)
}

// Load an <int32 /> tag with supported operation
func LoadWithConverter(tagInTemplate *xml.Tag, fieldDetails properties.Properties, valueConverter Int32Converter) (fieldint32.FieldInt32, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fieldint32.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fieldint32.NewDefaultOperation(fieldDetails), nil
		}

		operationValue, err := valueConverter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, err
		}

		return fieldint32.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fieldint32.FieldInt32{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue, err := valueConverter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, err
		}

		return fieldint32.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fieldint32.NewCopyOperation(fieldDetails), nil
		}

		operationValue, err := valueConverter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, err
		}

		return fieldint32.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fieldint32.FieldInt32{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}