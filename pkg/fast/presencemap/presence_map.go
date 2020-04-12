package presencemap

import (
	"bytes"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/decoder"
)

// PresenceMap represents the pMap in the FAST spec, allowing a message to specify whether a field is present or requires the use of an operation specified
// within the FAST template instead
type PresenceMap struct {
	pMap         []byte
	currentIndex int
}

// GetIsSetAndIncrement returns the next value in the pMap, incrementing the internal counter
func (pMap *PresenceMap) GetIsSetAndIncrement() bool {
	offset := pMap.currentIndex / 7
	if offset >= len(pMap.pMap) {
		return false
	}

	bit := 64 >> byte(pMap.currentIndex-(offset*7))
	isSet := (pMap.pMap[offset] & byte(bit)) > 0
	pMap.currentIndex = pMap.currentIndex + 1

	return isSet
}

// New pMap is created reading the next FAST encoded value off the message buffer to represent the pMap
func New(message *bytes.Buffer) (PresenceMap, error) {
	value, err := decoder.ReadValue(message)

	if err != nil {
		return PresenceMap{}, fmt.Errorf("unable to read a valid value from byte buffer for presence map, reason: %s", err)
	}

	return PresenceMap{pMap: value, currentIndex: 0}, nil
}
