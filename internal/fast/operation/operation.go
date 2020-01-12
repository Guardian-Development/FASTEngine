package operation

import (
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap, required bool) bool
	GetNotEncodedValue() (interface{}, error)
	Apply(value interface{}) (interface{}, error)
}

type None struct {
}

func (operation None) ShouldReadValue(pMap *presencemap.PresenceMap, required bool) bool {
	return true
}

func (operation None) GetNotEncodedValue() (interface{}, error) {
	return "", nil
}

func (operation None) Apply(value interface{}) (interface{}, error) {
	return value, nil
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

func (operation Constant) GetNotEncodedValue() (interface{}, error) {
	return operation.ConstantValue, nil
}

func (operation Constant) Apply(value interface{}) (interface{}, error) {
	return value, nil
}
