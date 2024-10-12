package nodes

import (
	"testing"
)

func TestServiceNode_NodeType(t *testing.T) {
	tests := []struct {
		name     string
		node     *ServiceNode
		expected string
	}{
		{
			name:     "Basic ServiceNode",
			node:     &ServiceNode{},
			expected: "Service",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.NodeType(); got != tt.expected {
				t.Errorf("NodeType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestServiceNode_Initialization(t *testing.T) {
	tests := []struct {
		name     string
		node     *ServiceNode
		expected *ServiceNode
	}{
		{
			name: "Full Initialization",
			node: &ServiceNode{
				Name:      "my-service",
				Namespace: "default",
				Ports: map[int]int{
					80:  8080,
					443: 8443,
				},
				Labels: map[string]string{
					"app": "my-app",
				},
			},
			expected: &ServiceNode{
				Name:      "my-service",
				Namespace: "default",
				Ports: map[int]int{
					80:  8080,
					443: 8443,
				},
				Labels: map[string]string{
					"app": "my-app",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.node.Name != tt.expected.Name {
				t.Errorf("expected Name to be '%s', got '%s'", tt.expected.Name, tt.node.Name)
			}
			if tt.node.Namespace != tt.expected.Namespace {
				t.Errorf("expected Namespace to be '%s', got '%s'", tt.expected.Namespace, tt.node.Namespace)
			}
			for port, targetPort := range tt.expected.Ports {
				if tt.node.Ports[port] != targetPort {
					t.Errorf("expected Ports[%d] to be %d, got %d", port, targetPort, tt.node.Ports[port])
				}
			}
			for key, value := range tt.expected.Labels {
				if tt.node.Labels[key] != value {
					t.Errorf("expected Labels['%s'] to be '%s', got '%s'", key, value, tt.node.Labels[key])
				}
			}
		})
	}
}
