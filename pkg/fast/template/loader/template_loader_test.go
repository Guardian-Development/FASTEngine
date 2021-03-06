package loader

import (
	"github.com/Guardian-Development/fastengine/pkg/fast/template/store"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldasciistring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldbytevector"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielddecimal"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldsequence"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint32"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fielduint64"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldunicodestring"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
)

var testLog = log.New(os.Stdout, "", log.LstdFlags)

func TestCanLoadAllSupportedTypesFromTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.New(properties.New(1, "StringDefaultAscii", true, testLog)),
					fielduint32.New(properties.New(2, "unsigned int32", true, testLog)),
					fieldint32.New(properties.New(3, "signed int32", true, testLog)),
					fielduint64.New(properties.New(4, "unsigned int64", true, testLog)),
					fieldint64.New(properties.New(5, "signed int64", true, testLog)),
					fielddecimal.New(properties.New(6, "decimal", true, testLog),
						fieldint32.New(properties.New(6, "decimalExponent", true, testLog)),
						fieldint64.New(properties.New(6, "decimalMantissa", true, testLog))),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true, testLog),
						fieldint32.New(properties.New(7, "custom decimal exp", true, testLog)),
						fieldint64.New(properties.New(7, "custom decimal man", true, testLog))),
					fieldunicodestring.New(properties.New(8, "StringUnicode", true, testLog)),
					fieldbytevector.New(properties.New(9, "byteVector", true, testLog)),
					fieldsequence.New(properties.New(10, "sequence", true, testLog),
						fielduint32.New(properties.New(11, "length", true, testLog)),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true, testLog)),
							fielduint32.New(properties.New(13, "sequence field 2", true, testLog)),
						}),
					fieldsequence.New(properties.New(14, "sequence implicit length", true, testLog),
						fielduint32.New(properties.New(0, "sequence implicit length", true, testLog)),
						[]store.Unit{
							fieldasciistring.New(properties.New(15, "sequence field 1", true, testLog)),
							fielduint32.New(properties.New(16, "sequence field 2", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%#v\nactual:\t\t%#v", expectedStore, loadedStore)
	}
}

func TestCanLoadAllSupportedOptionalTypesFromTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_all_supported_optional_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.New(properties.New(1, "String", false, testLog)),
					fielduint32.New(properties.New(2, "unsigned int32", false, testLog)),
					fieldint32.New(properties.New(3, "signed int32", false, testLog)),
					fielduint64.New(properties.New(4, "unsigned int64", false, testLog)),
					fieldint64.New(properties.New(5, "signed int64", false, testLog)),
					fielddecimal.New(properties.New(6, "decimal", false, testLog),
						fieldint32.New(properties.New(6, "decimalExponent", false, testLog)),
						fieldint64.New(properties.New(6, "decimalMantissa", true, testLog))),
					fielddecimal.New(properties.New(7, "decimal with exp/man", false, testLog),
						fieldint32.New(properties.New(7, "decimal with exp/manExponent", false, testLog)),
						fieldint64.New(properties.New(7, "decimal with exp/manMantissa", true, testLog))),
					fieldunicodestring.New(properties.New(8, "StringUnicode", false, testLog)),
					fieldbytevector.New(properties.New(9, "byteVector", false, testLog)),
					fieldsequence.New(properties.New(10, "sequence", false, testLog),
						fielduint32.New(properties.New(11, "length", false, testLog)),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true, testLog)),
							fielduint32.New(properties.New(13, "sequence field 2", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}

func TestCanLoadConstantOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_constant_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.NewConstantOperation(properties.New(1, "String", true, testLog), "Hello"),
					fielduint32.NewConstantOperation(properties.New(2, "unsigned int32", true, testLog), 10),
					fieldint32.NewConstantOperation(properties.New(3, "signed int32", true, testLog), -10),
					fielduint64.NewConstantOperation(properties.New(4, "unsigned int64", true, testLog), 10),
					fieldint64.NewConstantOperation(properties.New(5, "signed int64", true, testLog), -10),
					fielddecimal.New(properties.New(6, "decimal", true, testLog),
						fieldint32.NewConstantOperation(properties.New(6, "decimalExponent", true, testLog), -1),
						fieldint64.NewConstantOperation(properties.New(6, "decimalMantissa", true, testLog), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true, testLog),
						fieldint32.NewConstantOperation(properties.New(7, "decimal with exp/manExponent", true, testLog), -2),
						fieldint64.NewConstantOperation(properties.New(7, "decimal with exp/manMantissa", true, testLog), 2)),
					fieldunicodestring.NewConstantOperation(properties.New(8, "StringUnicode", true, testLog), "Hello: ϔ"),
					fieldbytevector.NewConstantOperation(properties.New(9, "byteVector", true, testLog), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true, testLog),
						fielduint32.NewConstantOperation(properties.New(11, "length", true, testLog), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}

func TestCanLoadDefaultOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_default_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.NewDefaultOperationWithValue(properties.New(1, "String", true, testLog), "Hello"),
					fielduint32.NewDefaultOperationWithValue(properties.New(2, "unsigned int32", true, testLog), 10),
					fieldint32.NewDefaultOperationWithValue(properties.New(3, "signed int32", true, testLog), -10),
					fielduint64.NewDefaultOperationWithValue(properties.New(4, "unsigned int64", true, testLog), 10),
					fieldint64.NewDefaultOperationWithValue(properties.New(5, "signed int64", true, testLog), -10),
					fielddecimal.New(properties.New(6, "decimal", true, testLog),
						fieldint32.NewDefaultOperationWithValue(properties.New(6, "decimalExponent", true, testLog), -1),
						fieldint64.NewDefaultOperationWithValue(properties.New(6, "decimalMantissa", true, testLog), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true, testLog),
						fieldint32.NewDefaultOperationWithValue(properties.New(7, "decimal with exp/manExponent", true, testLog), -2),
						fieldint64.NewDefaultOperationWithValue(properties.New(7, "decimal with exp/manMantissa", true, testLog), 2)),
					fieldunicodestring.NewDefaultOperationWithValue(properties.New(8, "StringUnicode", true, testLog), "Hello: ϔ"),
					fieldbytevector.NewDefaultOperationWithValue(properties.New(9, "byteVector", true, testLog), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true, testLog),
						fielduint32.NewDefaultOperationWithValue(properties.New(11, "length", true, testLog), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}

func TestCanLoadCopyOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_copy_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.NewCopyOperationWithInitialValue(properties.New(1, "String", true, testLog), "Hello"),
					fielduint32.NewCopyOperationWithInitialValue(properties.New(2, "unsigned int32", true, testLog), 10),
					fieldint32.NewCopyOperationWithInitialValue(properties.New(3, "signed int32", true, testLog), -10),
					fielduint64.NewCopyOperationWithInitialValue(properties.New(4, "unsigned int64", true, testLog), 10),
					fieldint64.NewCopyOperationWithInitialValue(properties.New(5, "signed int64", true, testLog), -10),
					fielddecimal.New(properties.New(6, "decimal", true, testLog),
						fieldint32.NewCopyOperationWithInitialValue(properties.New(6, "decimalExponent", true, testLog), -1),
						fieldint64.NewCopyOperationWithInitialValue(properties.New(6, "decimalMantissa", true, testLog), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true, testLog),
						fieldint32.NewCopyOperationWithInitialValue(properties.New(7, "decimal with exp/manExponent", true, testLog), -2),
						fieldint64.NewCopyOperationWithInitialValue(properties.New(7, "decimal with exp/manMantissa", true, testLog), 2)),
					fieldunicodestring.NewCopyOperationWithInitialValue(properties.New(8, "StringUnicode", true, testLog), "Hello: ϔ"),
					fieldbytevector.NewCopyOperationWithInitialValue(properties.New(9, "byteVector", true, testLog), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true, testLog),
						fielduint32.NewCopyOperationWithInitialValue(properties.New(11, "length", true, testLog), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}

func TestCanLoadIncrementOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_increment_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fielduint32.NewIncrementOperationWithInitialValue(properties.New(1, "unsigned int32", true, testLog), 10),
					fieldint32.NewIncrementOperationWithInitialValue(properties.New(2, "signed int32", true, testLog), -10),
					fielduint64.NewIncrementOperationWithInitialValue(properties.New(3, "unsigned int64", true, testLog), 10),
					fieldint64.NewIncrementOperationWithInitialValue(properties.New(4, "signed int64", true, testLog), -10),
					fieldsequence.New(properties.New(5, "sequence", true, testLog),
						fielduint32.NewIncrementOperationWithInitialValue(properties.New(6, "length", true, testLog), 2),
						[]store.Unit{
							fielduint32.NewIncrementOperation(properties.New(7, "sequence field 1", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}

func TestCanLoadTailOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_tail_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.NewTailOperationWithInitialValue(properties.New(1, "String", true, testLog), "Hello"),
					fieldunicodestring.NewTailOperationWithInitialValue(properties.New(2, "StringUnicode", true, testLog), "Hello: ϔ"),
					fieldbytevector.NewTailOperationWithInitialValue(properties.New(3, "byteVector", true, testLog), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}

func TestCanLoadDeltaOperationOnAllSupportedTypesFromTemplateFile(t *testing.T) {
	file, _ := os.Open("../../../../test/template-loader-tests/test_load_delta_operation_on_all_supported_types.xml")
	expectedStore := store.Store{
		Templates: map[uint32]store.Template{
			144: {
				Logger: testLog,
				TemplateUnits: []store.Unit{
					fieldasciistring.NewDeltaOperationWithInitialValue(properties.New(1, "String", true, testLog), "Hello"),
					fielduint32.NewDeltaOperationWithInitialValue(properties.New(2, "unsigned int32", true, testLog), 10),
					fieldint32.NewDeltaOperationWithInitialValue(properties.New(3, "signed int32", true, testLog), -10),
					fielduint64.NewDeltaOperationWithInitialValue(properties.New(4, "unsigned int64", true, testLog), 10),
					fieldint64.NewDeltaOperationWithInitialValue(properties.New(5, "signed int64", true, testLog), -10),
					fielddecimal.New(properties.New(6, "decimal", true, testLog),
						fieldint32.NewDeltaOperationWithInitialValue(properties.New(6, "decimalExponent", true, testLog), -1),
						fieldint64.NewDeltaOperationWithInitialValue(properties.New(6, "decimalMantissa", true, testLog), 57)),
					fielddecimal.New(properties.New(7, "decimal with exp/man", true, testLog),
						fieldint32.NewDeltaOperationWithInitialValue(properties.New(7, "decimal with exp/manExponent", true, testLog), -2),
						fieldint64.NewDeltaOperationWithInitialValue(properties.New(7, "decimal with exp/manMantissa", true, testLog), 2)),
					fieldunicodestring.NewDeltaOperationWithInitialValue(properties.New(8, "StringUnicode", true, testLog), "Hello: ϔ"),
					fieldbytevector.NewDeltaOperationWithInitialValue(properties.New(9, "byteVector", true, testLog), []byte{0x54, 0x45, 0x53, 0x54, 0x3F}),
					fieldsequence.New(properties.New(10, "sequence", true, testLog),
						fielduint32.NewDeltaOperationWithInitialValue(properties.New(11, "length", true, testLog), 2),
						[]store.Unit{
							fieldasciistring.New(properties.New(12, "sequence field 1", true, testLog)),
						}),
				},
			},
		},
	}

	// Act
	loadedStore, err := Load(file, testLog)

	// Assert
	if err != nil {
		t.Errorf("Got an error loading the template when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedStore, loadedStore)
	if !areEqual {
		t.Errorf("The returned store and expected store were not equal:\nexpected:\t%v\nactual:\t\t%v", expectedStore, loadedStore)
	}
}
