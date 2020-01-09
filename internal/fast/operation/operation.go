package operation

import (
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

type Operation interface {
	ShouldReadValue(pMap *presencemap.PresenceMap) bool
	GetNotEncodedValue() (interface{}, error)
}

type None struct {
}

func (operation None) ShouldReadValue(pMap *presencemap.PresenceMap) bool {
	return true
}

func (operation None) GetNotEncodedValue() (interface{}, error) {
	return "", nil
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
