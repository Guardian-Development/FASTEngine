package loader

import (
	"os"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
	"github.com/Guardian-Development/fastengine/internal/fast/field"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
)

func TestCanLoadAllSupportedTypesFromTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					field.AsciiString{
						FieldDetails: field.Field{
							ID:       1,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Required: true,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Required: true,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Required: true,
							},
							Operation: operation.None{},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Required: true,
								},
								Operation: operation.None{},
							},
							field.UInt32{
								FieldDetails: field.Field{
									ID:       13,
									Required: true,
								},
								Operation: operation.None{},
							},
						},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       14,
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       0,
								Required: true,
							},
							Operation: operation.None{},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       15,
									Required: true,
								},
								Operation: operation.None{},
							},
							field.UInt32{
								FieldDetails: field.Field{
									ID:       16,
									Required: true,
								},
								Operation: operation.None{},
							},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := Load(file)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, store)
	}
}

func TestCanLoadAllSupportedOptionalTypesFromTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_all_supported_optional_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					field.AsciiString{
						FieldDetails: field.Field{
							ID:       1,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Required: false,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Required: false,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Required: false,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Required: false,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Required: false,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Required: false,
							},
							Operation: operation.None{},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Required: true,
								},
								Operation: operation.None{},
							},
							field.UInt32{
								FieldDetails: field.Field{
									ID:       13,
									Required: true,
								},
								Operation: operation.None{},
							},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := Load(file)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, store)
	}
}

func TestCanLoadConstantOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_constant_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					field.AsciiString{
						FieldDetails: field.Field{
							ID:       1,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: "Hello"},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: uint32(10)},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: int32(-10)},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: uint64(10)},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: int64(-10)},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int32(-1)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int64(57)},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int32(-2)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int64(2)},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: "Hello: ϔ"},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: []byte{0x54, 0x45, 0x53, 0x54, 0x3F}},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: uint32(2)},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Required: true,
								},
								Operation: operation.None{},
							},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := Load(file)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, store)
	}
}
