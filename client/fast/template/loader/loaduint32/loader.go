package loaduint32

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/client/fast/template/structure"
	"github.com/Guardian-Development/fastengine/internal/converter"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/xml"
)

// Load an <uint32 /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fielduint32.FieldUInt32, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fielduint32.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fielduint32.NewDefaultOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToUInt32(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint32.FieldUInt32{}, err
		}

		return fielduint32.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fielduint32.FieldUInt32{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue, err := converter.ToUInt32(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint32.FieldUInt32{}, err
		}

		return fielduint32.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fielduint32.NewCopyOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToUInt32(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fielduint32.FieldUInt32{}, err
		}

		return fielduint32.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fielduint32.FieldUInt32{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}