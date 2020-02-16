package field

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
)

//<sequence>
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestCanDeseraliseRequiredSequenceOfLengthZero(t *testing.T) {
	// Arrange length = 10000000
	messageAsBytes := bytes.NewBuffer([]byte{128})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
	expectedMessage := fix.NewSequenceValue(0)
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
func TestCanDeseraliseRequiredSequenceOfLengthTwo(t *testing.T) {
	// Arrange length(2) = 10000010
	// 1: int64 = 10000011	string(TEST1) = 01010100 01000101 01010011 01010100 10110001
	// 2: int64 = 10000010	string(TEST2) = 01010100 01000101 01010011 01010100 10110010
	messageAsBytes := bytes.NewBuffer([]byte{130, 131, 84, 69, 83, 84, 177, 130, 84, 69, 83, 84, 178})
	pmap, _ := presencemap.New(bytes.NewBuffer([]byte{128}))
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	expectedMessage := fix.NewSequenceValue(0)
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			Sequence{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				LengthField: UInt32{
					FieldDetails: Field{
						ID:       4,
						Required: true,
					},
					Operation: operation.None{},
				},
				SequenceFields: []store.Unit{
					AsciiString{
						FieldDetails: Field{
							ID:       5,
							Required: true,
						},
						Operation: operation.None{},
					},
				},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			Sequence{
				FieldDetails: Field{
					ID:       3,
					Required: false,
				},
				LengthField: UInt32{
					FieldDetails: Field{
						ID:       4,
						Required: false,
					},
					Operation: operation.None{},
				},
				SequenceFields: []store.Unit{
					AsciiString{
						FieldDetails: Field{
							ID:       5,
							Required: true,
						},
						Operation: operation.None{},
					},
				},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: false,
				},
				Operation: operation.Constant{
					ConstantValue: int64(2),
				},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.Constant{
					ConstantValue: int64(2),
				},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: false,
				},
				Operation: operation.Constant{
					ConstantValue: int64(2),
				},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.Constant{
					ConstantValue: int64(2),
				},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.Constant{
				ConstantValue: uint32(1),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.Constant{
				ConstantValue: uint32(1),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.Constant{
				ConstantValue: uint32(1),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result, err := unitUnderTest.Deserialise(messageAsBytes, &pmap)
	if err != nil {
		t.Errorf("Got an error when none was expected: %s", err)
	}

	// Assert
	if result.Get() != nil {
		t.Errorf("Did not read the expected null value, expected: nil, result: %#v", result.Get())
	}
}

//<sequence>
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsFalseForRequiredSequenceWithNoLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
	}
}

//<sequence>
//	<length />
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsFalseForOptionalSequenceWithNoLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.None{},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != false {
		t.Errorf("Expected RequiresPmap to return false, but got true")
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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.Constant{
				ConstantValue: uint(3),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.Constant{
				ConstantValue: uint(3),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}

//<sequence>
//	<length>
//		<default value="3" />
//	</length>
// 	<int64 id="2"/>
// 	<string id="3"/>
//</sequence>
func TestRequiresPmapReturnsTrueForRequiredSequenceWithDefaultLengthOperator(t *testing.T) {
	// Arrange
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: true,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: true,
			},
			Operation: operation.Default{
				DefaultValue: uint(3),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

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
	unitUnderTest := Sequence{
		FieldDetails: Field{
			ID:       1,
			Required: false,
		},
		LengthField: UInt32{
			FieldDetails: Field{
				ID:       1,
				Required: false,
			},
			Operation: operation.Default{
				DefaultValue: uint(3),
			},
		},
		SequenceFields: []store.Unit{
			Int64{
				FieldDetails: Field{
					ID:       2,
					Required: true,
				},
				Operation: operation.None{},
			},
			AsciiString{
				FieldDetails: Field{
					ID:       3,
					Required: true,
				},
				Operation: operation.None{},
			},
		},
	}

	// Act
	result := unitUnderTest.RequiresPmap()

	// Assert
	if result != true {
		t.Errorf("Expected RequiresPmap to return true, but got false")
	}
}
