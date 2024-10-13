package nodes

import (
	"testing"
)

func TestDeploymentNode_NodeType(t *testing.T) {
	node := &DeploymentNode{}
	expected := "Deployment"
	if node.NodeType() != expected {
		t.Errorf("expected %s, got %s", expected, node.NodeType())
	}
}

func TestDeploymentNode_Initialization(t *testing.T) {
	node := &DeploymentNode{
		Name:      "my-app",
		Namespace: "default",
		Replicas:  3,
		Image:     "my-app:v1.0",
		Args:      []string{"--arg1", "--arg2"},
		Env: map[string]string{
			"DATABASE_URL": "postgres://user:password@host/db",
		},
		Ports: map[string]int{
			"http":    8080,
			"metrics": 2112,
		},
		Resources: &ResourceRequirementsNode{
			Limits: ResourceSpec{
				Memory: "512Mi",
				CPU:    "500m",
			},
			Requests: ResourceSpec{
				Memory: "256Mi",
				CPU:    "250m",
			},
		},
		Storage: &StorageConfigNode{
			Volume: "my-app-data",
			Size:   "1Gi",
		},
	}

	if node.Name != "my-app" {
		t.Errorf("expected Name to be 'my-app', got %s", node.Name)
	}
	if node.Namespace != "default" {
		t.Errorf("expected Namespace to be 'default', got %s", node.Namespace)
	}
	if node.Replicas != 3 {
		t.Errorf("expected Replicas to be 3, got %d", node.Replicas)
	}
	if node.Image != "my-app:v1.0" {
		t.Errorf("expected Image to be 'my-app:v1.0', got %s", node.Image)
	}
	if len(node.Args) != 2 || node.Args[0] != "--arg1" || node.Args[1] != "--arg2" {
		t.Errorf("expected Args to be ['--arg1', '--arg2'], got %v", node.Args)
	}
	if node.Env["DATABASE_URL"] != "postgres://user:password@host/db" {
		t.Errorf("expected Env['DATABASE_URL'] to be 'postgres://user:password@host/db', got %s", node.Env["DATABASE_URL"])
	}
	if node.Ports["http"] != 8080 {
		t.Errorf("expected Ports['http'] to be 8080, got %d", node.Ports["http"])
	}
	if node.Ports["metrics"] != 2112 {
		t.Errorf("expected Ports['metrics'] to be 2112, got %d", node.Ports["metrics"])
	}
	if node.Resources.Limits.Memory != "512Mi" {
		t.Errorf("expected Resources.Limits.Memory to be '512Mi', got %s", node.Resources.Limits.Memory)
	}
	if node.Resources.Limits.CPU != "500m" {
		t.Errorf("expected Resources.Limits.CPU to be '500m', got %s", node.Resources.Limits.CPU)
	}
	if node.Resources.Requests.Memory != "256Mi" {
		t.Errorf("expected Resources.Requests.Memory to be '256Mi', got %s", node.Resources.Requests.Memory)
	}
	if node.Resources.Requests.CPU != "250m" {
		t.Errorf("expected Resources.Requests.CPU to be '250m', got %s", node.Resources.Requests.CPU)
	}
	if node.Storage.Volume != "my-app-data" {
		t.Errorf("expected Storage.Volume to be 'my-app-data', got %s", node.Storage.Volume)
	}
	if node.Storage.Size != "1Gi" {
		t.Errorf("expected Storage.Size to be '1Gi', got %s", node.Storage.Size)
	}
}
