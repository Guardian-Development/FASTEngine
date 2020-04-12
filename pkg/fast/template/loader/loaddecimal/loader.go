package loaddecimal

import (
	"fmt"

	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielddecimal"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/loader/converter"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/loader/loadint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/loader/loadint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/template/structure"
)

// Load a <decimal /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fielddecimal.FieldDecimal, error) {
	if len(tagInTemplate.NestedTags) < 2 {
		exponentField, err := loadint32.LoadWithConverter(tagInTemplate, fieldDetails, converter.ToExponent)
		if err != nil {
			return fielddecimal.FieldDecimal{}, fmt.Errorf("[%s][%v] failed to load exponent, reason: %s", tagInTemplate.Type, fieldDetails, err)
		}
		exponentField.FieldDetails.Name = fmt.Sprintf("%sExponent", fieldDetails.Name)
		mantissaField, err := loadint64.LoadWithConverter(tagInTemplate, fieldDetails, converter.ToMantissa)
		if err != nil {
			return fielddecimal.FieldDecimal{}, fmt.Errorf("[%s][%v] failed to load mantissa, reason: %s", tagInTemplate.Type, fieldDetails, err)
		}

		mantissaField.FieldDetails.Required = true
		mantissaField.FieldDetails.Name = fmt.Sprintf("%sMantissa", fieldDetails.Name)

		return fielddecimal.New(fieldDetails, exponentField, mantissaField), nil
	}
	if len(tagInTemplate.NestedTags) == 2 {
		exponentTag := tagInTemplate.NestedTags[0]
		exponentField, err := loadint32.Load(&exponentTag, fieldDetails)
		if err != nil {
			return fielddecimal.FieldDecimal{}, fmt.Errorf("[%s][%v] failed to load exponent, reason: %s", tagInTemplate.Type, fieldDetails, err)
		}

		exponentName := exponentTag.Attributes["name"]
		if structure.IsNullString(exponentName) {
			exponentName = fmt.Sprintf("%sExponent", fieldDetails.Name)
		}
		exponentField.FieldDetails.Name = exponentName

		mantissaTag := tagInTemplate.NestedTags[1]
		mantissaField, err := loadint64.Load(&mantissaTag, fieldDetails)
		if err != nil {
			return fielddecimal.FieldDecimal{}, fmt.Errorf("[%s][%v] failed to load mantissa, reason: %s", tagInTemplate.Type, fieldDetails, err)
		}
		mantissaField.FieldDetails.Required = true
		mantissaName := mantissaTag.Attributes["name"]
		if structure.IsNullString(mantissaName) {
			mantissaName = fmt.Sprintf("%sMantissa", fieldDetails.Name)
		}
		mantissaField.FieldDetails.Name = mantissaName

		return fielddecimal.New(fieldDetails, exponentField, mantissaField), nil
	}

	return fielddecimal.FieldDecimal{}, fmt.Errorf("decimal must be declared with either no operation (empty), or with <exponent/> and <mantissa/>")
}
