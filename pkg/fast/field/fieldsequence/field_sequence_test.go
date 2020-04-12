package fieldsequence

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

var testLog = log.New(os.Stdout, "", log.LstdFlags)

//<sequence id="1">
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceOfLengthZero(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := fix.NewSequenceValue(0)
	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	areEqual := reflect.DeepEqual(expectedMessage, result)
	if !areEqual {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.Get())
	}
}

//<sequence id="1">
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceOfLengthTwo(t *testing.T) {
	// Arrange length(2) = 10000010
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{130, 131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=3|3=TEST1|2=2|3=TEST2|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence presence="optional">
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceOfLengthZero(t *testing.T) {
	// Arrange length = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := fix.NewSequenceValue(0)
	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.New(properties.New(1, "SequenceField", false, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
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
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceNull(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.New(properties.New(1, "SequenceField", false, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<sequence presence="optional">
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceOfLengthTwo(t *testing.T) {
	// Arrange length(2) = 10000011
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{131, 131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=3|3=TEST1|2=2|3=TEST2|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.New(properties.New(1, "SequenceField", false, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence>
//	<length />
// 	<int64 id="2"/>
//	<sequence id="3">
//		<length id="4"/>
// 		<string id="5"/>
//	</sequence>
//</sequence>
func TestCanDeseraliseSequenceWithNestedRequiredSequence(t *testing.T) {
	// Arrange length(2) = 10000010
	// 1: int64 = 10000011	nested - length(2) = 10000010
	// 			1. string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	//			2. string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	// 2: int64 = 10000100	nested - length(1) = 10000001
	//			1. string(TEST3) = 01010100 01000101 01010011 01010100 10110011
	messageAsBytes := bytes.NewBuffer([]byte{130, 131, 130, 84, 69, 83, 84, 177, 84, 69, 83, 84, 178, 132, 129, 84, 69, 83, 84, 179})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=3|3=2|5=TEST1|5=TEST2|2=4|3=1|5=TEST3|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			New(properties.New(3, "SequenceTwoField", true, testLog),
				fielduint32.New(properties.New(4, "SequenceTwoField", true, testLog)),
				[]store.Unit{
					fieldasciistring.New(properties.New(5, "AsciiStringTwoField", true, testLog)),
				}),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence>
//	<length />
// 	<int64 id="2"/>
//	<sequence id="3" presence="optional">
//		<length id="4"/>
// 		<string id="5"/>
//	</sequence>
//</sequence>
func TestCanDeseraliseSequenceWithNestedOptionalSequence(t *testing.T) {
	// Arrange length(2) = 10000010
	// 1: int64 = 10000011	nested - length(1) = 10000010
	// 			1. string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000100	nested - length(nil) = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{130, 131, 130, 84, 69, 83, 84, 177, 132, 128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=3|3=1|5=TEST1|2=4|3=nil|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true, testLog)),
			New(properties.New(3, "SequenceTwoField", false, testLog),
				fielduint32.New(properties.New(4, "SequenceTwoField", false, testLog)),
				[]store.Unit{
					fieldasciistring.New(properties.New(5, "AsciiStringTwoField", true, testLog)),
				}),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence>
//	<length />
// 	<int64 id="2" presence="optional">
//		<constant value="2" />
// 	<int64/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceWithRequiredPmapPerIteration(t *testing.T) {
	// Arrange length(2) = 10000010
	// 1: pmap = 10000000	int64 = nil (see pmap)	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: pmap = 11000000	int64 = 2 (see pmap)	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{130, 128, 84, 69, 83, 84, 177, 192, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=nil|3=TEST1|2=2|3=TEST2|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", false, testLog), 2),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence>
//	<length />
// 	<int64 id="2">
//		<constant value="2" />
// 	<int64/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceWithNoPmapPerIteration(t *testing.T) {
	// Arrange length(2) = 10000010
	// 1: int64 = 2	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 2 string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{130, 84, 69, 83, 84, 177, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=2|3=TEST1|2=2|3=TEST2|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", true, testLog), 2),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence presence="optional">
//	<length />
// 	<int64 id="2" presence="optional">
//		<constant value="2" />
// 	<int64/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceWithRequiredPmapPerIteration(t *testing.T) {
	// Arrange length(2) = 10000011
	// 1: pmap = 10000000	int64 = nil (see pmap)	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: pmap = 11000000	int64 = 2 (see pmap)	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{131, 128, 84, 69, 83, 84, 177, 192, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=nil|3=TEST1|2=2|3=TEST2|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.New(properties.New(1, "SequenceField", false, testLog)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", false, testLog), 2),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence presence="optional">
//	<length />
// 	<int64 id="2">
//		<constant value="2" />
// 	<int64/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceWithNoPmapPerIteration(t *testing.T) {
	// Arrange length(2) = 10000011
	// 1: int64 = 2	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 2 string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{131, 84, 69, 83, 84, 177, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := "2|2=2|3=TEST1|2=2|3=TEST2|"

	unitUnderTest := New(
		properties.New(1, "SequenceField", false, testLog),
		fielduint32.New(properties.New(1, "SequenceField", false, testLog)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", true, testLog), 2),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true, testLog)),
		})

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if expectedMessage != result.String() {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedMessage, result.String())
	}
}

//<sequence>
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsFalseForOptionalSequenceWithNoLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", true, testLog),
		fielduint32.New(properties.New(1, "SequenceField", true, testLog)),
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
