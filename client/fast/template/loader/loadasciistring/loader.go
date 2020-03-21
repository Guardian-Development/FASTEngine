package loadasciistring

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/client/fast/template/structure"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/xml"
)

// Load an <string /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fieldasciistring.FieldAsciiString, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fieldasciistring.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fieldasciistring.NewDefaultOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldasciistring.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fieldasciistring.FieldAsciiString{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldasciistring.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fieldasciistring.NewCopyOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldasciistring.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.TailOperation:
		if !hasOperationValue {
			return fieldasciistring.NewTailOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldasciistring.NewTailOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fieldasciistring.FieldAsciiString{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}
