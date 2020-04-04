package fieldsequence

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/pkg/fix"
	"github.com/Guardian-Development/fastengine/pkg/template/store"
)

//<sequence>
//	<length>
//		<copy/>
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestDictionaryIsUpdatedWithAssignedValueWhenSequenceLengthReadFromStream(t *testing.T) {
	// Arrange pmap = 1100000 length(2) = 10000010
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{130, 131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedValue := dictionary.AssignedValue{Value: fix.NewRawValue(uint32(2))}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.NewCopyOperation(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("SequenceField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<sequence presence="optional">
//	<length>
//		<copy/>
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestDictionaryIsUpdatedWithEmptyValueWhenNilSequenceLengthReadFromStream(t *testing.T) {
	// Arrange pmap = 1100000 length(2) = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dict := dictionary.New()
	expectedValue := dictionary.EmptyValue{}
	unitUnderTest := New(
		properties.New(1, "SequenceField", false),
		fielduint32.NewCopyOperation(properties.New(1, "SequenceField", false)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	unitUnderTest.Deserialise(messageAsBytes, &pmap, &dict)

	// Assert
	result := dict.GetValue("SequenceField")
	if result != expectedValue {
		t.Errorf("Expected value and deserialised value were not equal, expected: %v, actual: %v", expectedValue, result)
	}
}

//<sequence>
//	<length>
//		<copy/>
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceCopyOperatorLengthEncoded(t *testing.T) {
	// Arrange pmap = 1100000 length(2) = 10000010
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{130, 131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{192}))
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(3)),
					3: fix.NewRawValue("TEST1"),
				},
			},
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(2)),
					3: fix.NewRawValue("TEST2"),
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.NewCopyOperation(properties.New(1, "SequenceField", true)),
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

//<sequence>
//	<length>
//		<copy/>
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceCopyOperatorLengthNotEncodedGetsPreviousValue(t *testing.T) {
	// Arrange pmap = 1000000 length(2) = nil
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(3)),
					3: fix.NewRawValue("TEST1"),
				},
			},
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(2)),
					3: fix.NewRawValue("TEST2"),
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.NewCopyOperation(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	dict.SetValue("SequenceField", fix.NewRawValue(uint32(2)))
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

//<sequence>
//	<length>
//		<copy value="2"/>
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceCopyOperatorLengthNotEncodedGetsInitialValue(t *testing.T) {
	// Arrange pmap = 1000000 length(2) = nil
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(3)),
					3: fix.NewRawValue("TEST1"),
				},
			},
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(2)),
					3: fix.NewRawValue("TEST2"),
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.NewCopyOperationWithInitialValue(properties.New(1, "SequenceField", true), 2),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
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
//	<length>
//		<copy />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceCopyOperatorLengthNotEncodedGetsPreviousValueNil(t *testing.T) {
	// Arrange pmap = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dict := dictionary.New()
	expectedMessage := fix.NullValue{}
	unitUnderTest := New(
		properties.New(1, "SequenceField", false),
		fielduint32.NewCopyOperation(properties.New(1, "SequenceField", false)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			fieldasciistring.New(properties.New(3, "AsciiStringField", true)),
		})

	// Act
	dict.SetValue("SequenceField", fix.NullValue{})
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

//<sequence>
//	<length />
// 	<int64 id="2">
//		<copy value="1" />
//	</int64>
// 	<string id="3">
//		<copy />
//	</string>
//</sequence>
func TestShouldUsePreviousValueWhenElementNotEncodedWithinSequenceAndCopyOperatorPresent(t *testing.T) {
	// Arrange length(3) = 10000011
	// 1: pmap = 10100000 int64(1) = nil (use initial)	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: pmap = 10100000 int64(1) = nil (copy initial)	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	// 3: pmap = 11000000 int64(3) = 10000011			string = nil (copy TEST2)
	messageAsBytes := bytes.NewBuffer([]byte{131,
		160, 84, 69, 83, 84, 177,
		160, 84, 69, 83, 84, 178,
		192, 131,
	})
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
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(1)),
					3: fix.NewRawValue("TEST2"),
				},
			},
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(3)),
					3: fix.NewRawValue("TEST2"),
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.New(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.NewCopyOperationWithInitialValue(properties.New(2, "Int64Field", true), 1),
			fieldasciistring.NewCopyOperation(properties.New(3, "AsciiStringField", true)),
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

//<sequence>
//	<length>
//		<copy value="1" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsTrueForRequiredSequenceWithCopyLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.NewCopyOperationWithInitialValue(properties.New(1, "SequenceField", true), 1),
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

//<sequence presence="optional">
//	<length>
//		<copy value="1" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsTrueForOptionalSequenceWithCopyLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", false),
		fielduint32.NewCopyOperationWithInitialValue(properties.New(1, "SequenceField", false), 1),
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
