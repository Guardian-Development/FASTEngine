package engine

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"log"
	"os"
	"testing"
)

func TestCanReadInstrumentMessages(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../test/example-decoding-tests/instrument-messages-hex.txt")
	defer file.Close()

	logger := log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile)
	engine, err := NewFromTemplateFile("../../test/example-decoding-tests/templates.xml", logger)
	if err != nil {
		t.Fatalf("unable to load engine: %v", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		message, _ := hex.DecodeString(scanner.Text())

		// Act
		fix, err := engine.Deserialise(bytes.NewBuffer(message))

		// Assert
		if err != nil {
			t.Fatalf("unable to decode message: %s", message)
		}

		logger.Printf("message decoded: %s", fix.String())
	}
}

func TestCanReadSnapshotMessagesTest(t *testing.T) {
	// Arrange
	file, _ := os.Open("../../test/example-decoding-tests/snapshot-messages-hex.txt")
	defer file.Close()

	logger := log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile)
	engine, err := NewFromTemplateFile("../../test/example-decoding-tests/templates.xml", logger)
	if err != nil {
		t.Fatalf("unable to load engine: %v", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		message, _ := hex.DecodeString(scanner.Text())

		// Act
		fix, err := engine.Deserialise(bytes.NewBuffer(message))

		// Assert
		if err != nil {
			t.Fatalf("unable to decode message: %s", message)
		}

		logger.Printf("message decoded: %s", fix.String())
	}
}
