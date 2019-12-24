package templates

import (
	"encoding/xml"
	"os"
	"reflect"
	"testing"
)

func TestCanTokenizeHeartbeatTemplateFile(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../test/test_fast_templates.xml")
	decoder := xml.NewDecoder(file)

	expectedTokens := XMLTag{
		Type: "templates",
		Attributes: map[string]string{
			"xmlns": "http://www.fixprotocol.org/ns/fast/td/1.1",
		},
		NestedTags: []XMLTag{
			XMLTag{
				Type: "template",
				Attributes: map[string]string{
					"name":       "MDHeartbeat_144",
					"id":         "144",
					"dictionary": "144",
					"xmlns":      "http://www.fixprotocol.org/ns/fast/td/1.1",
				},
				NestedTags: []XMLTag{
					XMLTag{
						Type: "string",
						Attributes: map[string]string{
							"name": "ApplVerID",
							"id":   "1128",
						},
						NestedTags: []XMLTag{
							XMLTag{
								Type: "constant",
								Attributes: map[string]string{
									"value": "9",
								},
							},
						},
					},
					XMLTag{
						Type: "string",
						Attributes: map[string]string{
							"name": "MsgType",
							"id":   "35",
						},
						NestedTags: []XMLTag{
							XMLTag{
								Type: "constant",
								Attributes: map[string]string{
									"value": "0",
								},
							},
						},
					},
					XMLTag{
						Type: "uInt32",
						Attributes: map[string]string{
							"name": "MsgSeqNum",
							"id":   "34",
						},
					},
					XMLTag{
						Type: "uInt64",
						Attributes: map[string]string{
							"name": "SendingTime",
							"id":   "52",
						},
					},
				},
			},
		},
	}

	// Act
	tokens, err := LoadTagsFrom(decoder)

	// Assert
	if err != nil {
		t.Errorf("Got an error parsing the XML when none was expected: %s", err)
	}

	areEqual := reflect.DeepEqual(expectedTokens, *tokens)
	if !areEqual {
		t.Errorf("The returned tokens from parsing the XML did not equal the expected tokens:\nexpected:%s\nactual:%s", expectedTokens, tokens)
	}
}
