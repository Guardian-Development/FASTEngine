package fieldsequence

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
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
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(1)),
					3: fix.NewRawValue("TEST1"),
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", true), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result)
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
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
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(1)),
					3: fix.NewRawValue("TEST1"),
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", false),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", false), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dictionary)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result)
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
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
		properties.New(1, "SequenceField", false),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", false), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
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
		properties.New(1, "SequenceField", true),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", true), 1),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
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
		properties.New(1, "SequenceField", false),
		fielduint32.NewConstantOperation(properties.New(1, "SequenceField", false), 3),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}