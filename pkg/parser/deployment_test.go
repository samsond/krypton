package parser

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// MockAppDeployment is a mock implementation of the AppDeployment struct for testing purposes.
type MockAppDeployment struct {
	Name      string
	Namespace string
	Ports     map[string]int
	Env       map[string]string
}

// TestParseAppDeployment tests the parseAppDeployment function.
func TestParseAppDeployment(t *testing.T) {
	// Read the DSL content from examples/basic_app.kp.
	dslFilePath := "../../examples/basic_app.kp"
	dslContent, err := os.ReadFile(dslFilePath)
	if err != nil {
		t.Fatalf("failed to read DSL file: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(dslContent)))
	scanner.Scan() // Skip the first empty line

	resource, err := parseAppDeployment(scanner.Text(), scanner)
	if err != nil {
		t.Fatalf("parseAppDeployment() error = %v; expected no error", err)
	}

	app, ok := resource.(*AppDeployment)
	if !ok {
		t.Fatalf("parseAppDeployment() returned unexpected resource type")
	}

	// Verify the parsed values
	if app.Name != "my-app" {
		t.Errorf("expected Name to be 'my-app', got '%s'", app.Name)
	}
	if app.Namespace != "default" {
		t.Errorf("expected Namespace to be 'default', got '%s'", app.Namespace)
	}
	if app.Ports["http"] != 8080 {
		t.Errorf("expected http port to be 8080, got %d", app.Ports["http"])
	}
	if app.Ports["metrics"] != 2112 {
		t.Errorf("expected metrics port to be 2112, got %d", app.Ports["metrics"])
	}
	if app.Env["DATABASE_URL"] != "postgres://user:password@host/db" {
		t.Errorf("expected DATABASE_URL to be 'postgres://user:password@host/db', got '%s'", app.Env["DATABASE_URL"])
	}
}
