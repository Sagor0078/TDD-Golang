package ReadingFiles

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	// Load input from file
	inputPath := filepath.Join("testdata", "input.txt")
	input, err := os.ReadFile(inputPath)
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	// Load expected output
	expectedPath := filepath.Join("testdata", "expected_output.txt")
	expected, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("failed to read expected output: %v", err)
	}

	got := Process(input)

	gotStr := strings.TrimSpace(string(got))
	expectedStr := strings.TrimSpace(string(expected))

	if gotStr != expectedStr {
		t.Errorf("unexpected output:\nGot: %q\nWant: %q", gotStr, expectedStr)
	}
}
