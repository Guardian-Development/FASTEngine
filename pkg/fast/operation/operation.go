package operation

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/errors"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fast/value"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// Operation is applied when reading a fast value off a byte buffer
type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap) bool
	GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error)
	Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error)
	RequiresPmap(required bool) bool
}

// None represents applying no operation to the read value
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
func (operation None) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns false as no operator is present
func (operation None) RequiresPmap(required bool) bool {
	return false
}

// Constant represents the fast <constant/> operation
type Constant struct {
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
func (operation Constant) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// Default represents the fast <default/> operation
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
func (operation Default) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns true, as the Default operator always evaluates the next pMap bit.
func (operation Default) RequiresPmap(required bool) bool {
	return true
}

// Copy represents the fast <copy/> operation
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
				return nil, fmt.Errorf("%s", errors.D5)
			}
		}

		return operation.InitialValue, nil
	}

	return nil, fmt.Errorf("unsupported previous dictionary value for operation, value: %s", previousValue)
}

// Apply does not modify the value, as the Copy operator only applies to retrieving a value from the stream, not mutating it
func (operation Copy) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns true, as the Copy operator always evaluates the next pMap bit.
func (operation Copy) RequiresPmap(required bool) bool {
	return true
}

// Increment represents the fast <increment/> operation
type Increment struct {
	InitialValue fix.Value
}

// ShouldReadValue returns the result of reading the pMap. Increment operation always evaluates the next pMap bit as it needs to know whether to read the value or
// increment the previous value in the dictionary
func (operation Increment) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return pMap.GetIsSetAndIncrement()
}

// GetNotEncodedValue returns the previous value incremented by 1. If the previous value is undefined and its a required field, and there is no initial value, an error is returned. If the previous value
// is undefined and the field is not required, then the initial value is returned, which may be null.
func (operation Increment) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
	switch t := previousValue.(type) {
	case dictionary.AssignedValue:
		switch q := t.Value.Get().(type) {
		case uint32:
			return fix.NewRawValue(q + 1), nil
		case uint64:
			return fix.NewRawValue(q + 1), nil
		case int32:
			return fix.NewRawValue(q + 1), nil
		case int64:
			return fix.NewRawValue(q + 1), nil
		default:
			return nil, fmt.Errorf("unsupported type for increment operator, can only increment integers")
		}
	case dictionary.EmptyValue:
		if required {
			return nil, fmt.Errorf("%s", errors.D6)
		}
		return fix.NullValue{}, nil
	case dictionary.UndefinedValue:
		switch operation.InitialValue.(type) {
		case fix.NullValue:
			if required {
				return nil, fmt.Errorf("%s", errors.D5)
			}
		}

		return operation.InitialValue, nil
	}

	return nil, fmt.Errorf("unsupported previous dictionary value for operation, value: %s", previousValue)
}

// Apply does not modify the value, as the Increment operator only applies when there is no value in the stream
func (operation Increment) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	return readValue.GetAsFix(), nil
}

// RequiresPmap always returns true, as the Increment operator always evaluates the next pMap bit.
func (operation Increment) RequiresPmap(required bool) bool {
	return true
}

// Tail represents the fast <tail/> operation
type Tail struct {
	InitialValue fix.Value
	BaseValue    fix.Value
}

// ShouldReadValue returns the result of reading the pMap. Tail operation always evaluates the next pMap bit as it needs to know whether to read the value and
// combine it with the previous value, or to just use the previous value
func (operation Tail) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return pMap.GetIsSetAndIncrement()
}

// GetNotEncodedValue returns the previous value. If the previous value is undefined and its a required field, and there is no initial value, an error is returned. If the previous value
// is undefined and the field is not required, then the initial value is returned, which may be null.
func (operation Tail) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
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

// Apply takes the previous value and combines it with the read value. If the read value is larger than the previous value, the read value overwrites the
// previous value
func (operation Tail) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	// if null encoded, field is considered absent
	switch readValue.(type) {
	case value.NullValue:
		return readValue.GetAsFix(), nil
	}

	var baseValue fix.Value

	// work out base value based on previous value
	switch t := previousValue.(type) {
	case dictionary.AssignedValue:
		baseValue = t.Value
	case dictionary.EmptyValue, dictionary.UndefinedValue:
		switch operation.InitialValue.(type) {
		case fix.NullValue:
			baseValue = operation.BaseValue
		default:
			baseValue = operation.InitialValue
		}
	}

	// combine base value with read value
	switch t := readValue.(type) {
	case value.StringValue:
		return t.ApplyTail(baseValue)
	case value.ByteVector:
		return t.ApplyTail(baseValue)
	}

	return nil, fmt.Errorf("unsupported type for tail operator, you can only use this with strings and byte vectors")
}

// RequiresPmap always returns true, as the Tail operator always evaluates the next pMap bit
func (operation Tail) RequiresPmap(required bool) bool {
	return true
}

// Delta represents the fast <delta/> operation.
type Delta struct {
	InitialValue fix.Value
	BaseValue    fix.Value
}

// ShouldReadValue is always true as delta just informs you of what to do to build the read value, not what to do if the value is not present.
func (operation Delta) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return true
}

// GetNotEncodedValue returns nil, however value is always encoded
func (operation Delta) GetNotEncodedValue(pMap *presencemap.PresenceMap, required bool, previousValue dictionary.Value) (fix.Value, error) {
	return fix.NullValue{}, nil
}

// Apply returns the result of previous value + read value (delta). If theres no previous value, the initial value (or default value) is used.
// if the previous value is empty an error is returned (you cannot apply a delta to a null value)
func (operation Delta) Apply(readValue value.Value, previousValue dictionary.Value) (fix.Value, error) {
	switch readValue.(type) {
	case value.NullValue:
		return readValue.GetAsFix(), nil
	}

	var baseValue fix.Value

	// work out base value based on previous value
	switch t := previousValue.(type) {
	case dictionary.AssignedValue:
		baseValue = t.Value
	case dictionary.EmptyValue, dictionary.UndefinedValue:
		switch operation.InitialValue.(type) {
		case fix.NullValue:
			baseValue = operation.BaseValue
		default:
			baseValue = operation.InitialValue
		}
	}

	return readValue.Add(baseValue)
}

// RequiresPmap always returns false as value is always read
func (operation Delta) RequiresPmap(required bool) bool {
	return false
}
