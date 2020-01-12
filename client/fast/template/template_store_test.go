package template

import (
	"os"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/field"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
)

func TestCanLoadHeartbeatTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../test/test_heartbeat_template.xml")
	expectedStore := Store{
		Templates: map[uint32]Template{
			144: Template{
				TemplateUnits: []Unit{
					field.String{
						FieldDetails: field.Field{
							ID:       1128,
							Required: true,
							Operation: operation.Constant{
								ConstantValue: "9",
							},
						},
					},
					field.String{
						FieldDetails: field.Field{
							ID:       35,
							Required: true,
							Operation: operation.Constant{
								ConstantValue: "0",
							},
						},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:        34,
							Required:  true,
							Operation: operation.None{},
						},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:        52,
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
		t.Errorf("Got an error loading the heartbeat template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:%v\nactual:%v", expectedStore, store)
	}
}

func TestCanLoadOptionalValueTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../test/test_optional_value_template.xml")
	expectedStore := Store{
		Templates: map[uint32]Template{
			144: Template{
				TemplateUnits: []Unit{
					field.String{
						FieldDetails: field.Field{
							ID:       1128,
							Required: true,
							Operation: operation.Constant{
								ConstantValue: "9",
							},
						},
					},
					field.String{
						FieldDetails: field.Field{
							ID:       35,
							Required: true,
							Operation: operation.Constant{
								ConstantValue: "0",
							},
						},
					},
					field.UInt32{
						FieldDetails: field.Field{
							ID:        34,
							Required:  false,
							Operation: operation.None{},
						},
					},
					field.UInt64{
						FieldDetails: field.Field{
							ID:        52,
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
		t.Errorf("Got an error loading the optional value template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:%v\nactual:%v", expectedStore, store)
	}
}