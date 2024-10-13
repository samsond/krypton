package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenTypes(t *testing.T) {
	tests := []struct {
		name     string
		token    TokenType
		expected int
	}{
		{"TokenEOF", TokenEOF, 0},
		{"TokenDeployApp", TokenDeployApp, 1},
		{"TokenNamespace", TokenNamespace, 2},
		{"TokenReplicas", TokenReplicas, 3},
		{"TokenImage", TokenImage, 4},
		{"TokenEnv", TokenEnv, 5},
		{"TokenPorts", TokenPorts, 6},
		{"TokenResources", TokenResources, 7},
		{"TokenStorage", TokenStorage, 8},
		{"TokenLBrace", TokenLBrace, 9},
		{"TokenRBrace", TokenRBrace, 10},
		{"TokenColon", TokenColon, 11},
		{"TokenString", TokenString, 12},
		{"TokenNumber", TokenNumber, 13},
		{"TokenIdentifier", TokenIdentifier, 14},
		{"TokenMemory", TokenMemory, 15},
		{"TokenCPU", TokenCPU, 16},
		{"TokenVolume", TokenVolume, 17},
		{"TokenSize", TokenSize, 18},
		{"TokenLimits", TokenLimits, 19},
		{"TokenRequests", TokenRequests, 20},
		{"TokenArgs", TokenArgs, 21},
		{"TokenService", TokenService, 22},
		{"TokenPort", TokenPort, 23},
		{"TokenTargetPort", TokenTargetPort, 24},
		{"TokenTypeString", TokenTypeString, 25},
		{"TokenSeparator", TokenSeparator, 26},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, int(tt.token))
		})
	}
}
