package operation

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
	"github.com/Guardian-Development/fastengine/internal/fix"
)

type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap, required bool) bool
	GetNotEncodedValue() (fix.Value, error)
	Apply(readValue value.Value) (fix.Value, error)
}

type None struct {
}

func (operation None) ShouldReadValue(pMap *presencemap.PresenceMap, required bool) bool {
	return true
}

func (operation None) GetNotEncodedValue() (fix.Value, error) {
	return nil, nil
}

func (operation None) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

type Constant struct {
	ConstantValue interface{}
}

// ShouldReadValue will return false is the value is not marked as optional. If it is marked as optional, it will return the result of reading the
// next value in the pMap
func (operation Constant) ShouldReadValue(pMap *presencemap.PresenceMap, required bool) bool {
	if !required {
		return pMap.GetIsSetAndIncrement()
	}
	return false
}

func (operation Constant) GetNotEncodedValue() (fix.Value, error) {
	return fix.NewRawValue(operation.ConstantValue), nil
}

func (operation Constant) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

func convertToFixNoTransformation(readValue value.Value) (fix.Value, error) {
	switch t := readValue.(type) {
	case value.NullValue:
		return fix.NewRawValue(nil), nil
	case value.StringValue:
		return fix.NewRawValue(t.Value), nil
	case value.UInt32Value:
		return fix.NewRawValue(t.Value), nil
	case value.Int32Value:
		return fix.NewRawValue(t.Value), nil
	case value.UInt64Value:
		return fix.NewRawValue(t.Value), nil
	case value.Int64Value:
		return fix.NewRawValue(t.Value), nil
	}

	return nil, fmt.Errorf("Unsupported fast value for operation, value: %s", readValue)
}
