package loader

import (
	"os"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldbytevector"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielddecimal"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldsequence"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fielduint64"
	"github.com/Guardian-Development/fastengine/internal/fast/field/fieldunicodestring"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"

	"github.com/Guardian-Development/fastengine/client/fast/template/store"
)

func TestCanLoadAllSupportedTypesFromTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					fieldasciistring.New(properties.New(1, "StringDefaultAscii", true)),
					fielduint32.New(properties.New(2, "unsigned int32", true)),
					fieldint32.New(properties.New(3, "signed int32", true)),
					fielduint64.New(properties.New(4, "unsigned int64", true)),
					fieldint64.New(properties.New(5, "signed int64", true)),
					fielddecimal.New(properties.New(6, "decimal", true),
						fieldint32.New(properties.New(6, "decimalExponent", true)),
						fieldint64.New(properties.New(6, "decimalMantissa", true))),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true),
						fieldint32.New(properties.New(7, "custom decimal exp", true)),
						fieldint64.New(properties.New(7, "custom decimal man", true))),
					fieldunicodestring.New(properties.New(8, "StringUnicode", true)),
					fieldbytevector.New(properties.New(9, "byteVector", true)),
					fieldsequence.New(properties.New(10, "sequence", true),
						fielduint32.New(properties.New(11, "length", true)),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true)),
							fielduint32.New(properties.New(13, "sequence field 2", true)),
						}),
					fieldsequence.New(properties.New(14, "sequence implicit length", true),
						fielduint32.New(properties.New(0, "sequence implicit length", true)),
						[]store.Unit{
							fieldasciistring.New(properties.New(15, "sequence field 1", true)),
							fielduint32.New(properties.New(16, "sequence field 2", true)),
						}),
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
					fieldasciistring.New(properties.New(1, "String", false)),
					fielduint32.New(properties.New(2, "unsigned int32", false)),
					fieldint32.New(properties.New(3, "signed int32", false)),
					fielduint64.New(properties.New(4, "unsigned int64", false)),
					fieldint64.New(properties.New(5, "signed int64", false)),
					fielddecimal.New(properties.New(6, "decimal", false),
						fieldint32.New(properties.New(6, "decimalExponent", false)),
						fieldint64.New(properties.New(6, "decimalMantissa", true))),
					fielddecimal.New(properties.New(7, "decimal with exp/man", false),
						fieldint32.New(properties.New(7, "decimal with exp/manExponent", false)),
						fieldint64.New(properties.New(7, "decimal with exp/manMantissa", true))),
					fieldunicodestring.New(properties.New(8, "StringUnicode", false)),
					fieldbytevector.New(properties.New(9, "byteVector", false)),
					fieldsequence.New(properties.New(10, "sequence", false),
						fielduint32.New(properties.New(11, "length", false)),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true)),
							fielduint32.New(properties.New(13, "sequence field 2", true)),
						}),
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
					fieldasciistring.NewConstantOperation(properties.New(1, "String", true), "Hello"),
					fielduint32.NewConstantOperation(properties.New(2, "unsigned int32", true), 10),
					fieldint32.NewConstantOperation(properties.New(3, "signed int32", true), -10),
					fielduint64.NewConstantOperation(properties.New(4, "unsigned int64", true), 10),
					fieldint64.NewConstantOperation(properties.New(5, "signed int64", true), -10),
					fielddecimal.New(properties.New(6, "decimal", true),
						fieldint32.NewConstantOperation(properties.New(6, "decimalExponent", true), -1),
						fieldint64.NewConstantOperation(properties.New(6, "decimalMantissa", true), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true),
						fieldint32.NewConstantOperation(properties.New(7, "decimal with exp/manExponent", true), -2),
						fieldint64.NewConstantOperation(properties.New(7, "decimal with exp/manMantissa", true), 2)),
					fieldunicodestring.NewConstantOperation(properties.New(8, "StringUnicode", true), "Hello: ϔ"),
					fieldbytevector.NewConstantOperation(properties.New(9, "byteVector", true), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true),
						fielduint32.NewConstantOperation(properties.New(11, "length", true), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true)),
						}),
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
					fieldasciistring.NewDefaultOperationWithValue(properties.New(1, "String", true), "Hello"),
					fielduint32.NewDefaultOperationWithValue(properties.New(2, "unsigned int32", true), 10),
					fieldint32.NewDefaultOperationWithValue(properties.New(3, "signed int32", true), -10),
					fielduint64.NewDefaultOperationWithValue(properties.New(4, "unsigned int64", true), 10),
					fieldint64.NewDefaultOperationWithValue(properties.New(5, "signed int64", true), -10),
					fielddecimal.New(properties.New(6, "decimal", true),
						fieldint32.NewDefaultOperationWithValue(properties.New(6, "decimalExponent", true), -1),
						fieldint64.NewDefaultOperationWithValue(properties.New(6, "decimalMantissa", true), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true),
						fieldint32.NewDefaultOperationWithValue(properties.New(7, "decimal with exp/manExponent", true), -2),
						fieldint64.NewDefaultOperationWithValue(properties.New(7, "decimal with exp/manMantissa", true), 2)),
					fieldunicodestring.NewDefaultOperationWithValue(properties.New(8, "StringUnicode", true), "Hello: ϔ"),
					fieldbytevector.NewDefaultOperationWithValue(properties.New(9, "byteVector", true), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true),
						fielduint32.NewDefaultOperationWithValue(properties.New(11, "length", true), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true)),
						}),
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
					fieldasciistring.NewCopyOperationWithInitialValue(properties.New(1, "String", true), "Hello"),
					fielduint32.NewCopyOperationWithInitialValue(properties.New(2, "unsigned int32", true), 10),
					fieldint32.NewCopyOperationWithInitialValue(properties.New(3, "signed int32", true), -10),
					fielduint64.NewCopyOperationWithInitialValue(properties.New(4, "unsigned int64", true), 10),
					fieldint64.NewCopyOperationWithInitialValue(properties.New(5, "signed int64", true), -10),
					fielddecimal.New(properties.New(6, "decimal", true),
						fieldint32.NewCopyOperationWithInitialValue(properties.New(6, "decimalExponent", true), -1),
						fieldint64.NewCopyOperationWithInitialValue(properties.New(6, "decimalMantissa", true), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true),
						fieldint32.NewCopyOperationWithInitialValue(properties.New(7, "decimal with exp/manExponent", true), -2),
						fieldint64.NewCopyOperationWithInitialValue(properties.New(7, "decimal with exp/manMantissa", true), 2)),
					fieldunicodestring.NewCopyOperationWithInitialValue(properties.New(8, "StringUnicode", true), "Hello: ϔ"),
					fieldbytevector.NewCopyOperationWithInitialValue(properties.New(9, "byteVector", true), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true),
						fielduint32.NewCopyOperationWithInitialValue(properties.New(11, "length", true), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true)),
						}),
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

func TestCanLoadIncrementOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_increment_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					fielduint32.NewIncrementOperationWithInitialValue(properties.New(1, "unsigned int32", true), 10),
					fieldint32.NewIncrementOperationWithInitialValue(properties.New(2, "signed int32", true), -10),
					fielduint64.NewIncrementOperationWithInitialValue(properties.New(3, "unsigned int64", true), 10),
					fieldint64.NewIncrementOperationWithInitialValue(properties.New(4, "signed int64", true), -10),
					fieldsequence.New(properties.New(5, "sequence", true),
						fielduint32.NewIncrementOperationWithInitialValue(properties.New(6, "length", true), 2),
						[]store.Unit{
							fielduint32.NewIncrementOperation(properties.New(7, "sequence field 1", true)),
						}),
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

func TestCanLoadTailOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_tail_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					fieldasciistring.NewTailOperationWithInitialValue(properties.New(1, "String", true), "Hello"),
					fieldunicodestring.NewTailOperationWithInitialValue(properties.New(2, "StringUnicode", true), "Hello: ϔ"),
					fieldbytevector.NewTailOperationWithInitialValue(properties.New(3, "byteVector", true), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
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

func TestCanLoadDeltaOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_delta_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: store.Template{
				TemplateUnits: []store.Unit{
					fieldasciistring.NewDeltaOperationWithInitialValue(properties.New(1, "String", true), "Hello"),
					fielduint32.NewDeltaOperationWithInitialValue(properties.New(2, "unsigned int32", true), 10),
					fieldint32.NewDeltaOperationWithInitialValue(properties.New(3, "signed int32", true), -10),
					fielduint64.NewDeltaOperationWithInitialValue(properties.New(4, "unsigned int64", true), 10),
					fieldint64.NewDeltaOperationWithInitialValue(properties.New(5, "signed int64", true), -10),
					fielddecimal.New(properties.New(6, "decimal", true),
						fieldint32.NewDeltaOperationWithInitialValue(properties.New(6, "decimalExponent", true), -1),
						fieldint64.NewDeltaOperationWithInitialValue(properties.New(6, "decimalMantissa", true), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true),
						fieldint32.NewDeltaOperationWithInitialValue(properties.New(7, "decimal with exp/manExponent", true), -2),
						fieldint64.NewDeltaOperationWithInitialValue(properties.New(7, "decimal with exp/manMantissa", true), 2)),
					fieldunicodestring.NewDeltaOperationWithInitialValue(properties.New(8, "StringUnicode", true), "Hello: ϔ"),
					fieldbytevector.NewDeltaOperationWithInitialValue(properties.New(9, "byteVector", true), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true),
						fielduint32.NewDeltaOperationWithInitialValue(properties.New(11, "length", true), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true)),
						}),
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
