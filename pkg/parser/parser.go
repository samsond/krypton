package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samsond/krypton/pkg/lexer"
	"github.com/samsond/krypton/pkg/nodes"
)

// Type aliases for lexer types
type TokenType = lexer.TokenType
type Token = lexer.Token

// Constant aliases for lexer constants
const (
	TokenEOF       = lexer.TokenEOF
	TokenDeployApp = lexer.TokenDeployApp
	TokenService   = lexer.TokenService
	TokenSeparator = lexer.TokenSeparator
)

type Parser struct {
	lexer *lexer.Lexer
	token lexer.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer, token: lexer.NextToken()}
}

func (p *Parser) ParseResources() ([]nodes.Node, error) {
	var resources []nodes.Node

	for p.token.Type != TokenEOF {
		var resource nodes.Node
		var err error

		switch p.token.Type {
		case TokenDeployApp:
			resource, err = p.ParseDeployment()
		case TokenService:
			resource, err = p.ParseService()
		case TokenSeparator:
			p.token = p.lexer.NextToken() // Move to the next resource block.
			continue
		default:
			return nil, fmt.Errorf("unexpected token '%s' at line %d", p.token.Value, p.lexer.LineNumber)
		}

		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
		p.token = p.lexer.NextToken() // Move to the next token for further parsing.
	}

	return resources, nil
}

// parseKeyValuePairs parses key-value pairs from the input.
func (p *Parser) parseKeyValuePairs() (map[string]string, error) {
	pairs := make(map[string]string)
	p.token = p.lexer.NextToken() // Skip the opening brace

	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		keyValue := strings.SplitN(p.token.Value, ":", 2)
		if len(keyValue) != 2 {
			return nil, fmt.Errorf("invalid format: %s", p.token.Value)
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.Trim(strings.TrimSpace(keyValue[1]), `";`)
		pairs[key] = value
		p.token = p.lexer.NextToken()
	}
	return pairs, nil
}

// parseEnv parses environment variables.
func (p *Parser) parseEnv() (map[string]string, error) {
	return p.parseKeyValuePairs()
}

// parsePorts parses port mappings.
func (p *Parser) parsePorts() (map[string]int, error) {
	stringPairs, err := p.parseKeyValuePairs()
	if err != nil {
		return nil, err
	}

	ports := make(map[string]int)
	for key, value := range stringPairs {
		port, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("invalid port value: %s", value)
		}
		ports[key] = port
	}
	return ports, nil
}

func (p *Parser) parseArgs() ([]string, error) {
	var args []string
	p.token = p.lexer.NextToken() // Skip the "args {" token

	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		// Assuming each argument is represented as a quoted string
		arg := strings.Trim(p.token.Value, `"`)
		args = append(args, arg)
		p.token = p.lexer.NextToken()
	}

	return args, nil
}

// parseInt converts a string to an integer
func parseInt(value string) (int, error) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return n, nil
}
