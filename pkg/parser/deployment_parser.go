package parser

import (
	"fmt"
	"strings"

	"github.com/samsond/krypton/pkg/lexer"
	"github.com/samsond/krypton/pkg/nodes"
)

func (p *Parser) ParseDeployment() (*nodes.DeploymentNode, error) {
	node := &nodes.DeploymentNode{
		Env:   make(map[string]string),
		Ports: make(map[string]int),
	}

	// Expecting "deploy app <name> {"
	if p.token.Type != TokenDeployApp {
		return nil, fmt.Errorf("expected deploy app, got %s", p.token.Value)
	}
	// Use the token value for the app name.
	node.Name = p.token.Value

	p.token = p.lexer.NextToken()

	// Parsing the deployment details until we encounter "}"
	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		switch p.token.Type {
		case lexer.TokenNamespace:
			node.Namespace = p.token.Value
		case lexer.TokenReplicas:
			replicas, err := parseInt(p.token.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid replicas %s", p.token.Value)
			}
			node.Replicas = replicas
		case lexer.TokenImage:
			node.Image = p.token.Value
		case lexer.TokenEnv:
			env, err := p.parseEnv()
			if err != nil {
				return nil, err
			}
			node.Env = env
		case lexer.TokenPorts:
			ports, err := p.parsePorts()
			if err != nil {
				return nil, err
			}
			node.Ports = ports
		case lexer.TokenResources:
			resources, err := p.parseResources()
			if err != nil {
				return nil, err
			}
			node.Resources = resources
		case lexer.TokenStorage:
			storage, err := p.parseStorage()
			if err != nil {
				return nil, err
			}
			node.Storage = storage
		case lexer.TokenArgs:
			args, err := p.parseArgs()
			if err != nil {
				return nil, err
			}
			node.Args = args

		default:
			return nil, fmt.Errorf("unexpected token %s", p.token.Value)
		}
		p.token = p.lexer.NextToken()
	}

	return node, nil
}

func (p *Parser) parseResources() (*nodes.ResourceRequirementsNode, error) {
	resources := &nodes.ResourceRequirementsNode{}
	p.token = p.lexer.NextToken() // Skip "resources {"

	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		switch p.token.Type {
		case lexer.TokenLimits:
			limits, err := p.parseResourceSpec()
			if err != nil {
				return nil, err
			}
			resources.Limits = limits
		case lexer.TokenRequests:
			requests, err := p.parseResourceSpec()
			if err != nil {
				return nil, err
			}
			resources.Requests = requests
		default:
			return nil, fmt.Errorf("unexpected token in resources: %s", p.token.Value)
		}
		p.token = p.lexer.NextToken()
	}
	return resources, nil
}

func (p *Parser) parseResourceSpec() (nodes.ResourceSpec, error) {
	spec := nodes.ResourceSpec{}
	p.token = p.lexer.NextToken() // Skip "limits {" or "requests {"

	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		switch p.token.Type {
		case lexer.TokenMemory:
			spec.Memory = strings.Trim(p.token.Value, `";`)
		case lexer.TokenCPU:
			spec.CPU = strings.Trim(p.token.Value, `";`)
		}
		p.token = p.lexer.NextToken()
	}
	return spec, nil
}

func (p *Parser) parseStorage() (*nodes.StorageConfigNode, error) {
	storage := &nodes.StorageConfigNode{}
	p.token = p.lexer.NextToken() // Skip the "storage {"

	for p.token.Type != lexer.TokenRBrace && p.token.Type != TokenEOF {
		switch p.token.Type {
		case lexer.TokenVolume:
			storage.Volume = p.token.Value
		case lexer.TokenSize:
			storage.Size = p.token.Value
		default:
			return nil, fmt.Errorf("unexpected token in storage: %s", p.token.Value)
		}
		p.token = p.lexer.NextToken()
	}
	return storage, nil
}
