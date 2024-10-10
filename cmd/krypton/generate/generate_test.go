package generate

import (
	"os"
	"strings"
	"testing"
)

func TestNewGenerateCommand(t *testing.T) {

	// Set the working directory
	err := os.Chdir("../../../")
	if err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	// Initialize the command
	cmd := NewGenerateCommand()

	// temp file
	tmpFile, err := os.CreateTemp("", "output.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the file afterwards

	// Set the command arguments, including the output file path
	cmd.SetArgs([]string{"examples/basic_app.kp", "-o", tmpFile.Name()})

	// Execute the command
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Read the content of the temporary file
	actualOutput, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	expectedYAML, err := os.ReadFile("examples/deployment.yaml")

	if err != nil {
		t.Fatalf("failed to read expected YAML file: %v", err)
	}

	trimmedExpectedOutput := strings.TrimSpace(string(expectedYAML))
	trimmedActualOutput := strings.TrimSpace(string(actualOutput))

	if trimmedActualOutput != trimmedExpectedOutput {
		t.Errorf("expected output:\n%s\nbut got:\n%s", trimmedExpectedOutput, trimmedActualOutput)
	}
}
