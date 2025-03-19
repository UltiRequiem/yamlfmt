package yamlfmt

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestFormatStreamSuccess(t *testing.T) {
	// Check if it works well with correct YAML.
	input := "key: value\n"
	r := strings.NewReader(input)
	var out bytes.Buffer

	err := FormatStream(r, &out, 2)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(out.Bytes(), &data); err != nil {
		t.Fatalf("Output is not valid YAML: %v", err)
	}
}

func TestFormatStreamInvalid(t *testing.T) {
	// Check if the function returns an error when the YAML is incorrect.
	input := "key: [unclosed"
	r := strings.NewReader(input)
	var out bytes.Buffer

	err := FormatStream(r, &out, 2)
	if err == nil {
		t.Fatalf("Expected an error for invalid YAML, got nil")
	}
}

func TestDumpStreamOverwriteTrue(t *testing.T) {
	// Write content to a temporary file and checks that the file contains the expected content.
	content := "test content"
	buf := bytes.NewBufferString(content)

	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filename := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(filename)

	err = DumpStream(buf, filename, true)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if string(data) != content {
		t.Fatalf("Expected file content %q, got %q", content, string(data))
	}
}

func TestDumpStreamOverwriteFalse(t *testing.T) {
	// Capture the standard output to check that DumpStream writes to stdout.
	content := "stdout test"
	buf := bytes.NewBufferString(content)

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	err = DumpStream(buf, "", false)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	w.Close()
	var output bytes.Buffer
	io.Copy(&output, r)
	os.Stdout = oldStdout

	if output.String() != content {
		t.Fatalf("Expected stdout content %q, got %q", content, output.String())
	}
}

func TestFormatFileSuccess(t *testing.T) {
	// Create a temporary file, writes valid YAML into it, runs FormatFile,
	// and then verifies that the file was correctly formatted.
	content := "key: value\n"
	tmpFile, err := os.CreateTemp("", "formatfile_test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filename := tmpFile.Name()
	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	defer os.Remove(filename)

	ch := make(chan string, 1)
	FormatFile(filename, 2, true, ch)

	msg := <-ch
	expectedMsg := filename + " formatted!"
	if msg != expectedMsg {
		t.Fatalf("Expected message %q, got %q", expectedMsg, msg)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var parsed map[string]interface{}
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("Formatted file is not valid YAML: %v", err)
	}
}
