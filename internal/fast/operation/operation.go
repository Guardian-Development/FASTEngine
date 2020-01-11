package operation

import (
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap) bool
	GetNotEncodedValue() (interface{}, error)
	Apply(value interface{}) (interface{}, error)
}

type None struct {
}

func (operation None) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
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

func (operation Constant) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return false
}

func (operation Constant) GetNotEncodedValue() (interface{}, error) {
	return operation.ConstantValue, nil
}

func (operation Constant) Apply(value interface{}) (interface{}, error) {
	return value, nil
}
