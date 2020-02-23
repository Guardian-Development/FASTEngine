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
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns false as no operator is present
func (operation None) RequiresPmap(required bool) bool {
	return false
}

type Constant struct {
	// TODO: cleanup how we build up the elements from XML.
	// TODO: look at the public interface, just have the engine, everything else private
	// TODO: look at test names and make them better
	ConstantValue fix.Value
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
		return operation.ConstantValue, nil
	}
	if pMap.GetIsSetAndIncrement() {
		return operation.ConstantValue, nil
	}

	return fix.NullValue{}, nil
}

// Apply does to modify the value, as the Constant operator only applies to retrieving a value from the stream, not mutating it
func (operation Constant) Apply(readValue value.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

type Default struct {
	DefaultValue fix.Value
}

// ShouldReadValue returns the result of reading the pMap. Default operation always evaluates the next pMap bit.
func (operation Default) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return pMap.GetIsSetAndIncrement()
}

// GetNotEncodedValue returns the configured DefaultValue. If this is null, it returns fix.NullValue{}.
func (operation Default) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
	return operation.DefaultValue, nil
}

// Apply does to modify the value, as the Default operator only applies to retrieving a value from the stream, not mutating it
func (operation Default) Apply(readValue value.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns true, as the Default operator always evaluates the next pMap bit.
func (operation Default) RequiresPmap(required bool) bool {
	return true
}

type Copy struct {
	InitialValue fix.Value
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
		return fix.NullValue{}, nil
	case dictionary.UndefinedValue:
		switch operation.InitialValue.(type) {
		case fix.NullValue:
			if required {
				return nil, fmt.Errorf("no value supplied in message and no initial value with required field")
			}
		}

		return operation.InitialValue, nil
	}

	return nil, fmt.Errorf("unsupported previous dictionary value for operation, value: %s", previousValue)
}

// Apply does to modify the value, as the Copy operator only applies to retrieving a value from the stream, not mutating it
func (operation Copy) Apply(readValue value.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns true, as the Copy operator always evaluates the next pMap bit.
func (operation Copy) RequiresPmap(required bool) bool {
	return true
}
