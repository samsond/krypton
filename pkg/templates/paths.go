package templates

import (
	"path/filepath"
)

func GetDeploymentTemplatePath() string {
	return filepath.Join("pkg", "templates", "deployment.tmpl")
}
