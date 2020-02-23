package loaduint64

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/client/fast/template/structure"
	"github.com/Guardian-Development/fastengine/internal/converter"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielduint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/xml"
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
	default:
		return fielduint64.FieldUInt64{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}