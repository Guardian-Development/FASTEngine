package fieldsequence

import (
	"bytes"
	"github.com/Guardian-Development/fastengine/pkg/fast/dictionary"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/fast/presencemap"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
	"github.com/Guardian-Development/fastengine/pkg/fix"
)

//<sequence id="1">
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceOfLengthZero(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := fix.NewSequenceValue(0)
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.New(properties.New(1, "SequenceField", true)),
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
		fielduint32.New(properties.New(1, "SequenceField", true)),
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
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceOfLengthZero(t *testing.T) {
	// Arrange length = 10000001
	messageAsBytes := bytes.NewBuffer([]byte{129})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	expectedMessage := fix.NewSequenceValue(0)
	unitUnderTest := New(
		properties.New(1, "SequenceField", false),
		fielduint32.New(properties.New(1, "SequenceField", false)),
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
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseOptionalSequenceNull(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	dictionary := dictionary.New()
	unitUnderTest := New(
		properties.New(1, "SequenceField", false),
		fielduint32.New(properties.New(1, "SequenceField", false)),
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
		properties.New(1, "SequenceField", false),
		fielduint32.New(properties.New(1, "SequenceField", false)),
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
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(3)),
					3: fix.SequenceValue{
						Values: []fix.Message{
							fix.Message{
								Tags: map[uint64]fix.Value{
									5: fix.NewRawValue("TEST1"),
								},
							},
							fix.Message{
								Tags: map[uint64]fix.Value{
									5: fix.NewRawValue("TEST2"),
								},
							},
						},
					},
				},
			},
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(4)),
					3: fix.SequenceValue{
						Values: []fix.Message{
							fix.Message{
								Tags: map[uint64]fix.Value{
									5: fix.NewRawValue("TEST3"),
								},
							},
						},
					},
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.New(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			New(properties.New(3, "SequenceTwoField", true),
				fielduint32.New(properties.New(4, "SequenceTwoField", true)),
				[]store.Unit{
					fieldasciistring.New(properties.New(5, "AsciiStringTwoField", true)),
				}),
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
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(3)),
					3: fix.SequenceValue{
						Values: []fix.Message{
							fix.Message{
								Tags: map[uint64]fix.Value{
									5: fix.NewRawValue("TEST1"),
								},
							},
						},
					},
				},
			},
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(4)),
					3: fix.NullValue{},
				},
			},
		},
	}
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.New(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.New(properties.New(2, "Int64Field", true)),
			New(properties.New(3, "SequenceTwoField", false),
				fielduint32.New(properties.New(4, "SequenceTwoField", false)),
				[]store.Unit{
					fieldasciistring.New(properties.New(5, "AsciiStringTwoField", true)),
				}),
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
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NullValue{},
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
		fielduint32.New(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", false), 2),
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
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(2)),
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
		fielduint32.New(properties.New(1, "SequenceField", true)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", true), 2),
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
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NullValue{},
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
		properties.New(1, "SequenceField", false),
		fielduint32.New(properties.New(1, "SequenceField", false)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", false), 2),
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
	dictionary := dictionary.New()
	expectedMessage := fix.SequenceValue{
		Values: []fix.Message{
			fix.Message{
				Tags: map[uint64]fix.Value{
					2: fix.NewRawValue(int64(2)),
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
		properties.New(1, "SequenceField", false),
		fielduint32.New(properties.New(1, "SequenceField", false)),
		[]store.Unit{
			fieldint64.NewConstantOperation(properties.New(2, "Int64Field", true), 2),
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
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsFalseForOptionalSequenceWithNoLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := New(
		properties.New(1, "SequenceField", true),
		fielduint32.New(properties.New(1, "SequenceField", true)),
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
