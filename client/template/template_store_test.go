package template

import (
	"os"
	"reflect"
	"testing"
)

func TestCanLoadHeartbeatTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../test/test_heartbeat_template.xml")
	expectedStore := Store{
		Templates: []Template{
			Template{
				TemplateUnits: []Unit{
					FieldString{
						fieldDetails: Field{
							ID: 1128,
							Operation: OperationConstant{
								constantValue: "9",
							},
						},
					},
					FieldString{
						fieldDetails: Field{
							ID: 35,
							Operation: OperationConstant{
								constantValue: "0",
							},
						},
					},
					FieldUInt32{
						fieldDetails: Field{
							ID:        34,
							Operation: OperationNone{},
						},
					},
					FieldUInt64{
						fieldDetails: Field{
							ID:        52,
							Operation: OperationNone{},
						},
					},
				},
			},
		},
	}

	// Act
	store, err := Create(file)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the heartbeat template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, store)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:%s\nactual:%s", expectedStore, store)
	}
}
