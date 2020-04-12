package fieldsequence

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
)

//<sequence>
//	<length>
//		<constant value="1" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceWithConstantLength(t *testing.T) {
	// Arrange
	// 1: int64 = 10000001	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{129, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := "1|2=1|3=TEST1|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", true, testLog), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence presence="optional">
//	<length>
//		<constant value="1" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceWithEncodedConstantLength(t *testing.T) {
	// Arrange pmap = 11000000
	// 1: int64 = 10000001	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	messageAsBytes := bytes.NewBuffer([]byte{129, 84, 69, 83, 84, 177})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := "1|2=1|3=TEST1|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", false, testLog), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence presence="optional">
//	<length>
//		<constant value="1" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceWithNotEncodedConstantLength(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", false, testLog), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<sequence>
//	<length>
//		<constant value="1" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsFalseForRequiredSequenceWithConstantLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", true, testLog), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<sequence presence="optional">
//	<length>
//		<constant value="3" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsTrueForOptionalSequenceWithConstantLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", false, testLog), 3),
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
