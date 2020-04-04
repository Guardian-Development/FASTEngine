package loadunicodestring

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/template/structure"

	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldunicodestring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
)

// Load an <string charset="unicode"/> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fieldunicodestring.FieldUnicodeString, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fieldunicodestring.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fieldunicodestring.NewDefaultOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldunicodestring.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fieldunicodestring.FieldUnicodeString{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldunicodestring.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fieldunicodestring.NewCopyOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldunicodestring.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.TailOperation:
		if !hasOperationValue {
			return fieldunicodestring.NewTailOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldunicodestring.NewTailOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.DeltaOperation:
		if !hasOperationValue {
			return fieldunicodestring.NewDeltaOperation(fieldDetails), nil
		}

		operationValue := operationTag.Attributes[structure.ValueAttribute]
		return fieldunicodestring.NewDeltaOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fieldunicodestring.FieldUnicodeString{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}
