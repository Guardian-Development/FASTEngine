package fielddecimal

import (
	"bytes"
	"fmt"
	"math"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

// FieldDecimal represents a FAST template <decimal/> type
type FieldDecimal struct {
	FieldDetails  properties.Properties
	ExponentField fieldint32.FieldInt32
	MantissaField fieldint64.FieldInt64
}

// Deserialise a <decimal/> from the input source
func (field FieldDecimal) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dict *dictionary.Dictionary) (fix.Value, error) {
	exponentValue, err := field.ExponentField.Deserialise(inputSource, pMap, dict)
	if err != nil {
		return nil, err
	}
	switch exponentValue.(type) {
	case fix.NullValue:
		return fix.NullValue{}, nil
	case fix.RawValue:
		mantissaValue, err := field.MantissaField.Deserialise(inputSource, pMap, dict)
		if err != nil {
			return nil, fmt.Errorf("unable to decode mantissa after successful decoding of exponent")
		}
		decimalValue := math.Pow(10, float64(exponentValue.Get().(int32))) * float64(mantissaValue.Get().(int64))
		fixValue := fix.NewRawValue(decimalValue)
		dict.SetValue(field.FieldDetails.Name, fixValue)
		return fixValue, nil
	}

	return nil, fmt.Errorf("Exponent value of decimal was not expected type: %v", exponentValue)
}

// GetTagId for this field
func (field FieldDecimal) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether either the exponent or mantissa require a pmap bit being set
func (field FieldDecimal) RequiresPmap() bool {
	return field.ExponentField.RequiresPmap() || field.MantissaField.RequiresPmap()
}

// New <decimal/> field with the given properties, exponent and mantissa
func New(properties properties.Properties, exponent fieldint32.FieldInt32, mantissa fieldint64.FieldInt64) store.Unit {
	field := FieldDecimal{
		FieldDetails:  properties,
		ExponentField: exponent,
		MantissaField: mantissa,
	}

	return field
}
