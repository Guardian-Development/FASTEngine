package presencemap

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/decoder"
)

// PresenceMap represents the pMap in the FAST spec, allowing a message to specify whether a field is present or requires the use of an operation specified
// within the FAST template instead
type PresenceMap struct {
	pMap         []byte
	currentIndex uint32
	currentByte  uint32
}

// GetIsSetAndIncrement returns the next value in the pMap, incrementing the internal counter
func (pMap *PresenceMap) GetIsSetAndIncrement() bool {
	byteInPmap := pMap.pMap[pMap.currentByte]
	valueInByte := (1 << pMap.currentIndex) & byteInPmap
	isSet := valueInByte == (1 << pMap.currentIndex)

	if pMap.currentIndex == 1 {
		pMap.currentByte = pMap.currentByte + 1
		pMap.currentIndex = 7
	} else {
		pMap.currentIndex = pMap.currentIndex - 1
	}

	return isSet
}

// New pMap is created reading the next FAST encoded value off the message buffer to represent the pMap
func New(message *bytes.Buffer) (PresenceMap, error) {
	value, err := decoder.ReadValue(message)

	if err != nil {
		return PresenceMap{}, err
	}

	return PresenceMap{pMap: value, currentIndex: 7}, nil
}
