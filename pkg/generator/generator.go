package generator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/samsond/krypton/pkg/parser"
	"github.com/samsond/krypton/pkg/templates"
)

// GenerateFromApp generates Kubernetes YAML manifests from the parsed AppDeployment.
func GenerateFromApp(app *parser.AppDeployment) (string, error) {
	tmplPath := templates.GetDeploymentTemplatePath()
	return generateTemplate(tmplPath, app)
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
