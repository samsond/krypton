package nodes

import (
	"bytes"
	"fmt"
	"text/template"
)

type DeploymentNode struct {
	Name      string
	Namespace string
	Replicas  int
	Image     string
	Args      []string
	Env       map[string]string
	Ports     map[string]int
	Resources *ResourceRequirementsNode
	Storage   *StorageConfigNode
}

type ResourceRequirementsNode struct {
	Limits   ResourceSpec
	Requests ResourceSpec
}

type StorageConfigNode struct {
	Volume string
	Size   string
}

type ResourceSpec struct {
	Memory string
	CPU    string
}

func (n *DeploymentNode) NodeType() string {
	return "Deployment"
}

func (n *ResourceRequirementsNode) NodeType() string {
	return "ResourceRequirements"
}

func (n *StorageConfigNode) NodeType() string {
	return "StorageConfig"
}

// generateTemplate is a common function to execute a template with the given data.
func generateTemplate(templatePath string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var manifests bytes.Buffer
	err = tmpl.Execute(&manifests, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return manifests.String(), nil
}
