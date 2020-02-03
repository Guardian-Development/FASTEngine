package operation

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

type Operation interface {
	ShouldReadValue() bool
	GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool) (fix.Value, error)
	Apply(readValue value.Value) (fix.Value, error)
	RequiresPmap(required bool) bool
}

type None struct {
}

func (operation None) ShouldReadValue() bool {
	return true
}

func (operation None) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool) (fix.Value, error) {
	return fix.NullValue{}, nil
}

func (operation None) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

func (operation None) RequiresPmap(required bool) bool {
	return false
}

type Constant struct {
	ConstantValue interface{}
}

// ShouldReadValue always returns false for constant operations
func (operation Constant) ShouldReadValue() bool {
	return false
}

// RequiresPmap will return true is the value is marked as optional.
func (operation Constant) RequiresPmap(required bool) bool {
	return !required
}

// GetNotEncodedValue returns default value if required field. If optional and pmap bit set to 1, returns default value, else returns null
func (operation Constant) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool) (fix.Value, error) {
	if required {
		return fix.NewRawValue(operation.ConstantValue), nil
	}
	if pMap.GetIsSetAndIncrement() {
		return fix.NewRawValue(operation.ConstantValue), nil
	}

	return fix.NullValue{}, nil
}

func (operation Constant) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

func convertToFixNoTransformation(readValue value.Value) (fix.Value, error) {
	// TODO: use visitor pattern instead
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
	case value.ByteVector:
		return fix.NewRawValue(t.Value), nil
	}

	return nil, fmt.Errorf("Unsupported fast value for operation, value: %s", readValue)
}
