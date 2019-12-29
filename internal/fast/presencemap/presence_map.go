package presencemap

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/internal/fast"
)

type PresenceMap struct {
	pMap uint32
}

func New(message *bytes.Buffer) (PresenceMap, error) {
	pMap, err := fast.ReadUInt32(message)

	if err != nil {
		return PresenceMap{}, err
	}

	return PresenceMap{pMap: pMap}, nil
}
