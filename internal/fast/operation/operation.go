package operation

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap) bool
	GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool) (fix.Value, error)
	Apply(readValue value.Value) (fix.Value, error)
	RequiresPmap(required bool) bool
}

type None struct {
}

// ShouldReadValue if no operator is present must always read the value from the stream
func (operation None) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return true
}

// GetNotEncodedValue if the value is not encoded, and there is no operator, the value is always nil
func (operation None) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool) (fix.Value, error) {
	return fix.NullValue{}, nil
}

// Apply does no transformation on the value as no operator is present
func (operation None) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

// RequiresPmap always returns false as no operator is present
func (operation None) RequiresPmap(required bool) bool {
	return false
}

type Constant struct {
	ConstantValue interface{}
}

// ShouldReadValue always returns false for constant operations
func (operation Constant) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
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

// Apply does to modify the value, as the Constant operator only applies to retrieving a value from the stream, not mutating it
func (operation Constant) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

type Default struct {
	DefaultValue interface{}
}

// ShouldReadValue returns the result of reading the pMap. Default operation always evaluates the next pMap bit.
func (operation Default) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return pMap.GetIsSetAndIncrement()
}

// GetNotEncodedValue returns the configured DefaultValue. If this is null, it returns fix.NullValue{}.
func (operation Default) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool) (fix.Value, error) {
	return fix.NewRawValue(operation.DefaultValue), nil
}

// Apply does to modify the value, as the Default operator only applies to retrieving a value from the stream, not mutating it
func (operation Default) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

// RequiresPmap always returns true, as the Default operator always evaluates the next pMap bit.
func (operation Default) RequiresPmap(required bool) bool {
	return true
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
