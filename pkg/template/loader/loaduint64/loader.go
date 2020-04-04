package loaduint64

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/template/structure"

	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/converter"
)

// Load an <uint64 /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fielduint64.FieldUInt64, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fielduint64.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fielduint64.NewDefaultOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToUInt64(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint64.FieldUInt64{}, err
		}

		return fielduint64.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fielduint64.FieldUInt64{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue, err := converter.ToUInt64(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint64.FieldUInt64{}, err
		}

		return fielduint64.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fielduint64.NewCopyOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToUInt64(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint64.FieldUInt64{}, err
		}

		return fielduint64.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.IncrementOperation:
		if !hasOperationValue {
			return fielduint64.NewIncrementOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToUInt64(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint64.FieldUInt64{}, err
		}

		return fielduint64.NewIncrementOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.DeltaOperation:
		if !hasOperationValue {
			return fielduint64.NewDeltaOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToUInt64(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint64.FieldUInt64{}, err
		}

		return fielduint64.NewDeltaOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fielduint64.FieldUInt64{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}
