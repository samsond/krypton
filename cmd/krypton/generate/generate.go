package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/samsond/krypton/pkg/generator"
	"github.com/samsond/krypton/pkg/lexer"
	"github.com/samsond/krypton/pkg/parser"
	"github.com/spf13/cobra"
)

// NewGenerateCommand creates a new generate command with dependency injection
func NewGenerateCommand() *cobra.Command {
	var outputFile string

	cmd := &cobra.Command{
		Use:   "generate [path to DSL script]",
		Short: "Generate Kubernetes manifests from a DSL script",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dslFilePath := args[0]

			// Initialize the lexer and parser
			// Read the DSL script from the file
			rawInput, err := os.ReadFile(dslFilePath)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "Error reading DSL file: %v\n", err)
				os.Exit(1)
			}

			// Trim the input DSL script
			input := strings.TrimSpace(string(rawInput))

			lexer := lexer.NewLexer(input)
			parser := parser.NewParser(lexer)

			// Parse the DSL into resources
			resources, err := parser.ParseResources()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "Error parsing resources: %v\n", err)
				os.Exit(1)
			}

			var manifests string
			for _, resource := range resources {
				yamlContent, err := generator.GenerateYAML(resource)
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "Error getting template path: %v\n", err)
					os.Exit(1)
				}

				manifests += yamlContent + "\n---\n"
			}

			// Write the generated YAML to the specified output file or stdout
			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(manifests), 0644)
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "Error writing to file: %v\n", err)
					os.Exit(1)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "Manifests written to %s\n", outputFile)
			} else {
				fmt.Fprint(cmd.OutOrStdout(), manifests)
			}
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for the generated Kubernetes manifests")
	return cmd
}
