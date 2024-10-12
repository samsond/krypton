package parser

import (
	"strings"
	"testing"

	"github.com/samsond/krypton/pkg/lexer"
	"github.com/stretchr/testify/require"
)

func TestParseService(t *testing.T) {
	input := `
        service my-service {
            namespace: "default";
            port: 80;
            targetPort: 8080;
        }`

	// Create a new lexer from the input
	lex := lexer.NewLexer(strings.TrimSpace(input))

	// Initialize the parser with the lexer
	p := NewParser(lex)

	// Call ParseService to parse the input
	serviceNode, err := p.ParseService()

	// Ensure no errors were encountered
	require.NoError(t, err)

	// Verify the parsed service node fields
	require.NotNil(t, serviceNode)
	require.Equal(t, "my-service", serviceNode.Name)
	require.Equal(t, "default", serviceNode.Namespace)
	require.Equal(t, 80, serviceNode.Ports[80])
	require.Equal(t, 8080, serviceNode.Ports[80])

	// Test edge cases, like missing or invalid ports
	t.Run("Test invalid port", func(t *testing.T) {
		invalidInput := `
        service my-service {
            namespace: "default";
            port: abc;  // Invalid port
            targetPort: 8080;
        }`
		lexInvalid := lexer.NewLexer(strings.TrimSpace(invalidInput))
		pInvalid := NewParser(lexInvalid)
		_, err := pInvalid.ParseService()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid port abc")
	})

	t.Run("Test missing targetPort", func(t *testing.T) {
		missingTargetPortInput := `
        service my-service {
            namespace: "default";
            port: 80;
        }`
		lexMissingTargetPort := lexer.NewLexer(strings.TrimSpace(missingTargetPortInput))
		pMissingTargetPort := NewParser(lexMissingTargetPort)
		_, err := pMissingTargetPort.ParseService()
		require.Error(t, err)
		require.Contains(t, err.Error(), "missing targetPort for port 80")
	})
}
