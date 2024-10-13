package generator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/samsond/krypton/pkg/nodes"
)

func GenerateYAML(resource nodes.Node) (string, error) {
	if resource.NodeType() == "" {
		return "", errors.New("node type is empty")
		
	}
	nodeType := strings.ToLower(resource.NodeType())
	tmplPath, err := getTemplatePath(nodeType)
	if err != nil {
		return "", err
	}
	return generateTemplate(tmplPath, resource)
}

func getTemplatePath(nodeType string) (string, error) {

	templatePath := filepath.Join("pkg", "templates", nodeType + ".tmpl")
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return "", errors.New("template not found for nodeType: " + nodeType)
	}
	return templatePath, nil
}

// generateTemplate is a common function to execute a template with the given data.
func generateTemplate(tmplPath string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(tmplPath)
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
