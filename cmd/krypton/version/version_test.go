package version

import (
	"bytes"
	"fmt"
	"testing"
)

func TestVersionString(t *testing.T) {
	// Define test cases with different rawVersion and stages
	tests := []struct {
		rawVersion string
		stage      string
		expected   string
	}{
		{"v1.0.0", "dev", "v1.0.0-beta"},
		{"v1.0.0", "prod", "v1.0.0"},
	}

	for _, test := range tests {
		setVersionAndStage(test.rawVersion, test.stage)

		// Compare the output
		result := versionString()
		if result != test.expected {
			t.Errorf("versionString() = %s; expected %s", result, test.expected)
		}
	}
}

func TestNewVersionCommand(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create the command and set its output to the buffer
	cmd := NewVersionCommand()
	cmd.SetOut(&buf)

	// Execute the command
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Get the output from the buffer
	output := buf.String()
	expected := fmt.Sprintf("kptn version: %s\n", versionString())

	// Compare the output with the expected value
	if output != expected {
		t.Errorf("NewVersionCommand output = %q; expected %q", output, expected)
	}
}
