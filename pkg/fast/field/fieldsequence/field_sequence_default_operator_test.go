package fieldsequence

import (
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
)

//<sequence>
//	<length>
//		<default value="3" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsTrueForRequiredSequenceWithDefaultLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.NewDefaultOperationWithValue(properties.New(1, "SequenceField", true, testLog), 3),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<sequence presence="optional">
//	<length>
//		<default value="3" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsTrueForOptionalSequenceWithDefaultLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.NewDefaultOperationWithValue(properties.New(1, "SequenceField", false, testLog), 3),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
