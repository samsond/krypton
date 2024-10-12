package lexer

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanToken(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   Token
		lineNumber int
	}{
		{"Separator", "---", Token{Type: TokenSeparator, Value: separatorValue, LineNumber: 1}, 1},
		{"DeployApp", "deploy app myapp {", Token{Type: TokenDeployApp, Value: "myapp ", LineNumber: 2}, 2},
		{"Namespace", "namespace: mynamespace", Token{Type: TokenNamespace, Value: "mynamespace", LineNumber: 3}, 3},
		{"Replicas", "replicas: 3", Token{Type: TokenReplicas, Value: "3", LineNumber: 4}, 4},
		{"Image", "image: myimage", Token{Type: TokenImage, Value: "myimage", LineNumber: 5}, 5},
		{"Env", "env {", Token{Type: TokenEnv, Value: envTokenValue, LineNumber: 6}, 6},
		{"Ports", "ports {", Token{Type: TokenPorts, Value: portsTokenValue, LineNumber: 7}, 7},
		{"Resources", "resources {", Token{Type: TokenResources, Value: resourcesTokenValue, LineNumber: 8}, 8},
		{"Limits", "limits {", Token{Type: TokenLimits, Value: limitsTokenValue, LineNumber: 9}, 9},
		{"Requests", "requests {", Token{Type: TokenRequests, Value: requestsTokenValue, LineNumber: 10}, 10},
		{"Memory", "memory: 512Mi", Token{Type: TokenMemory, Value: "512Mi", LineNumber: 11}, 11},
		{"CPU", "cpu: 1", Token{Type: TokenCPU, Value: "1", LineNumber: 12}, 12},
		{"Storage", "storage {", Token{Type: TokenStorage, Value: storageTokenValue, LineNumber: 13}, 13},
		{"Volume", "volume: myvolume", Token{Type: TokenVolume, Value: "myvolume", LineNumber: 14}, 14},
		{"Size", "size: 10Gi", Token{Type: TokenSize, Value: "10Gi", LineNumber: 15}, 15},
		{"RightBrace", "}", Token{Type: TokenRBrace, Value: rightbraceValue, LineNumber: 16}, 16},
		{"Service", "service myservice {", Token{Type: TokenService, Value: "myservice", LineNumber: 17}, 17},
		{"Port", "port: 8080", Token{Type: TokenPort, Value: "8080", LineNumber: 18}, 18},
		{"TargetPort", "targetPort: 80", Token{Type: TokenTargetPort, Value: "80", LineNumber: 19}, 19},
		{"Identifier", "someIdentifier", Token{Type: TokenIdentifier, Value: "someIdentifier", LineNumber: 20}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := &Lexer{LineNumber: tt.lineNumber}
			token := lexer.scanToken(tt.input)
			require.Equal(t, tt.expected, token)
		})
	}
}

func TestParseStringValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Namespace with colon", `namespace: "default";`, "default"},
		{"Replicas with colon", `replicas: 3;`, "3"},
		{"Image with colon", `image: "myimage";`, "myimage"},
		{"Memory with colon", `memory: 512Mi;`, "512Mi"},
		{"CPU with colon", `cpu: 1;`, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringValue(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestLexer(t *testing.T) {
	// Read input from testdata/dsl_input.txt
	data, err := os.ReadFile("testdata/dsl_input.txt")
	require.NoError(t, err)

	input := strings.TrimSpace(string(data))

	expectedTokens := []Token{
		{Type: 1, Value: "my-app ", LineNumber: 1},
		{Type: 2, Value: "default", LineNumber: 2},
		{Type: 3, Value: "3", LineNumber: 3},
		{Type: 4, Value: "my-image:latest", LineNumber: 4},
		{Type: 5, Value: "env", LineNumber: 5},
		{Type: 14, Value: `ENV_VAR: "value";`, LineNumber: 6},
		{Type: 10, Value: "}", LineNumber: 7},
		{Type: 6, Value: "ports", LineNumber: 8},
		{Type: 14, Value: `http: 8080;`, LineNumber: 9},
		{Type: 10, Value: "}", LineNumber: 10},
		{Type: 7, Value: "resources", LineNumber: 11},
		{Type: 19, Value: "limits", LineNumber: 12},
		{Type: 15, Value: "512Mi", LineNumber: 13},
		{Type: 16, Value: "0.5", LineNumber: 14},
		{Type: 10, Value: "}", LineNumber: 15},
		{Type: 20, Value: "requests", LineNumber: 16},
		{Type: 15, Value: "256Mi", LineNumber: 17},
		{Type: 16, Value: "0.25", LineNumber: 18},
		{Type: 10, Value: "}", LineNumber: 19},
		{Type: 10, Value: "}", LineNumber: 20},
		{Type: 8, Value: "storage", LineNumber: 21},
		{Type: 17, Value: "my-volume", LineNumber: 22},
		{Type: 18, Value: "10Gi", LineNumber: 23},
		{Type: 10, Value: "}", LineNumber: 24},
		{Type: 10, Value: "}", LineNumber: 25},
		{Type: 26, Value: "---", LineNumber: 26},
		{Type: 22, Value: "my-service", LineNumber: 27},
		{Type: 2, Value: "default", LineNumber: 28},
		{Type: 23, Value: "80", LineNumber: 29},
		{Type: 24, Value: "8080", LineNumber: 30},
		{Type: 10, Value: "}", LineNumber: 31},
		{Type: 0, Value: "", LineNumber: 32},
	}

	// Initialize the Lexer using NewLexer
	lexer := NewLexer(input)

	// Verify the tokens
	require.Equal(t, expectedTokens, lexer.tokens)

	// Additional test for directly calling tokenize method
	scanner := bufio.NewScanner(strings.NewReader(input))
	lexer = &Lexer{scanner: scanner, LineNumber: 1}
	lexer.tokenize()

	require.Equal(t, expectedTokens, lexer.tokens)
}
