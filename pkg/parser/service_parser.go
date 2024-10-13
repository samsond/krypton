package parser

import (
	"fmt"
	"strings"

	"github.com/samsond/krypton/pkg/lexer"
	"github.com/samsond/krypton/pkg/nodes"
)

func (p *Parser) ParseService() (*nodes.ServiceNode, error) {
	node := &nodes.ServiceNode{
		Ports:  make(map[int]int),
		Labels: make(map[string]string),
	}

	if p.token.Type != TokenService {
		return nil, fmt.Errorf("expected 'service', got '%s'", p.token.Value)
	}

	node.Name = strings.TrimSpace(p.token.Value)
	p.token = p.lexer.NextToken()

	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		switch p.token.Type {
		case lexer.TokenNamespace:
			node.Namespace = p.token.Value
		case lexer.TokenPort:
			port, err := parseInt(p.token.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid port %s", p.token.Value)
			}
			p.token = p.lexer.NextToken()
			if p.token.Type == lexer.TokenTargetPort {
				targetPort, err := parseInt(p.token.Value)
				if err != nil {
					return nil, fmt.Errorf("invalid targetPort %s", p.token.Value)
				}
				node.Ports[port] = targetPort
			}
		default:
			return nil, fmt.Errorf("unexpected token '%s' at line %d", p.token.Value, p.lexer.LineNumber)
		}
		p.token = p.lexer.NextToken()
	}

	return node, nil
}
