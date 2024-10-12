package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeploymentLiterals(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"deployAppPrefix", deployAppPrefix, "deploy app"},
		{"namespacePrefix", namespacePrefix, "namespace:"},
		{"replicasPrefix", replicasPrefix, "replicas:"},
		{"imagePrefix", imagePrefix, "image:"},
		{"envPrefix", envPrefix, "env {"},
		{"envTokenValue", envTokenValue, "env"},
		{"resourcesPrefix", resourcesPrefix, "resources {"},
		{"resourcesTokenValue", resourcesTokenValue, "resources"},
		{"limitsPrefix", limitsPrefix, "limits {"},
		{"limitsTokenValue", limitsTokenValue, "limits"},
		{"requestsPrefix", requestsPrefix, "requests {"},
		{"memoryPrefix", memoryPrefix, "memory:"},
		{"cpuPrefix", cpuPrefix, "cpu:"},
		{"storagePrefix", storagePrefix, "storage {"},
		{"storageTokenValue", storageTokenValue, "storage"},
		{"volumePrefix", volumePrefix, "volume:"},
		{"sizePrefix", sizePrefix, "size:"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.actual)
		})
	}
}
