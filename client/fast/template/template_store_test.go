package template

import (
	"os"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/field"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
)

func TestCanLoadAllSupportedTypesFromTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../test/template-loader-tests/test_load_all_supported_types.xml")
	expectedStore := Store{
		Templates: map[uint32]Template{
			144: Template{
				TemplateUnits: []Unit{
					field.String{
						FieldDetails: field.Field{
							ID:        1,
							Required:  true,
							Operation: operation.None{},
						},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:        2,
							Required:  true,
							Operation: operation.None{},
						},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:        3,
							Required:  true,
							Operation: operation.None{},
						},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:        4,
							Required:  true,
							Operation: operation.None{},
						},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:        5,
							Required:  true,
							Operation: operation.None{},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := New(file)

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
	file, _ := os.Open("../../../test/template-loader-tests/test_load_all_supported_optional_types.xml")
	expectedStore := Store{
		Templates: map[uint32]Template{
			144: Template{
				TemplateUnits: []Unit{
					field.String{
						FieldDetails: field.Field{
							ID:        1,
							Required:  false,
							Operation: operation.None{},
						},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:        2,
							Required:  false,
							Operation: operation.None{},
						},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:        3,
							Required:  false,
							Operation: operation.None{},
						},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:        4,
							Required:  false,
							Operation: operation.None{},
						},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:        5,
							Required:  false,
							Operation: operation.None{},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := New(file)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, store)
	}
}

func TestCanLoadAllSupportedOperationsFromTemplateFile(t *testing.T) {
	t.Skip("TODO BUG: constant not loading type information, need to refactor loader")

	// Arrange
	file, _ := os.Open("../../../test/template-loader-tests/test_load_all_supported_operations.xml")
	expectedStore := Store{
		Templates: map[uint32]Template{
			144: Template{
				TemplateUnits: []Unit{
					field.String{
						FieldDetails: field.Field{
							ID:        1,
							Required:  true,
							Operation: operation.Constant{ConstantValue: "Hello"},
						},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:        2,
							Required:  true,
							Operation: operation.Constant{ConstantValue: uint32(10)},
						},
					},
					field.Int32{
						FieldDetails: field.Field{
							ID:        3,
							Required:  true,
							Operation: operation.Constant{ConstantValue: int32(-10)},
						},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:        4,
							Required:  true,
							Operation: operation.Constant{ConstantValue: uint64(10)},
						},
					},
					field.Int64{
						FieldDetails: field.Field{
							ID:        5,
							Required:  true,
							Operation: operation.Constant{ConstantValue: int64(-10)},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := New(file)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, store)
	}
}
