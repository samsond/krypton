package parser

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// parseFunc is a function type for parsing DSL commands
type parseFunc func(line string, scanner *bufio.Scanner) (Resource, error)

// resourceParsers maps DSL commands to their respective parsing functions
var resourceParsers = map[string]parseFunc{
    "deploy app":    parseAppDeployment,
    // Add more commands here as needed.
}

// ParseDSL reads a DSL file and parses it into a Resource
func  ParseDSL(filePath string) (Resource, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to read DSL file: %w", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var resource Resource

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        // Find the appropriate parser function based on the command.
        for command, parser := range resourceParsers {
            if strings.HasPrefix(line, command) {
                resource, err = parser(line, scanner)
                if err != nil {
                    return nil, fmt.Errorf("error parsing line '%s': %w", line, err)
                }
                break
            }
        }

        if resource == nil {
            return nil, fmt.Errorf("unknown DSL command: %s", line)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading DSL: %w", err)
    }

    return resource, nil
}
