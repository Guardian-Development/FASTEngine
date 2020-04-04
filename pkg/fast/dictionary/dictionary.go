package dictionary

import "github.com/Guardian-Development/fastengine/pkg/fix"

// Value represents a value associated with a key
type Value interface {
}

// EmptyValue means the previous value has been evaluated and is nil
type EmptyValue struct {
}

// UndefinedValue means the previous value has not been set
type UndefinedValue struct {
}

// AssignedValue means the previous value has been evaluated and is not nil
type AssignedValue struct {
	Value fix.Value
}

// Dictionary represents a key value store of values
type Dictionary struct {
	keys map[string]Value
}

// SetValue sets the associated value with the key
func (dictionary *Dictionary) SetValue(key string, value fix.Value) {
	switch t := value.(type) {
	case fix.NullValue:
		dictionary.keys[key] = EmptyValue{}
	case fix.RawValue:
		dictionary.keys[key] = AssignedValue{Value: t}
	}
}

// GetValue gets the associated value with the key. If no value is associated this returns UndefinedValue
func (dictionary Dictionary) GetValue(key string) Value {
	if value, exists := dictionary.keys[key]; exists {
		return value
	}

	return UndefinedValue{}
}

// Reset the internal set of key/value pairs to be empty
func (dictionary *Dictionary) Reset() {
	dictionary.keys = make(map[string]Value)
}

// New dictionary to hold key/value pairs within
func New() Dictionary {
	return Dictionary{
		keys: make(map[string]Value),
	}
}
