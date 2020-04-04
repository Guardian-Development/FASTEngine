package fielddecimal

import (
	"bytes"
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	fieldint322 "github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	fieldint642 "github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	properties2 "github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"math"

	"github.com/Guardian-Development/fastengine/pkg/fix"
)

// FieldDecimal represents a FAST template <decimal/> type
type FieldDecimal struct {
	FieldDetails  properties2.Properties
	ExponentField fieldint322.FieldInt32
	MantissaField fieldint642.FieldInt64
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
		exponentRawValue := exponentValue.Get().(int32)
		if exponentRawValue < -63 || exponentRawValue > 63 {
			return nil, fmt.Errorf("a decimal exponent must fall in the range of [-63...63]")
		}
		mantissaValue, err := field.MantissaField.Deserialise(inputSource, pMap, dict)
		if err != nil {
			return nil, fmt.Errorf("unable to decode mantissa after successful decoding of exponent: %s", err)
		}
		decimalValue := math.Pow(10, float64(exponentValue.Get().(int32))) * float64(mantissaValue.Get().(int64))
		fixValue := fix.NewRawValue(decimalValue)
		dict.SetValue(field.FieldDetails.Name, fixValue)
		return fixValue, nil
	}

	return nil, fmt.Errorf("exponent value of decimal was not expected type: %v", exponentValue)
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
func New(properties properties2.Properties, exponent fieldint322.FieldInt32, mantissa fieldint642.FieldInt64) FieldDecimal {
	field := FieldDecimal{
		FieldDetails:  properties,
		ExponentField: exponent,
		MantissaField: mantissa,
	}

	return field
}
