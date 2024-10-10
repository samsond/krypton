package parser

import (
	"bufio"
	"os"
	"testing"
)

// MockResource is a mock implementation of the Resource interface for testing purposes.
type MockResource struct{}

func (m *MockResource) SomeMethod() {}

// Implement the GetName method to satisfy the Resource interface.
func (m *MockResource) GetName() string {
	return "MockResource"
}

// Implement the GetNamespace method to satisfy the Resource interface.
func (m *MockResource) GetNamespace() string {
	return "default"
}

// Mock parse function for testing.
func mockParseFunc(line string, scanner *bufio.Scanner) (Resource, error) {
	return &MockResource{}, nil
}

func TestParseDSL(t *testing.T) {
	// Read the DSL content from examples/basic_app.kp
	dslFilePath := "../../examples/basic_app.kp"
	_, err := os.ReadFile(dslFilePath)
	if err != nil {
		t.Fatalf("failed to read DSL file: %v", err)
	}

	// Override the resourceParsers map for testing.
	originalParsers := resourceParsers
	resourceParsers = map[string]parseFunc{
		"deploy app": mockParseFunc,
	}
	defer func() { resourceParsers = originalParsers }()

	// Test the ParseDSL function.
	resource, err := ParseDSL(dslFilePath)
	if err != nil {
		t.Fatalf("ParseDSL() error = %v; expected no error", err)
	}
	if _, ok := resource.(*MockResource); !ok {
		t.Fatalf("ParseDSL() returned unexpected resource type")
	}
}

func TestParseDSLError(t *testing.T) {
	// Test with a non-existent file.
	_, err := ParseDSL("non_existent_file.kp")
	if err == nil {
		t.Fatalf("ParseDSL() error = nil; expected an error")
	}
}
