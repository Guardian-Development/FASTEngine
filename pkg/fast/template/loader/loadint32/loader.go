package loadint32

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/loader/converter"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/structure"
)

type Int32Converter func(string) (int32, error)

// Load an <int32 /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fieldint32.FieldInt32, error) {
	return LoadWithConverter(tagInTemplate, fieldDetails, converter.ToInt32)
}

// Load an <int32 /> tag with supported operation
func LoadWithConverter(tagInTemplate *xml.Tag, fieldDetails properties.Properties, int32Converter Int32Converter) (fieldint32.FieldInt32, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fieldint32.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue && fieldDetails.Required {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s", errors.S5)
		}

		if !hasOperationValue {
			return fieldint32.NewDefaultOperation(fieldDetails), nil
		}

		operationValue, err := int32Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s: %s", errors.S3, err)
		}

		return fieldint32.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s", errors.S4)
		}

		operationValue, err := int32Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s: %s", errors.S3, err)
		}

		return fieldint32.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fieldint32.NewCopyOperation(fieldDetails), nil
		}

		operationValue, err := int32Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s: %s", errors.S3, err)
		}

		return fieldint32.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.IncrementOperation:
		if !hasOperationValue {
			return fieldint32.NewIncrementOperation(fieldDetails), nil
		}

		operationValue, err := int32Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s: %s", errors.S3, err)
		}

		return fieldint32.NewIncrementOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.DeltaOperation:
		if !hasOperationValue {
			return fieldint32.NewDeltaOperation(fieldDetails), nil
		}

		operationValue, err := int32Converter(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldint32.FieldInt32{}, fmt.Errorf("%s: %s", errors.S3, err)
		}

		return fieldint32.NewDeltaOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fieldint32.FieldInt32{}, fmt.Errorf("%s: %s", errors.S2, operationTag)
	}
}
