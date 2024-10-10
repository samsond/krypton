package parser

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// GetName returns the name of the deployment
func (a *AppDeployment) GetName() string {
	return a.Name
}

// GetNamespace returns the namespace of the deployment
func (a *AppDeployment) GetNamespace() string {
	return a.Namespace
}

// parseAppDeployment parses an app deployment block from the DSL
func parseAppDeployment(line string, scanner *bufio.Scanner) (Resource, error) {
	app := &AppDeployment{
		Ports: make(map[string]int),
		Env:   make(map[string]string),
	}
	app.Name = strings.Trim(strings.TrimPrefix(line, "deploy app"), `" {`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "}" {
			break
		}
		if err := parseDeploymentLine(line, app, scanner); err != nil {
			return nil, err
		}
	}
	return app, nil
}

// parseDeploymentLine handles the details of an app deployment line.
func parseDeploymentLine(line string, app *AppDeployment, scanner *bufio.Scanner) error {
	switch {
	case strings.HasPrefix(line, "namespace:"):
		app.Namespace = parseStringValue(line)
	case strings.HasPrefix(line, "replicas:"):
		app.Replicas = parseIntValue(line)
	case strings.HasPrefix(line, "image:"):
		app.Image = parseStringValue(line)
	case strings.HasPrefix(line, "ports {"):
		return parsePorts(app, scanner)
	case strings.HasPrefix(line, "env {"):
		return parseEnv(app, scanner)
	case strings.HasPrefix(line, "resources {"):
		return parseResources(app, scanner)
	case strings.HasPrefix(line, "storage {"):
		return parseStorage(app, scanner)
	default:
		return fmt.Errorf("unknown directive in deployment: %s", line)
	}
	return nil
}

func parseResources(app *AppDeployment, scanner *bufio.Scanner) error {
	resources := &ResourceRequirements{
		Limits:   ResourceSpec{},
		Requests: ResourceSpec{},
	}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "}" {
			break // End of resources block.
		}

		if strings.HasPrefix(line, "limits {") {
			parseResourceSpec(&resources.Limits, scanner)
		} else if strings.HasPrefix(line, "requests {") {
			parseResourceSpec(&resources.Requests, scanner)
		} else {
			return fmt.Errorf("unknown directive in resources block: %s", line)
		}
	}

	app.Resources = resources
	return nil
}

func parseResourceSpec(spec *ResourceSpec, scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "}" {
			break // End of resource spec block.
		}

		if strings.HasPrefix(line, "memory:") {
			spec.Memory = parseStringValue(line)
		} else if strings.HasPrefix(line, "cpu:") {
			spec.CPU = parseStringValue(line)
		} else {
			return fmt.Errorf("unknown directive in resource spec: %s", line)
		}
	}
	return nil
}

func parseIntValue(line string) int {
	parts := strings.Split(line, ":")
	value, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	return value
}

func parseStringValue(line string) string {
	parts := strings.Split(line, ":")
	value := strings.TrimSpace(parts[1])
	// Remove leading and trailing quotes, if present.
	value = strings.Trim(value, `";`)
	return value
}

func parsePorts(app *AppDeployment, scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "}" {
			break // End of ports block.
		}

		// Expecting format: http: 8080;
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid port format: %s", line)
		}

		portName := strings.TrimSpace(parts[0])
		portValue, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[1], ";")))
		if err != nil {
			return fmt.Errorf("invalid port value: %s", parts[1])
		}

		app.Ports[portName] = portValue
	}
	return nil
}

func parseEnv(app *AppDeployment, scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "}" {
			break // End of env block.
		}

		// Expecting format: KEY: "value";
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid environment variable format: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		// Trim the quotes and semicolon from the value part
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, ` ";`)

		// Check if the value is properly formatted, just as a precaution.
		if key == "" || value == "" {
			return fmt.Errorf("invalid environment variable key or value: %s", line)
		}

		app.Env[key] = value
	}
	return nil
}

func parseStorage(app *AppDeployment, scanner *bufio.Scanner) error {
	storage := &StorageConfig{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "}" {
			break // End of storage block.
		}

		if strings.HasPrefix(line, "volume:") {
			storage.Volume = parseStringValue(line)
		} else if strings.HasPrefix(line, "size:") {
			storage.Size = parseStringValue(line)
		} else {
			return fmt.Errorf("unknown directive in storage block: %s", line)
		}
	}

	app.Storage = storage
	return nil
}
