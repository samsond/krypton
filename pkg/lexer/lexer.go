package lexer

import (
	"bufio"
	"strings"
)

// LexerInterface defines the contract for any lexer implementation
type LexerInterface interface {
	NextToken() Token
}

// Token represents a token with its type, value, and line number.
type Token struct {
	Type       TokenType
	Value      string
	LineNumber int
}

// Lexer struct represents a lexer for the DSL.
type Lexer struct {
	scanner    *bufio.Scanner
	tokens     []Token // Store tokens for debugging
	pos        int     // Position for tracking the current token in tokens slice
	LineNumber int     // Current line number
}

// NewLexer initializes a new Lexer and pre-tokenizes the input for debugging purposes.
func NewLexer(input string) *Lexer {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	lexer := &Lexer{
		scanner:    scanner,
		LineNumber: 1,
	}
	lexer.tokenize() // Pre-tokenize to store tokens without consuming.
	return lexer
}

// tokenize reads all tokens without affecting the scanning process.
func (l *Lexer) tokenize() {
	for l.scanner.Scan() {
		text := strings.TrimSpace(l.scanner.Text())
		token := l.scanToken(text)
		l.tokens = append(l.tokens, token)
		l.LineNumber++
	}
	l.tokens = append(l.tokens, Token{Type: TokenEOF, LineNumber: l.LineNumber}) // Ensure EOF is included.
}

// NextToken returns the next token in sequence.
func (l *Lexer) NextToken() Token {
	if l.pos < len(l.tokens) {
		token := l.tokens[l.pos]
		l.pos++
		return token
	}
	return Token{Type: TokenEOF}
}

// scanToken processes a single line of text into a token.
func (l *Lexer) scanToken(text string) Token {
	switch {
	case text == separatorValue:
		return Token{Type: TokenSeparator, Value: separatorValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, deployAppPrefix):
		value := strings.TrimSpace(strings.TrimPrefix(text, deployAppPrefix))
		value = strings.TrimSuffix(value, leftbraceValue)
		return Token{Type: TokenDeployApp, Value: value, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, namespacePrefix):
		return Token{Type: TokenNamespace, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, replicasPrefix):
		return Token{Type: TokenReplicas, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, imagePrefix):
		return Token{Type: TokenImage, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, envPrefix):
		return Token{Type: TokenEnv, Value: envTokenValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, portsPrefix):
		return Token{Type: TokenPorts, Value: portsTokenValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, resourcesPrefix):
		return Token{Type: TokenResources, Value: resourcesTokenValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, limitsPrefix):
		return Token{Type: TokenLimits, Value: limitsTokenValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, requestsPrefix):
		return Token{Type: TokenRequests, Value: requestsTokenValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, memoryPrefix):
		return Token{Type: TokenMemory, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, cpuPrefix):
		return Token{Type: TokenCPU, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, storagePrefix):
		return Token{Type: TokenStorage, Value: storageTokenValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, volumePrefix):
		return Token{Type: TokenVolume, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, sizePrefix):
		return Token{Type: TokenSize, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case text == rightbraceValue:
		return Token{Type: TokenRBrace, Value: rightbraceValue, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, servicePrefix):
		value := strings.TrimSpace(strings.TrimPrefix(text, serviceTokenValue))
		value = strings.TrimSuffix(value, leftbraceValue)
		value = strings.TrimSpace(value)
		return Token{Type: TokenService, Value: value, LineNumber: l.LineNumber}
	case strings.HasPrefix(text, portPrefix):
		return Token{Type: TokenPort, Value: parseStringValue(text), LineNumber: l.LineNumber}
	case strings.HasPrefix(text, targetPortPrefix):
		return Token{Type: TokenTargetPort, Value: parseStringValue(text), LineNumber: l.LineNumber}
	default:
		return Token{Type: TokenIdentifier, Value: text, LineNumber: l.LineNumber}
	}
}

// parseStringValue extracts and cleans up the string value from a line.
func parseStringValue(line string) string {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return ""
	}
	value := strings.TrimSpace(parts[1])
	value = strings.Trim(value, `";{}`)
	return value
}
