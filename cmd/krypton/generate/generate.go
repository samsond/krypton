package generate

import (
	"fmt"
	"os"

	"github.com/samsond/krypton/pkg/parser"
    "github.com/samsond/krypton/pkg/generator"
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

            // Parse the DSL into a Resource.
            resource, err := parser.ParseDSL(dslFilePath)
            if err != nil {
                fmt.Fprintf(cmd.OutOrStderr(), "Error parsing DSL: %v\n", err)
                os.Exit(1)
            }

            // Use type assertion or type switch to handle different resource types
            appDeployment, ok := resource.(*parser.AppDeployment)
            if !ok {
                fmt.Fprintf(cmd.OutOrStderr(), "Unsupported resource type: %T\n", resource)
                os.Exit(1)
            }
            

            manifests, err := generator.GenerateFromApp(appDeployment)

            if err != nil {
                fmt.Fprintf(cmd.OutOrStderr(), "Error generating manifests: %v\n", err)
                os.Exit(1)
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