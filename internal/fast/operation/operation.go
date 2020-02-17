package operation

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap) bool
	GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error)
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
func (operation None) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
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
func (operation Constant) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
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
func (operation Default) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
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

type Copy struct {
	InitialValue interface{}
}

// ShouldReadValue returns the result of reading the pMap. Copy operation always evaluates the next pMap bit as it needs to know whether to read the value or
// copy the previous value in the dictionary
func (operation Copy) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return pMap.GetIsSetAndIncrement()
}

// GetNotEncodedValue returns the previous value. If the previous value is undefined and its a required field, and there is no initial value, an error is returned. If the previous value
// is undefined and the field is not required, then the initial value is returned, which may be null.
func (operation Copy) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
	switch t := previousValue.(type) {
	case dictionary.AssignedValue:
		return t.Value, nil
	case dictionary.EmptyValue:
		return nil, nil
	case dictionary.UndefinedValue:
		if operation.InitialValue == nil && required {
			return nil, fmt.Errorf("no value supplied in message and no initial value with required field")
		}
		return fix.NewRawValue(operation.InitialValue), nil
	}

	return nil, fmt.Errorf("unsupported previous dictionary value for operation, value: %s", previousValue)
}

// Apply does to modify the value, as the Copy operator only applies to retrieving a value from the stream, not mutating it
func (operation Copy) Apply(readValue value.Value) (fix.Value, error) {
	return convertToFixNoTransformation(readValue)
}

// RequiresPmap always returns true, as the Copy operator always evaluates the next pMap bit.
func (operation Copy) RequiresPmap(required bool) bool {
	return true
}

func convertToFixNoTransformation(readValue value.Value) (fix.Value, error) {
	// TODO: just have a fast value have a get method on it with interface{} return type
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
