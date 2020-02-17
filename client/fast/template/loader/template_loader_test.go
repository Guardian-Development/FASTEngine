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
							Name:     "StringDefaultAscii",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Name:     "unsigned int32",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Name:     "signed int32",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Name:     "unsigned int64",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Name:     "signed int64",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Name:     "decimal",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalExponent",
								Required: true,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalMantissa",
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Name:     "decimal with exp/man",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "custom decimal exp",
								Required: true,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "custom decimal man",
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Name:     "StringUnicode",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Name:     "byteVector",
							Required: true,
						},
						Operation: operation.None{},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Name:     "sequence",
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Name:     "length",
								Required: true,
							},
							Operation: operation.None{},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Name:     "sequence field 1",
									Required: true,
								},
								Operation: operation.None{},
							},
							field.UInt32{
								FieldDetails: field.Field{
									ID:       13,
									Name:     "sequence field 2",
									Required: true,
								},
								Operation: operation.None{},
							},
						},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       14,
							Name:     "sequence implicit length",
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       0,
								Name:     "sequence implicit length",
								Required: true,
							},
							Operation: operation.None{},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       15,
									Name:     "sequence field 1",
									Required: true,
								},
								Operation: operation.None{},
							},
							field.UInt32{
								FieldDetails: field.Field{
									ID:       16,
									Name:     "sequence field 2",
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
							Name:     "String",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Name:     "unsigned int32",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Name:     "signed int32",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Name:     "unsigned int64",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Name:     "signed int64",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Name:     "decimal",
							Required: false,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalExponent",
								Required: false,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalMantissa",
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Name:     "decimal with exp/man",
							Required: false,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manExponent",
								Required: false,
							},
							Operation: operation.None{},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manMantissa",
								Required: true,
							},
							Operation: operation.None{},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Name:     "StringUnicode",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Name:     "byteVector",
							Required: false,
						},
						Operation: operation.None{},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Name:     "sequence",
							Required: false,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Name:     "length",
								Required: false,
							},
							Operation: operation.None{},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Name:     "sequence field 1",
									Required: true,
								},
								Operation: operation.None{},
							},
							field.UInt32{
								FieldDetails: field.Field{
									ID:       13,
									Name:     "sequence field 2",
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
							Name:     "String",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: "Hello"},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Name:     "unsigned int32",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: uint32(10)},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Name:     "signed int32",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: int32(-10)},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Name:     "unsigned int64",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: uint64(10)},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Name:     "signed int64",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: int64(-10)},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Name:     "decimal",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalExponent",
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int32(-1)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalMantissa",
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int64(57)},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Name:     "decimal with exp/man",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manExponent",
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int32(-2)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manMantissa",
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: int64(2)},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Name:     "StringUnicode",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: "Hello: ϔ"},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Name:     "byteVector",
							Required: true,
						},
						Operation: operation.Constant{ConstantValue: []byte{0x54, 0x45, 0x53, 0x54, 0x3F}},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Name:     "sequence",
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Name:     "length",
								Required: true,
							},
							Operation: operation.Constant{ConstantValue: uint32(2)},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Name:     "sequence field 1",
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

func TestCanLoadDefaultOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_default_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					field.AsciiString{
						FieldDetails: field.Field{
							ID:       1,
							Name:     "String",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: "Hello"},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Name:     "unsigned int32",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: uint32(10)},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Name:     "signed int32",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: int32(-10)},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Name:     "unsigned int64",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: uint64(10)},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Name:     "signed int64",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: int64(-10)},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Name:     "decimal",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalExponent",
								Required: true,
							},
							Operation: operation.Default{DefaultValue: int32(-1)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalMantissa",
								Required: true,
							},
							Operation: operation.Default{DefaultValue: int64(57)},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Name:     "decimal with exp/man",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manExponent",
								Required: true,
							},
							Operation: operation.Default{DefaultValue: int32(-2)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manMantissa",
								Required: true,
							},
							Operation: operation.Default{DefaultValue: int64(2)},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Name:     "StringUnicode",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: "Hello: ϔ"},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Name:     "byteVector",
							Required: true,
						},
						Operation: operation.Default{DefaultValue: []byte{0x54, 0x45, 0x53, 0x54, 0x3F}},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Name:     "sequence",
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Name:     "length",
								Required: true,
							},
							Operation: operation.Default{DefaultValue: uint32(2)},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Name:     "sequence field 1",
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

func TestCanLoadCopyOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_copy_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					field.AsciiString{
						FieldDetails: field.Field{
							ID:       1,
							Name:     "String",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: "Hello"},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:       2,
							Name:     "unsigned int32",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: uint32(10)},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:       3,
							Name:     "signed int32",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: int32(-10)},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:       4,
							Name:     "unsigned int64",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: uint64(10)},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:       5,
							Name:     "signed int64",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: int64(-10)},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       6,
							Name:     "decimal",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalExponent",
								Required: true,
							},
							Operation: operation.Copy{InitialValue: int32(-1)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       6,
								Name:     "decimalMantissa",
								Required: true,
							},
							Operation: operation.Copy{InitialValue: int64(57)},
						},
					},
					field.Decimal{
						FieldDetails: field.Field{
							ID:       7,
							Name:     "decimal with exp/man",
							Required: true,
						},
						ExponentField: field.Int32{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manExponent",
								Required: true,
							},
							Operation: operation.Copy{InitialValue: int32(-2)},
						},
						MantissaField: field.Int64{
							FieldDetails: field.Field{
								ID:       7,
								Name:     "decimal with exp/manMantissa",
								Required: true,
							},
							Operation: operation.Copy{InitialValue: int64(2)},
						},
					},
					field.UnicodeString{
						FieldDetails: field.Field{
							ID:       8,
							Name:     "StringUnicode",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: "Hello: ϔ"},
					},
					field.ByteVector{
						FieldDetails: field.Field{
							ID:       9,
							Name:     "byteVector",
							Required: true,
						},
						Operation: operation.Copy{InitialValue: []byte{0x54, 0x45, 0x53, 0x54, 0x3F}},
					},
					field.Sequence{
						FieldDetails: field.Field{
							ID:       10,
							Name:     "sequence",
							Required: true,
						},
						LengthField: field.UInt32{
							FieldDetails: field.Field{
								ID:       11,
								Name:     "length",
								Required: true,
							},
							Operation: operation.Copy{InitialValue: uint32(2)},
						},
						SequenceFields: []store.Unit{
							field.AsciiString{
								FieldDetails: field.Field{
									ID:       12,
									Name:     "sequence field 1",
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
