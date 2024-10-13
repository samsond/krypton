package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonLiterals(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"separatorValue", separatorValue, "---"},
		{"rightbraceValue", rightbraceValue, "}"},
		{"leftbraceValue", leftbraceValue, "{"},
		{"portsPrefix", portsPrefix, "ports {"},
		{"portsTokenValue", portsTokenValue, "ports"},
		{"portPrefix", portPrefix, "port:"},
		{"targetPortPrefix", targetPortPrefix, "targetPort:"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.actual)
		})
	}
}
