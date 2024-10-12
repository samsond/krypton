package parser

import (
	"testing"

	"github.com/samsond/krypton/pkg/lexer"
	"github.com/stretchr/testify/require"
)

type MockLexer struct {
	*lexer.Lexer
	token lexer.Token
}

func (m *MockLexer) NextToken() lexer.Token {
	return m.token
}

func TestNewParser(t *testing.T) {
	mockToken := lexer.Token{
		Type:  lexer.TokenIdentifier,
		Value: "mock",
	}

	// Initialize the real lexer with an empty input
	realLexer := lexer.NewLexer("")

	// Create the mock lexer embedding the real lexer
	mockLexer := &MockLexer{
		Lexer: realLexer,
		token: mockToken,
	}

	parser := NewParser(mockLexer)

	require.NotNil(t, parser)
	require.Equal(t, mockLexer, parser.lexer)
	require.Equal(t, mockToken, parser.token)
}
