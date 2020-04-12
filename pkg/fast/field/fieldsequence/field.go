package fieldsequence

import (
	"bytes"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"

	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// FieldSequence represents a FAST template <sequence /> type
type FieldSequence struct {
	FieldDetails   properties.Properties
	LengthField    fielduint32.FieldUInt32
	SequenceFields []store.Unit
}

// Deserialise an <sequence/> from the input source
func (field FieldSequence) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, previousValues *dictionary.Dictionary) (fix.Value, error) {
	numberOfElements, err := field.LengthField.Deserialise(inputSource, pMap, previousValues)
	if err != nil {
		field.FieldDetails.Logger.Printf("[FieldSequence][%#v] failed to decode number of elements in sequence from byte buffer, reason: %s", field.FieldDetails, err)
		return nil, fmt.Errorf("[FieldSequence][%#v] failed to decode number of elements in sequence from byte buffer, reason: %s", field.FieldDetails, err)
	}

	switch t := numberOfElements.(type) {
	case fix.NullValue:
		return t, nil
	}

	numberOfRepeatingGroups := numberOfElements.Get().(uint32)
	sequenceValue := fix.NewSequenceValue(numberOfRepeatingGroups)

	for repeatingGroup := uint32(0); repeatingGroup < numberOfRepeatingGroups; repeatingGroup++ {
		sequencePmap := presencemap.PresenceMap{}
		if field.subFieldsRequirePmap() {
			sequencePmap, err = presencemap.New(inputSource)
			if err != nil {
				field.FieldDetails.Logger.Printf("[FieldSequence][%#v] failed to decode pmap for repeating group [%d] in sequence, reason: %s", field.FieldDetails, repeatingGroup, err)
				field.FieldDetails.Logger.Printf("[FieldSequence][%#v] sequence currently decoded before failure %d=%s", field.FieldDetails, field.FieldDetails.ID, sequenceValue.String())
				return nil, fmt.Errorf("[FieldSequence][%#v] failed to decode pmap for repeating group [%d] in sequence, reason: %s", field.FieldDetails, repeatingGroup, err)
			}
		}

		for _, element := range field.SequenceFields {
			value, err := element.Deserialise(inputSource, &sequencePmap, previousValues)
			if err != nil {
				field.FieldDetails.Logger.Printf("[FieldSequence][%#v] failed to decode element for repeating group [%d] in sequence, reason: %s", field.FieldDetails, repeatingGroup, err)
				field.FieldDetails.Logger.Printf("[FieldSequence][%#v] sequence currently decoded before failure %d=%s", field.FieldDetails, field.FieldDetails.ID, sequenceValue.String())
				return nil, fmt.Errorf("[FieldSequence][%#v] failed to decode element for repeating group [%d] in sequence, reason: %s", field.FieldDetails, repeatingGroup, err)
			}

			sequenceValue.SetValue(repeatingGroup, element.GetTagId(), value)
		}
	}

	return sequenceValue, nil
}

func (field FieldSequence) subFieldsRequirePmap() bool {
	for _, element := range field.SequenceFields {
		if element.RequiresPmap() {
			return true
		}
	}

	return false
}

// GetTagId for this field
func (field FieldSequence) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the length element for this sequence requires a pmap
func (field FieldSequence) RequiresPmap() bool {
	return field.LengthField.RequiresPmap()
}

// New <sequence/> field with the given properties, legnth field, and sub fields
func New(properties properties.Properties, length fielduint32.FieldUInt32, sequenceFields []store.Unit) FieldSequence {
	field := FieldSequence{
		FieldDetails:   properties,
		LengthField:    length,
		SequenceFields: sequenceFields,
	}

	return field
}
